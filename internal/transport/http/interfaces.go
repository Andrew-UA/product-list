package http

import "net/http"

type HealthHandlerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
	Health(w http.ResponseWriter, r *http.Request)
}

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type ResourceHandlerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
