package service

import (
	"github.com/nuvotlyuba/study-go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/store"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
	"go.uber.org/zap"
)

type Service struct {
	memRepo  store.MemRepository
	fileRepo store.FileRepository
}

func New(store *store.Store) *Service {
	return &Service{
		memRepo:  *store.MemRepo(),
		fileRepo: *store.FileRepo(),
	}
}

type Services interface {
	CreateURL(originalURL *models.URL) (*models.ObjURL, error)
	GetURL(shortURL *models.URL) (*models.URL, error)
}

func (s Service) CreateURL(originalURL *models.URL) (*models.ObjURL, error) {
	token := utils.MakeToken(8)

	newURL := models.ObjURL{
		UUID:        *utils.MakeToken(8),
		ShortURL:    *utils.MakeShortURL(token),
		OriginalURL: *originalURL,
	}
	storage := utils.SwitchStorage()
	switch storage {
	case "file":
		if err := s.fileRepo.WriteNewURL(&newURL); err != nil {
			return nil, err
		}
	case "mem":
		s.memRepo.SaveURL(&newURL)
	}
	logger.Info("create URL in", zap.String("storage", storage))

	return &newURL, nil
}

func (s Service) GetURL(shortURL *models.URL) (*models.URL, error) {
	var originalURL *models.URL
	var err error

	storage := utils.SwitchStorage()
	switch storage {
	case "file":
		if originalURL, err = s.fileRepo.ReadURL(shortURL); err != nil {
			return nil, err
		}
	case "mem":
		originalURL, err = s.memRepo.GetURL(shortURL)
		if err != nil {
			return nil, err
		}
	}
	logger.Info("get URL from", zap.String("storage", storage))

	return originalURL, nil
}
