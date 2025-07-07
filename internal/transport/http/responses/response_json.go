package responses

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type ResponseInterface interface {
	Write(closer io.Writer)
	SetStatusCode(statusCode int) ResponseInterface
	SetData(data any) ResponseInterface
}

type ResponseJson struct {
	statusCode int
	data       any
}

func NewResponseJson(statusCode int, data any) *ResponseJson {
	return &ResponseJson{
		statusCode: statusCode,
		data:       data,
	}
}

func (r *ResponseJson) Write(w io.Writer) {
	if rc, ok := w.(http.ResponseWriter); ok {
		rc.Header().Set("Content-Type", "application/json")
		rc.WriteHeader(r.statusCode)
	}

	err := json.NewEncoder(w).Encode(r.data)
	if err != nil {
		log.Err(err).Msg("error encoding response")
	}
}

func (r *ResponseJson) SetStatusCode(statusCode int) ResponseInterface {
	r.statusCode = statusCode
	return r
}

func (r *ResponseJson) SetData(data any) ResponseInterface {
	r.data = data
	return r
}
