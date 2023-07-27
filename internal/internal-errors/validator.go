package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "min":
		return errors.New(field + " requires " + validationError.Param() + " minimum characters")
	case "max":
		return errors.New(field + " requires " + validationError.Param() + " maximum characters")
	case "email":
		return errors.New(field + " is an invalid email")
	}
	return nil
}
