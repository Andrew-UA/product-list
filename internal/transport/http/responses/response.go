package responses

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseInterface interface {
	Write(closer io.Writer)
}

type Response struct {
	statusCode int
}

func (r *Response) Write(w io.Writer) error {
	if rc, ok := w.(http.ResponseWriter); ok {
		rc.Header().Set("Content-Type", "application/json")
		rc.WriteHeader(r.statusCode)
	}

	return json.NewEncoder(w).Encode(r)
}

func (r *Response) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}
