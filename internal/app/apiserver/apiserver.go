package apiserver

import (
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/repository"
	"github.com/nuvotlyuba/study-go-yandex/internal/service"
	"github.com/nuvotlyuba/study-go-yandex/internal/transport/handler"
)

type APIServer struct {
	config *APIConfig
	router *http.ServeMux
}

func New(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
		router: http.NewServeMux(),
	}
}

func (s *APIServer) Start() error {
	repo := repository.New()
	service := service.New(repo)
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

func (s *APIServer) addRouter(h *handler.Handler) *http.ServeMux {
	s.router.HandleFunc(`/`, h.GetURL)

	return s.router
}
