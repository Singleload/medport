package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

// Logging is a middleware that logs each request
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			// Get request ID from context
			requestID := middleware.GetReqID(r.Context())

			// Log request details
			log.Info().
				Str("request_id", requestID).
				Str("method", r.Method).
				Str("url", r.URL.Path).
				Str("query", r.URL.RawQuery).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Int("status", ww.Status()).
				Int("bytes_written", ww.BytesWritten()).
				Dur("duration", time.Since(start)).
				Msg("Request processed")
		}()

		next.ServeHTTP(ww, r)
	})
}