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
	Handler(w http.ResponseWriter, r *http.Request)
}

func (h Handler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.PostURL(w, r)
	case http.MethodGet:
		h.GetURL(w, r)
	default:
		http.Error(w, "Неверный метод", http.StatusBadRequest)
		return
	}
}
