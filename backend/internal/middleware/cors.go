package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/svenskhalsovard/api/internal/config"
)

// CORS creates a middleware for handling Cross-Origin Resource Sharing
func CORS(config config.CORSConfig) func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   config.AllowedMethods,
		AllowedHeaders:   config.AllowedHeaders,
		AllowCredentials: true,
		MaxAge:           config.MaxAge,
	})
}