package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/dto"
)

// Recovery is a middleware that recovers from panics and logs the error
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Get request ID from context
				requestID := middleware.GetReqID(r.Context())

				// Log error with stack trace
				stack := debug.Stack()
				log.Error().
					Str("request_id", requestID).
					Str("method", r.Method).
					Str("url", r.URL.Path).
					Interface("error", err).
					Bytes("stack", stack).
					Msg("Recovered from panic")

				// Convert error to string
				var errorMsg string
				switch v := err.(type) {
				case string:
					errorMsg = v
				case error:
					errorMsg = v.Error()
				default:
					errorMsg = fmt.Sprintf("%v", v)
				}

				// Respond with error
				response, _ := dto.NewErrorResponse(
					dto.ErrorCodeInternalServerError,
					"Internal server error",
					nil,
				).ToJSON()

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(response)
			}
		}()

		next.ServeHTTP(w, r)
	})
}