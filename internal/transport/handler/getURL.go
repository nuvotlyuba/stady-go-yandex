package handler

import (
	"io"
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
)

func (h Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.String()[1:]
		shortURL := utils.MakeShortURL(&id)
		longURL, err := h.service.GetURL(shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", string(*longURL))
		w.WriteHeader(http.StatusPermanentRedirect)
	}

	if r.Method == http.MethodPost {
		// contentType := r.Header.Get("Content-Type")
		// if strings.Contains(contentType, "text/plain") {
		// 	http.Error(w, "Неверный тип данных", http.StatusBadRequest)
		// 	return
		// }
		bytesData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
			return
		}
		strData := string(bytesData)

		shortURL, err := h.service.CreateURL(models.URL(strData).Point())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, string(*shortURL))
	}
}
