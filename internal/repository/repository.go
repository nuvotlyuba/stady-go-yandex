package repository

import (
	"fmt"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

var DataURL = make(map[models.URL]models.URL)

func (r Repository) SaveURL(shotURL *models.URL, longURL *models.URL) {
	DataURL[*shotURL] = *longURL
}

func (r Repository) GetURL(shotURL *models.URL) (*models.URL, error) {
	longURL, ok := DataURL[*shotURL]
	if !ok {
		return nil, fmt.Errorf("Данной %s ссылки нет", *shotURL)
	}
	return &longURL, nil
}
