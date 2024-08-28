package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/study-go-yandex/internal/models"
	"go.uber.org/zap"
)

func (h Handler) PostJsonURL(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != string(JSON_CONTENT_TYPE) {
		http.Error(w, "Unexpected content type", http.StatusBadRequest)
		return
	}

	var jsonReq models.JsonURLRequest
	if err := json.NewDecoder(r.Body).Decode(&jsonReq); err != nil {
		logger.Debug("cannot decode request JSON body", zap.Error(err))
		http.Error(w, err.Error()+" -> unmarshal", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.CreateURL(models.URL(jsonReq.URL).Point())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", string(JSON_CONTENT_TYPE))
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	if err := enc.Encode(models.JsonURLResponse{Result: string(*shortURL)}); err != nil {
		logger.Debug("error encoding response", zap.Error(err))
		http.Error(w, err.Error()+" -> marshal", http.StatusBadRequest)
		return
	}
}
