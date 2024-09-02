package store

import "github.com/nuvotlyuba/study-go-yandex/internal/models"

type MemRepository struct {
	data models.URLData
}

var DataURL = make(map[models.URL]models.URL)

type MemRepo interface {
	SaveURL(shotURL *models.URL, originalURL *models.URL)
	GetURL(shotURL *models.URL) (*models.URL, error)
}

func (r MemRepository) SaveURL(data *models.ObjURL) {
	DataURL[*&data.ShortURL] = *&data.OriginalURL
}

func (r MemRepository) GetURL(shotURL *models.URL) (*models.URL, error) {
	originalURL, ok := DataURL[*shotURL]
	if !ok {
		return nil, ErrNotFoundURL
	}
	return &originalURL, nil
}
