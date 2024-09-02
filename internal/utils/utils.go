package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

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

func SwitchStorage() string {
	fmt.Println(config.FileStoragePath, "config.FileStoragePath")
	if config.FileStoragePath != "" {
		return "file"
	}

	return "mem"
}

func GetDirsFromPath(path string) string {
	sl := strings.Split(path, "/")
	sl = sl[:len(sl)-1]
	st := strings.Join(sl, "/")
	return st
}

func MakeDir(dirname string) {
	os.MkdirAll(GetDirsFromPath(dirname), 0777)
}
