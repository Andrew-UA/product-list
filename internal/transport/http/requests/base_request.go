package requests

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
	"reflect"
)

type Request interface {
	ReadAndClose(data io.ReadCloser) error
	Validate() []error
	ToMap() map[string]any
}

type BaseRequest struct {
	RawData json.RawMessage `json:"-"`
}

func (r *BaseRequest) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if len(r.RawData) == 0 {
		return result
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(r.RawData, &rawMap); err != nil {
		return result
	}

	rv := reflect.ValueOf(r).Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Field(i)

		if field.Name == "RawData" {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" || jsonTag == "" {
			continue
		}

		if _, exists := rawMap[jsonTag]; !exists {
			continue
		}

		result[jsonTag] = fieldValue.Interface()
	}

	return result
}

func (r *BaseRequest) ReadAndClose(reader io.ReadCloser) error {
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	r.RawData = body

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
