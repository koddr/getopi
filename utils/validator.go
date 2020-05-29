package utils

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	emailRegexpPattern string = `^.+@.+\..+$`
)

// Validate create a new validator for expected fields,
// register function to get tag name from `json` tags
// and add validation to expected fields
func Validate(obj string) *validator.Validate {
	// Create a new validator
	validate := validator.New()

	// Get tag name from `json`
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	switch obj {
	case "auth":
		// Check for regexp parrtern and length (varchar(255))
		validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return regexp.MustCompile(emailRegexpPattern).MatchString(field) && len(field) <= 254
		})

		// Check for length
		validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return len(field) >= 6
		})
	case "forget-password":
		// Check for regexp parrtern and length (varchar(255))
		validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return regexp.MustCompile(emailRegexpPattern).MatchString(field) && len(field) <= 254
		})
	case "reset-code":
		// Check for length (varchar(6))
		validate.RegisterValidation("reset_code", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return len(field) == 6
		})
	case "password":
		// Check for length (min: 6)
		validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return len(field) >= 6
		})
	case "user":
		// Check for valid UUID
		validate.RegisterValidation("id", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			if _, errParseUUID := uuid.Parse(field); errParseUUID != nil {
				return true
			}
			return false
		})

		// Check for regexp parrtern and length (varchar(255))
		validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return regexp.MustCompile(emailRegexpPattern).MatchString(field) && len(field) <= 254
		})

		// Check for length (varchar(13))
		validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
			field := fl.Field().String()
			return len(field) <= 13
		})
	}

	return validate
}

// ValidateErrors show validation error for each invalid fields
func ValidateErrors(err error) map[string]string {
	errorFields := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		// Make error message
		errorFields[err.Field()] = "field " + err.StructField() + " is not valid"
	}

	return errorFields
}
