package schema

import (
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate        *validator.Validate
	once            sync.Once
	ethAddressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
)

type RequestValidator interface {
	Validate(i interface{}) error
	FormatErrors(err error) map[string]string
}

// initSchemaValidator initializes global **validate** var
func initSchemaValidator() {
	validate = validator.New()

	validate.RegisterValidation("eth_addr", validateEthAddress)
}

// GetSchemaValidator return validator instance to be used for
// validating a given schema
func GetSchemaValidator() RequestValidator {
	once.Do(initSchemaValidator)

	return &requestValidator{
		validator: validate,
	}
}

type requestValidator struct {
	validator *validator.Validate
}

func (rv *requestValidator) Validate(i interface{}) error {
	return rv.validator.Struct(i)
}

func (rv *requestValidator) FormatErrors(err error) map[string]string {
	errors := make(map[string]string)

	if err == nil {
		return errors
	}

	// Type assert to get validation errors
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["_general"] = "Validation failed"
		return errors
	}

	// Convert each error to a user-friendly message
	for _, e := range validationErrors {
		errors[e.Field()] = formatError(e)
	}

	return errors
}

// formatError converts a single validation error to a message
func formatError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "eth_addr":
		return "Must be a valid Ethereum address"
	default:
		return "Invalid value"
	}
}

// validateEthAddress is the custom validation function for Ethereum addresses
func validateEthAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	return ethAddressRegex.MatchString(address)
}
