package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/nuvotlyuba/study-go-yandex/config"
	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/service"
	"github.com/nuvotlyuba/study-go-yandex/internal/store"
	"github.com/nuvotlyuba/study-go-yandex/internal/types"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostURL(t *testing.T) {
	config.FileStoragePath = "test.json"
	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name    string
		request string
		body    string
		method  string
		want    want
	}{
		{
			name:    "POST запрос",
			request: "/",
			body:    "https://yandex.ru",
			method:  http.MethodPost,
			want: want{
				contentType: types.TextContentType,
				statusCode:  201,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.request, strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", tt.want.contentType)
			repo := store.New()
			s := service.New(repo)
			h := New(s)
			h.PostURL(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)
			assert.Contains(t, res.Header.Get("Content-Type"), tt.want.contentType, "Отличный от %s  Conent-Type", tt.want.contentType)

			body, err := io.ReadAll(res.Body)
			fmt.Println(string(body), "BODY")
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)
			assert.NotEmpty(t, string(body), "Тело ответа пустое")

			os.Remove(config.FileStoragePath)

		})
	}
}

func TestGetURL(t *testing.T) {
	config.FileStoragePath = "test.json"
	token := utils.MakeToken(8)
	EnsureNewURL(*token, "https://yandex.ru")
	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name    string
		request string
		method  string
		want    want
	}{
		{
			name:    "GET запрос",
			request: fmt.Sprintf("/%s", *token),
			method:  http.MethodGet,
			want: want{
				contentType: types.TextContentType,
				statusCode:  307,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.request, nil)
			w := httptest.NewRecorder()
			repo := store.New()
			s := service.New(repo)
			h := New(s)
			h.GetURL(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Contains(t, res.Header.Get("Content-Type"), tt.want.contentType, "Отличный от %s  Conent-Type", tt.want.contentType)
			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)

			os.Remove(config.FileStoragePath)
		})
	}
}

func TestPostJSONURL(t *testing.T) {
	config.FileStoragePath = "test.json"

	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name    string
		request string
		body    string
		method  string
		want    want
	}{
		{
			name:    "POST запрос JSON",
			request: "/api/shorten",
			body:    `{ "url": "https://yandex.ru" }`,
			method:  http.MethodPost,
			want: want{
				contentType: types.JSONContentType,
				statusCode:  201,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.request, strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", tt.want.contentType)
			repo := store.New()
			s := service.New(repo)
			h := New(s)
			h.PostJSONURL(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Contains(t, res.Header.Get("Content-Type"), tt.want.contentType, "Отличный от %s  Conent-Type", tt.want.contentType)
			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)
			assert.NotEmpty(t, string(body), "Тело ответа пустое")

			os.Remove(config.FileStoragePath)
		})
	}
}

func EnsureNewURL(token string, originalURL string) {
	shortURL := utils.MakeShortURL(&token)
	store.DataURL[*shortURL] = models.URL(originalURL)

	obj := models.ObjURL{
		UUID:        token,
		ShortURL:    *shortURL,
		OriginalURL: models.URL(originalURL),
	}
	f := store.NewFileRepository(config.FileStoragePath)
	f.WriteNewURL(&obj)
}
