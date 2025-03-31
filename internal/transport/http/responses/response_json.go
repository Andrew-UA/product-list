package responses

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseInterface interface {
	Write(closer io.Writer) error
}

type ResponseJson struct {
	statusCode int
}

func (r *ResponseJson) Write(w io.Writer) error {
	if rc, ok := w.(http.ResponseWriter); ok {
		rc.Header().Set("Content-Type", "application/json")
		rc.WriteHeader(r.statusCode)
	}

	return json.NewEncoder(w).Encode(r)
}

func (r *ResponseJson) SetStatusCode(statusCode int) ResponseInterface {
	r.statusCode = statusCode
	return r
}
