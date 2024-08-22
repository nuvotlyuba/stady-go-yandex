package repository

import (
	"errors"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
)

type Repository struct{}

func NewVarRepository() VarRepository {
	return &Repository{}
}

var ErrNotFoundURL = errors.New("данная ссылка не найдена")

var DataURL = make(map[models.URL]models.URL)

type VarRepository interface {
	SaveURL(shotURL *models.URL, longURL *models.URL)
	GetURL(shotURL *models.URL) (*models.URL, error)
}

func (r Repository) SaveURL(shotURL *models.URL, longURL *models.URL) {
	DataURL[*shotURL] = *longURL
}

func (r Repository) GetURL(shotURL *models.URL) (*models.URL, error) {
	longURL, ok := DataURL[*shotURL]
	if !ok {
		return nil, ErrNotFoundURL
	}
	return &longURL, nil
}
