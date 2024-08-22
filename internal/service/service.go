package service

import (
	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/repository"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
)

type Service struct {
	repoVar repository.VarRepository
}

func New(repoVar repository.VarRepository) *Service {
	return &Service{
		repoVar: repoVar,
	}
}

type Services interface {
	CreateURL(longURL *models.URL) (*models.URL, error)
	GetURL(shortURL *models.URL) (*models.URL, error)
}

func (s Service) CreateURL(longURL *models.URL) (*models.URL, error) {
	token := utils.MakeToken(8)
	shortURL := utils.MakeShortURL(token)
	s.repoVar.SaveURL(shortURL, longURL)
	return shortURL, nil
}

func (s Service) GetURL(shortURL *models.URL) (*models.URL, error) {
	longURL, err := s.repoVar.GetURL(shortURL)
	if err != nil {
		return nil, err
	}
	return longURL, nil
}
