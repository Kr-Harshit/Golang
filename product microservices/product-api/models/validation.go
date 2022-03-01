package models

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError wrpas the validator FieldError
type ValidationError struct {
	error validator.FieldError
}

// Error convert the ValidationError to String
func (v *ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Erorr: Field validation for '%s' failed on the '%s' tag",
		v.error.Namespace(),
		v.error.Field(),
		v.error.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Erros convert all ValidationErrors to slice of string
func (v *ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range *v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation wrapper
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation Type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate: validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs ValidationErrors
	for _, err := range errs {
		// casting FieldError to ValidationError and appending to the slice
		ve := ValidationError{error: err}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}

// validateSKU
func validateSKU(fl validator.FieldLevel) bool {
	// SKU format is  abc-abc-abc
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	sku := re.FindAllString(fl.Field().String(), -1)

	return len(sku) == 1
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}
