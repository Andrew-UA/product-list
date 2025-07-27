package requests

import (
	"encoding/json"
	"net/http"
)

type BaseRequest struct {
}

func (r *BaseRequest) BeforeValidation() error {
	return nil
}

func (r *BaseRequest) AfterValidation() error {
	return nil
}

func ReadAndCLose(r *http.Request, data any) error {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	return d.Decode(data)
}
