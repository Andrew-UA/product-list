package requests

import (
	"encoding/json"
	"errors"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/go-playground/validator/v10"
	"io"
)

type Request interface {
	ReadAndClose(data io.ReadCloser) error
	Validate() []error
	ToDTO() dto.DTO
}

type BaseRequest struct {
	RawMap map[string]json.RawMessage `json:"-"`
}

func (r *BaseRequest) ReadAndClose(reader io.ReadCloser) error {
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &r.RawMap)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, r)
}

func (r *BaseRequest) Validate() []error {
	validate := validator.New()
	err := validate.Struct(r)

	if err == nil {
		return nil
	}

	var errs []error
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldErr := range validationErrors {
			errs = append(errs, fieldErr)
		}
	}

	return errs
}

func (r *BaseRequest) HasField(fieldName string) (exists bool, isNull bool) {
	value, exists := r.RawMap[fieldName]
	if !exists {
		return false, false
	}

	return true, string(value) == "null"
}
