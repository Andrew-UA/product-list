package handlers

import (
	"github.com/Andrew-UA/product-list/internal/config"
	"net/http"
)

type HealthHandler struct {
	config *config.Config
}

func NewHealthHandler(config *config.Config) *HealthHandler {
	return &HealthHandler{
		config: config,
	}
}

func (h *HealthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.config.AppName))
}
