package validation

import "github.com/go-playground/validator/v10"

type ValidatorInterface interface {
	ValidateStruct(data any) error
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

func (s *Validator) ValidateStruct(data any) error {
	return s.validator.Struct(data)
}
