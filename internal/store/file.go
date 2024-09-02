package store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
)

type URLRecorder struct {
	file   *os.File
	writer *bufio.Writer
}

func NewURLRecorder(filename string) (*URLRecorder, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &URLRecorder{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (w *URLRecorder) Close() error {
	return w.file.Close()
}

func (w *URLRecorder) WriteURL(url *models.ObjURL) error {
	data, err := json.Marshal(&url)
	if err != nil {
		return err
	}

	// записываем событие в буфер
	if _, err = w.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := w.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return w.writer.Flush()
}

type URLScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewURLScanner(filename string) (*URLScanner, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &URLScanner{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (s *URLScanner) ScanURL(shortURL models.URL) (*models.ObjURL, error) {
	url := models.ObjURL{}
	var d models.ObjURL

	for s.scanner.Scan() {
		data := s.scanner.Bytes()
		if err := json.Unmarshal(data, &url); err != nil {
			return nil, err
		}

		if url.ShortURL == shortURL {
			d = url
			break
		}
	}

	if err := s.scanner.Err(); err != nil {
		return nil, err
	}

	return &d, nil
}

func (s *URLScanner) Split() {
	s.scanner.Split(bufio.ScanLines)
}

type FileRepository struct {
	FileStoragePath string
}

func NewFileRepository(fileStoragePath string) *FileRepository {
	return &FileRepository{
		FileStoragePath: fileStoragePath,
	}
}

type FileRepo interface {
	ReadURL()
	WriteURL()
}

func (f *FileRepository) WriteNewURL(data *models.ObjURL) error {
	w, err := NewURLRecorder(f.FileStoragePath)
	if err != nil {
		return fmt.Errorf("error in FileRepository: WriteNewURL: %v", err)
	}
	err = w.WriteURL(data)
	if err != nil {
		return fmt.Errorf("error in FileRepository: WriteURL: %v", err)
	}

	return nil
}

func (f *FileRepository) ReadURL(shortURL *models.URL) (*models.URL, error) {
	rr, err := NewURLScanner(f.FileStoragePath)
	if err != nil {
		return nil, fmt.Errorf("error in FileRepository: ReadURL: %v", err)
	}

	rr.Split()
	data, err := rr.ScanURL(*shortURL)
	if err != nil {
		return nil, fmt.Errorf("error in FileRepository: ReadURL -> %v", err)
	}

	return &data.OriginalURL, nil
}
