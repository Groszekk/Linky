package middlewares

import (
	"strconv"

	"github.com/go-playground/validator"
)

func MinimalValidation(field validator.FieldLevel) bool {
	param, _ := strconv.Atoi(field.Param())
	if len(field.Field().String()) < param {
		return false
	}

	return true
}

func MaximumValidation(field validator.FieldLevel) bool {
	param, _ := strconv.Atoi(field.Param())
	if len(field.Field().String()) > param {
		return false
	}

	return true
}
