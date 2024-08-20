package handler

import (
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type Handlers interface {
	PostURL(w http.ResponseWriter, r *http.Request)
	GetURL(w http.ResponseWriter, r *http.Request)
}
