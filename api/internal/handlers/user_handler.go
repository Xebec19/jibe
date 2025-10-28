package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Xebec19/jibe/api/internal/domain"
	"github.com/Xebec19/jibe/api/internal/service"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/gorilla/mux"
)

// UserHandler handles user-related HTTP requests
// It follows the Handler pattern with dependency injection
type UserHandler struct {
	userService service.UserService
	logger      *logger.Logger
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService service.UserService, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

// GetUsers handles GET /api/v1/users
// @Summary      List users
// @Description  Get a list of users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit   query     int  false  "Limit"  default(10)
// @Param        offset  query     int  false  "Offset" default(0)
// @Success      200     {object}  Response{data=[]domain.User}
// @Failure      500     {object}  Response
// @Router       /users [get]
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse query parameters for pagination
	limit := 10
	offset := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	// Get users from service
	users, err := h.userService.ListUsers(ctx, limit, offset)
	if err != nil {
		h.logger.Error().
			Err(err).
			Msg("Failed to list users")
		respondError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    users,
	})
}

// GetUser handles GET /api/v1/users/{id}
// @Summary      Get a user
// @Description  Get a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  Response{data=domain.User}
// @Failure      400  {object}  Response
// @Failure      404  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract user ID from URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Get user from service
	user, err := h.userService.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respondError(w, http.StatusNotFound, "User not found")
			return
		}
		h.logger.Error().
			Err(err).
			Int64("user_id", id).
			Msg("Failed to get user")
		respondError(w, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    user,
	})
}

// CreateUser handles POST /api/v1/users
// @Summary      Create a user
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      domain.CreateUserRequest  true  "User data"
// @Success      201   {object}  Response{data=domain.User}
// @Failure      400   {object}  Response
// @Failure      409   {object}  Response
// @Failure      500   {object}  Response
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse request body
	var req domain.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Create user through service
	user, err := h.userService.CreateUser(ctx, &req)
	if err != nil {
		// Handle different error types
		var validationErr domain.ErrInvalidInput
		if errors.As(err, &validationErr) {
			respondError(w, http.StatusBadRequest, validationErr.Error())
			return
		}
		if errors.Is(err, domain.ErrAlreadyExists) {
			respondError(w, http.StatusConflict, "User with this email already exists")
			return
		}

		h.logger.Error().
			Err(err).
			Str("email", req.Email).
			Msg("Failed to create user")
		respondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	respondJSON(w, http.StatusCreated, Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

// UpdateUser handles PUT /api/v1/users/{id}
// @Summary      Update a user
// @Description  Update an existing user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int                       true  "User ID"
// @Param        user  body      domain.UpdateUserRequest  true  "User data"
// @Success      200   {object}  Response{data=domain.User}
// @Failure      400   {object}  Response
// @Failure      404   {object}  Response
// @Failure      409   {object}  Response
// @Failure      500   {object}  Response
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract user ID from URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Parse request body
	var req domain.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Update user through service
	user, err := h.userService.UpdateUser(ctx, id, &req)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respondError(w, http.StatusNotFound, "User not found")
			return
		}
		if errors.Is(err, domain.ErrAlreadyExists) {
			respondError(w, http.StatusConflict, "User with this email already exists")
			return
		}

		h.logger.Error().
			Err(err).
			Int64("user_id", id).
			Msg("Failed to update user")
		respondError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

// DeleteUser handles DELETE /api/v1/users/{id}
// @Summary      Delete a user
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      404  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract user ID from URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Delete user through service
	if err := h.userService.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respondError(w, http.StatusNotFound, "User not found")
			return
		}

		h.logger.Error().
			Err(err).
			Int64("user_id", id).
			Msg("Failed to delete user")
		respondError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "User deleted successfully",
	})
}
