package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/study-go-yandex/internal/repository"
	"github.com/nuvotlyuba/study-go-yandex/internal/service"
	"github.com/nuvotlyuba/study-go-yandex/internal/transport/handler"
)

type APIServer struct {
	config *APIConfig
	router *chi.Mux
}

func New(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	repoVar := repository.NewVarRepository()
	service := service.New(repoVar)
	handler := handler.New(service)

	s.addRouter(handler)

	server := &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: s.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) addRouter(h *handler.Handler) *chi.Mux {
	s.router.HandleFunc(`/`, h.Handler)
	s.router.Get("/{id}", h.GetURL)
	s.router.Post("/", h.PostURL)

	return s.router
}
