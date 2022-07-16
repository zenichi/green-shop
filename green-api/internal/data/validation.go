package data

import "github.com/go-playground/validator/v10"

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
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {

			// Examples from Doc:
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())

			errors = append(errors, err.Field())
		}

		return errors
	}

	// validated
	return nil
}
