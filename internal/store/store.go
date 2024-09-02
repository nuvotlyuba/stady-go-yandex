package store

import (
	"errors"

	"github.com/nuvotlyuba/study-go-yandex/config"
)

type Store struct {
	memRepository  *MemRepository
	fileRepository *FileRepository
}

func New() *Store {
	return &Store{}
}

var ErrNotFoundURL = errors.New("данная ссылка не найдена")

func (s *Store) MemRepo() *MemRepository {
	if s.memRepository != nil {
		return s.memRepository
	}
	s.memRepository = &MemRepository{
		data: DataURL,
	}
	return s.memRepository
}

func (s *Store) FileRepo() *FileRepository {
	if s.fileRepository != nil {
		return s.fileRepository
	}

	s.fileRepository = &FileRepository{
		FileStoragePath: config.FileStoragePath,
	}
	return s.fileRepository
}
