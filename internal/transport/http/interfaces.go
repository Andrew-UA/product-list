package http

import "net/http"

type HealthHandlerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
	Health(w http.ResponseWriter, r *http.Request)
}
