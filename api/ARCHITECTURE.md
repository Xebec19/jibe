# Architecture Documentation

## Overview

This application follows **Clean Architecture** principles with clear separation of concerns across multiple layers. Each layer has specific responsibilities and dependencies flow inward (from infrastructure toward domain).

## Architectural Layers

```
┌─────────────────────────────────────────────────────────┐
│                   HTTP Layer (main.go)                   │
│              Router, Middleware, Server Setup            │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│              Handler Layer (internal/handlers)           │
│         HTTP Request/Response, Serialization             │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│              Service Layer (internal/service)            │
│         Business Logic, Validation, Orchestration        │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│           Repository Layer (internal/repository)         │
│              Data Access Abstraction (Interface)         │
└─────┬──────────────────────────────────────────┬────────┘
      │                                           │
┌─────▼──────────────┐                  ┌────────▼────────┐
│ Memory Repository  │                  │ PostgreSQL Repo │
│   (Development)    │                  │   (Production)  │
└────────────────────┘                  └─────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│              Domain Layer (internal/domain)              │
│         Entities, Business Rules, Domain Errors          │
└─────────────────────────────────────────────────────────┘
```

## Dependency Flow

Dependencies always point **inward** toward the domain:

```
Handler → Service → Repository → Domain
  ↓         ↓          ↓           ↓
Logger   Logger    Logger      (none)
```

- **Domain** has no external dependencies
- **Repository** depends only on Domain
- **Service** depends on Repository and Domain
- **Handler** depends on Service and Domain

## Design Patterns

### 1. Dependency Injection (DI)

**Location**: `internal/container/container.go`

The Container pattern wires all dependencies together:

```go
type Container struct {
    Config         *config.Config
    Logger         *logger.Logger
    UserRepository repository.UserRepository
    UserService    service.UserService
    CircuitBreaker *resilience.CircuitBreaker
    RetryPolicy    *resilience.RetryPolicy
}
```

**Benefits**:
- Single place to manage all dependencies
- Easy to swap implementations
- Simplifies testing with mocks
- Clear visibility of application structure

### 2. Repository Pattern

**Location**: `internal/repository/`

Abstracts data access behind interfaces:

```go
type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    GetByID(ctx context.Context, id int64) (*domain.User, error)
    // ...
}
```

**Implementations**:
- `memory/` - In-memory (fast, for dev/test)
- `postgres/` - PostgreSQL (production)

**Benefits**:
- Database-agnostic business logic
- Easy to test with in-memory repo
- Can switch databases without code changes
- Supports multiple data sources

### 3. Service Layer Pattern

**Location**: `internal/service/`

Encapsulates business logic:

```go
type UserService interface {
    CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
    GetUser(ctx context.Context, id int64) (*User, error)
    // ...
}
```

**Responsibilities**:
- Business rule enforcement
- Input validation
- Transaction coordination
- Error handling
- Logging business events

**Benefits**:
- Business logic separate from HTTP concerns
- Reusable across different interfaces (HTTP, gRPC, CLI)
- Easier to test business logic in isolation

### 4. Handler Pattern

**Location**: `internal/handlers/`

Struct-based handlers with dependency injection:

```go
type UserHandler struct {
    userService service.UserService
    logger      *logger.Logger
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    // Handle HTTP specifics
    users, err := h.userService.ListUsers(ctx, limit, offset)
    // Serialize response
}
```

**Benefits**:
- Dependencies injected at creation
- Easy to test with mock services
- Clear separation from business logic

## Resilience Patterns

### 1. Circuit Breaker

**Location**: `pkg/resilience/circuit_breaker.go`

Prevents cascading failures by stopping requests to failing services:

```
[Closed] ──(too many failures)──> [Open] ──(timeout)──> [Half-Open]
    ↑                                                         │
    └───────────────────(success)─────────────────────────────┘
```

**States**:
- **Closed**: Normal operation
- **Open**: Fail fast, don't call service
- **Half-Open**: Try one request to test recovery

**Configuration**:
```go
CircuitBreaker(
    maxFailures: 5,      // Open after 5 failures
    timeout: 10s,        // Request timeout
    resetTimeout: 30s    // Try recovery after 30s
)
```

### 2. Retry with Exponential Backoff

**Location**: `pkg/resilience/retry.go`

Automatically retries failed operations with increasing delays:

```
Attempt 1: immediate
Attempt 2: wait 100ms
Attempt 3: wait 200ms
Attempt 4: wait 400ms
Attempt 5: wait 800ms (capped at maxBackoff)
```

**Features**:
- Exponential backoff: delay *= multiplier
- Maximum backoff cap
- Optional jitter to prevent thundering herd
- Context-aware cancellation

**Configuration**:
```go
RetryPolicy(
    maxRetries: 3,
    initialBackoff: 100ms,
    backoffMultiplier: 2.0,
    maxBackoff: 5000ms
)
```

## Request Flow

Here's how a typical request flows through the system:

```
1. HTTP Request
   ↓
2. Middleware Stack
   ├─ RequestID: Add unique ID
   ├─ Logger: Log request
   ├─ Recoverer: Catch panics
   └─ CORS: Handle CORS headers
   ↓
3. Router (Gorilla Mux)
   ↓
4. Handler (e.g., UserHandler.CreateUser)
   ├─ Parse request body
   ├─ Basic validation
   └─ Call service
   ↓
5. Service (UserService.CreateUser)
   ├─ Domain validation
   ├─ Business rules
   ├─ Check for duplicates (via repository)
   └─ Create user (via repository)
   ↓
6. Repository (UserRepository.Create)
   ├─ Database interaction
   └─ Error translation
   ↓
7. Response
   ├─ Service returns domain entity
   ├─ Handler serializes to JSON
   └─ Middleware logs response
```

## Testing Strategy

### Unit Tests

**Domain Layer**:
```go
// Test business rules in isolation
func TestUserValidation(t *testing.T) {
    user := &domain.CreateUserRequest{Name: "", Email: "test@example.com"}
    err := user.Validate()
    // Assert error
}
```

**Service Layer**:
```go
// Test with mock repository
func TestUserService_CreateUser(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo, logger)
    // Test business logic
}
```

**Handler Layer**:
```go
// Test with mock service
func TestUserHandler_CreateUser(t *testing.T) {
    mockService := &MockUserService{}
    handler := NewUserHandler(mockService, logger)
    // Test HTTP handling
}
```

### Integration Tests

**Repository Layer**:
```go
// Test with real database
func TestUserRepository_Create(t *testing.T) {
    db := setupTestDB()
    repo := postgres.NewUserRepository(db, logger)
    // Test database operations
}
```

### End-to-End Tests

```go
// Test full request flow
func TestAPI_CreateUser(t *testing.T) {
    server := setupTestServer()
    resp := httptest.NewRecorder()
    // Make HTTP request
    // Assert response
}
```

## Error Handling

### Domain Errors

Defined in `internal/domain/errors.go`:
```go
var (
    ErrNotFound       = errors.New("resource not found")
    ErrAlreadyExists  = errors.New("resource already exists")
    ErrUnauthorized   = errors.New("unauthorized")
    // ...
)
```

### Error Flow

```
Repository → Service → Handler → HTTP Response

ErrNotFound → propagate → 404 Not Found
ErrAlreadyExists → propagate → 409 Conflict
ErrInvalidInput → propagate → 400 Bad Request
Other errors → log → 500 Internal Server Error
```

### Error Handling in Handlers

```go
user, err := h.userService.CreateUser(ctx, &req)
if err != nil {
    switch {
    case errors.Is(err, domain.ErrAlreadyExists):
        respondError(w, http.StatusConflict, "User already exists")
    case errors.As(err, &domain.ErrInvalidInput{}):
        respondError(w, http.StatusBadRequest, err.Error())
    default:
        h.logger.Error().Err(err).Msg("Failed to create user")
        respondError(w, http.StatusInternalServerError, "Internal error")
    }
    return
}
```

## Logging Strategy

### Structured Logging

Using `zerolog` for structured JSON logs:

```go
log.Info().
    Str("user_id", userID).
    Str("action", "create_user").
    Dur("duration", elapsed).
    Msg("User created successfully")
```

### Log Levels

- **Debug**: Development information, detailed flow
- **Info**: Important business events (user created, order placed)
- **Warn**: Recoverable errors, validation failures
- **Error**: Unhandled errors, service failures
- **Fatal**: Unrecoverable errors, application shutdown

### Context Propagation

Request ID flows through all layers:

```go
requestID := middleware.GetRequestID(ctx)
log.WithRequestID(requestID).Info().Msg("Processing request")
```

## Configuration Management

### Environment-Based Config

`internal/config/config.go` loads configuration from:
1. Environment variables
2. `.env` file (development)
3. Default values

### Validation

Production configs are validated:
```go
if cfg.Environment == "production" {
    if cfg.JWTSecret == "" {
        return fmt.Errorf("JWT_SECRET required in production")
    }
}
```

## Database Strategy

### Development: In-Memory Repository

- Fast startup, no setup required
- Perfect for local development
- Data lost on restart
- Thread-safe with sync.RWMutex

### Production: PostgreSQL Repository

- Persistent storage
- ACID transactions
- Connection pooling
- Prepared statements for performance
- Migrations in `internal/repository/postgres/migrations.sql`

### Switching Implementations

In `internal/container/container.go`:
```go
// Development
c.UserRepository = memory.NewUserRepository()

// Production
db := connectDB(cfg.DatabaseURL)
c.UserRepository = postgres.NewUserRepository(db, log)
```

## Scalability Considerations

### Horizontal Scaling

- **Stateless design**: No session state in memory
- **Database connection pooling**: Configurable pool size
- **Read replicas**: Repository pattern supports read/write splitting

### Performance

- **Context deadlines**: Prevent long-running requests
- **Connection pooling**: Reuse database connections
- **Prepared statements**: Faster query execution
- **Indexing**: Database indexes on common queries

### Monitoring

- **Request ID tracking**: Trace requests across services
- **Structured logging**: Easy to parse and analyze
- **Health checks**: `/health` endpoint for load balancers
- **Metrics**: Duration logging for performance monitoring

## Future Enhancements

- [ ] Add caching layer (Redis)
- [ ] Implement rate limiting per user
- [ ] Add distributed tracing (OpenTelemetry)
- [ ] Add metrics collection (Prometheus)
- [ ] Implement event sourcing for audit trail
- [ ] Add message queue integration
- [ ] Implement feature flags
- [ ] Add API versioning strategy
- [ ] Implement GraphQL endpoint
- [ ] Add WebSocket support
