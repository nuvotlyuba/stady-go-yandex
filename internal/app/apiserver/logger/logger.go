package logger

import (
	"net/http"

	"go.uber.org/zap"
)

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size = size

	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func Info(msg string, field ...interface{})  { zap.S().Debugw(msg, field...) }
func Debug(msg string, field ...interface{}) { zap.S().Debugw(msg, field...) }
func Fatal(msg string, field ...interface{}) { zap.S().Fatalw(msg, field...) }
