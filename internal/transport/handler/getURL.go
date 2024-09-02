package handler

import (
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
)

func (h Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.String()[1:]
		shortURL := utils.MakeShortURL(&id)
		originalURL, err := h.service.GetURL(shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", string(*originalURL))
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
