package controller

import (
	"github.com/EmreZURNACI/url-shortener/app"
)

type Handler struct {
	Repository app.Repository
}

func NewRepository(handler app.Repository) *Handler {
	return &Handler{
		Repository: handler,
	}
}
