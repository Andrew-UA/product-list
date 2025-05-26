package validation

import "github.com/go-playground/validator/v10"

type Validatable interface {
	BeforeValidation() error
	AfterValidation() error
}

type ValidatorInterface interface {
	ValidateStruct(data Validatable) error
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

func (s *Validator) ValidateStruct(data Validatable) error {
	if err := data.BeforeValidation(); err != nil {
		return err
	}

	if err := s.validator.Struct(data); err != nil {
		return err
	}

	if err := data.AfterValidation(); err != nil {
		return err
	}

	return nil
}
