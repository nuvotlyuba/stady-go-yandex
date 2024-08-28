package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/study-go-yandex/config"
	"github.com/nuvotlyuba/study-go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/study-go-yandex/internal/repository"
	"github.com/nuvotlyuba/study-go-yandex/internal/service"
	"github.com/nuvotlyuba/study-go-yandex/internal/transport/handler"
	"go.uber.org/zap"
)

type APIServer struct {
	config *APIConfig
	router *chi.Mux
	logger *zap.Logger
}

func New(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.addLogger(); err != nil {
		s.logger.Fatal("Don't add logger to server", zap.Error(err))
	}
	s.logger.Info("Server running ...", zap.String("address", s.config.ServerAddress))

	s.router.Use(logger.Middleware)
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

	s.logger.Info("Server running ...")

	return nil
}

func (s *APIServer) addRouter(h *handler.Handler) *chi.Mux {
	s.router.Get("/{id}", h.GetURL)
	s.router.Post("/", h.PostURL)
	s.router.Post("/api/shorten", h.PostJsonURL)

	return s.router
}

func (s *APIServer) addLogger() error {
	var logger *zap.Logger
	var err error
	if s.config.AppLevel == config.DEVELOPMENT {
		logger, err = zap.NewDevelopment()
	}
	if s.config.AppLevel == config.PRODUCTION {
		logger, err = zap.NewProduction()
	}

	s.logger = logger

	defer logger.Sync()

	zap.ReplaceGlobals(s.logger)

	return err
}
