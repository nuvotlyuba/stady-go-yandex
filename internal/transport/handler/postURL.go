package handler

import (
	"io"
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/types"
)

func (h Handler) PostURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytesData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
			return
		}
		strData := string(bytesData)

		result, err := h.service.CreateURL(models.URL(strData).Point())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", string(types.TextContentType))
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, string(result.ShortURL))
	}
}
