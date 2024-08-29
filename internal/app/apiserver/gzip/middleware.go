package gzip

import (
	"net/http"
	"strings"

	"github.com/nuvotlyuba/study-go-yandex/internal/types"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ow := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		isSupportGzip := strings.Contains(acceptEncoding, types.GzipCompressType)

		currContentType := r.Header.Get("Content-Type")
		isSupportContentType := strings.Contains(currContentType, types.JSONContentType) ||
			strings.Contains(currContentType, types.HTMLContentType)

		if isSupportContentType && isSupportGzip {
			cw := NewWriter(w)
			ow = cw
			defer cw.Close()

			ow.Header().Set("Content-Encoding", types.GzipCompressType)
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		isSendGzip := strings.Contains(contentEncoding, types.GzipCompressType)

		if isSendGzip {
			cr, err := NewReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			r.Body = cr
			defer cr.Close()
		}

		next.ServeHTTP(ow, r)
	})
}
