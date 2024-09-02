package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		Debug("httpRequest",
			zap.String("uri", r.URL.Path),
			zap.String("method", r.Method),
		)

		next.ServeHTTP(&lw, r)
		duration := time.Since(start)

		Debug("httpResponse",
			zap.Int("status", responseData.status),
			zap.Int("size", responseData.size),
			zap.Duration("duration", duration),
		)
	})
}
