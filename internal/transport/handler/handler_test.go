package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/repository"
	"github.com/nuvotlyuba/study-go-yandex/internal/service"
	"github.com/nuvotlyuba/study-go-yandex/internal/types"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostURL(t *testing.T) {
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
			repo := repository.NewVarRepository()
			s := service.New(repo)
			h := New(s)
			h.PostURL(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "Отличный от %s  Conent-Type", tt.want.contentType)
			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)
			assert.NotEmpty(t, string(body), "Тело ответа пустое")
		})
	}
}

func TestGetURL(t *testing.T) {
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
			repo := repository.NewVarRepository()
			s := service.New(repo)
			h := New(s)
			h.GetURL(w, r)
			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "Отличный от %s  Content-Type", tt.want.contentType)
			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)
		})
	}
}

func TestPostJSONURL(t *testing.T) {
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
			repo := repository.NewVarRepository()
			s := service.New(repo)
			h := New(s)
			h.PostJSONURL(w, r)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "Отличный от %s  Content-Type", tt.want.contentType)
			assert.Equal(t, tt.want.statusCode, res.StatusCode, "Отличный от %d статус код", tt.want.statusCode)

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err, "Ошибка чтения тела ответа")
			err = res.Body.Close()
			require.NoError(t, err)
			assert.NotEmpty(t, string(body), "Тело ответа пустое")
		})
	}
}

func EnsureNewURL(token string, longURL string) {
	shortURL := utils.MakeShortURL(&token)
	repository.DataURL[*shortURL] = models.URL(longURL)
}
