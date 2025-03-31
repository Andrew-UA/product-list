package http

import (
	"github.com/Andrew-UA/product-list/internal/config"
	"net/http"
)

const (
	GET     = "GET "
	POST    = "POST "
	PUT     = "PUT "
	DELETE  = "DELETE "
	PATCH   = "PATCH "
	OPTIONS = "OPTIONS "

	v1 = "/api/v1"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

type Router struct {
	Mux    *http.ServeMux
	Config *config.Config

	// Middlewares
	AuthMiddleware Middleware

	// Handlers
	HealthHandler HealthHandlerInterface
	AuthHandler   AuthHandlerInterface
	UserHandler   ResourceHandlerInterface
}

func NewRouter(
	config *config.Config,
	authMiddleware Middleware,
	healthHandler HealthHandlerInterface,
	authHandler AuthHandlerInterface,
	userHandler ResourceHandlerInterface,
) *Router {
	mux := http.NewServeMux()

	router := &Router{
		Mux:    mux,
		Config: config,

		// Middleware
		AuthMiddleware: authMiddleware,

		// Handlers
		HealthHandler: healthHandler,
		AuthHandler:   authHandler,
		UserHandler:   userHandler,
	}

	router.registerHealthRoutes()
	router.registerAuthRoutes()
	router.registerUserRoutes()

	return router
}

func (r *Router) registerHealthRoutes() {
	r.Mux.HandleFunc(GET+"/health", r.HealthHandler.Health)
	r.Mux.HandleFunc(GET+"/ping", r.HealthHandler.Ping)
}

func (r *Router) registerAuthRoutes() {
	r.Mux.HandleFunc(POST+v1+"/login", r.AuthHandler.Login)
	r.Mux.HandleFunc(POST+v1+"/logout", r.AuthMiddleware(r.AuthHandler.Logout))
}

func (r *Router) registerUserRoutes() {
	r.Mux.HandleFunc(GET+v1+"/users", r.AuthMiddleware(r.UserHandler.Index))
	r.Mux.HandleFunc(GET+v1+"/users/{id}}", r.AuthMiddleware(r.UserHandler.Show))
	r.Mux.HandleFunc(POST+v1+"/users", r.AuthMiddleware(r.UserHandler.Create))
	r.Mux.HandleFunc(PUT+v1+"/users/{id}", r.AuthMiddleware(r.UserHandler.Update))
	r.Mux.HandleFunc(DELETE+v1+"/users/{id}", r.AuthMiddleware(r.UserHandler.Delete))
}
