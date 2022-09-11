package validator

import (
	goValidator "github.com/go-playground/validator/v10"
)

var validator = goValidator.New()

func ValidateStruct(structToValidate interface{}) error {
	return validator.Struct(structToValidate)
}
