package schema

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

// GetSchemaValidator return validator instance to be used for
// validating a given schema
func GetSchemaValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})

	return validate
}
