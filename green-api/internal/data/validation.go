package data

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationErrors struct {
	validator.ValidationErrors
}

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, f := range v.ValidationErrors {
		s := fmt.Sprintf(
			"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
			f.Namespace(),
			f.Field(),
			f.Tag(),
		)
		errs = append(errs, s)
	}

	return errs
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	// use a single instance of Validate, it caches struct info
	v := validator.New()
	return &Validator{v}
}

func (v *Validator) Validate(s interface{}) []string {
	err := v.validate.Struct(s)
	if err != nil {
		return ValidationErrors{err.(validator.ValidationErrors)}.Errors()
	}

	// validated
	return nil
}
