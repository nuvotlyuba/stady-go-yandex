package service

import (
	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/repository"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) CreateURL(longURL *models.URL) (*models.URL, error) {
	token := utils.MakeToken(8)
	shortURL := utils.MakeShortURL(token)
	s.repo.SaveURL(shortURL, longURL)
	return shortURL, nil
}

func (s Service) GetURL(shortURL *models.URL) (*models.URL, error) {
	longURL, err := s.repo.GetURL(shortURL)
	if err != nil {
		return nil, err
	}
	return longURL, nil
}
