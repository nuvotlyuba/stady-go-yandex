package utils

import (
	"fmt"
	"math/rand"

	"github.com/nuvotlyuba/study-go-yandex/config"
	"github.com/nuvotlyuba/study-go-yandex/internal/models"
)

func MakeToken(length int) *string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	str := string(b)
	return &str
}

func MakeShortURL(token *string) *models.URL {
	return models.URL(fmt.Sprintf("%s/%s", config.BaseURL, *token)).Point()
}
