package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"

	"github.com/svenskhalsovard/api/internal/config"
	"github.com/svenskhalsovard/api/internal/dto"
)

// Client represents a client for rate limiting
type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimiterMiddleware is a middleware that limits request rates
type RateLimiterMiddleware struct {
	clients map[string]*Client
	mu      sync.Mutex
	config  config.RateLimitConfig
}

// NewRateLimiterMiddleware creates a new RateLimiterMiddleware
func NewRateLimiterMiddleware(config config.RateLimitConfig) *RateLimiterMiddleware {
	go func() {
		// Clean up old clients every minute
		for range time.Tick(time.Minute) {
			r.cleanup(5 * time.Minute)
		}
	}()

	return &RateLimiterMiddleware{
		clients: make(map[string]*Client),
		config:  config,
	}
}

// getClientLimiter gets or creates a rate limiter for a client
func (m *RateLimiterMiddleware) getClientLimiter(ip string) *rate.Limiter {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Get client if exists
	client, exists := m.clients[ip]
	if !exists {
		// Create new client
		client = &Client{
			limiter: rate.NewLimiter(
				rate.Limit(m.config.RequestsPerMinute/60),
				m.config.BurstSize),
			lastSeen: time.Now(),
		}
		m.clients[ip] = client
	} else {
		// Update last seen
		client.lastSeen = time.Now()
	}

	return client.limiter
}

// cleanup removes old clients
func (m *RateLimiterMiddleware) cleanup(maxAge time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for ip, client := range m.clients {
		if time.Since(client.lastSeen) > maxAge {
			delete(m.clients, ip)
		}
	}

	log.Debug().Int("active_clients", len(m.clients)).Msg("Rate limiter cleanup completed")
}

// Handler is the middleware handler
func (m *RateLimiterMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If rate limiting is disabled, skip
		if !m.config.Enabled {
			next.ServeHTTP(w, r)
			return
		}

		// Get client IP
		ip := r.RemoteAddr

		// Get client limiter
		limiter := m.getClientLimiter(ip)

		// Check if request is allowed
		if !limiter.Allow() {
			log.Warn().Str("ip", ip).Msg("Rate limit exceeded")

			// Return rate limit error
			response := dto.NewErrorResponse(
				"RATE_LIMIT_EXCEEDED",
				"Rate limit exceeded. Please try again later.",
				nil,
			)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write(response.ToJSON())
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RateLimiter creates a new rate limiter middleware
func RateLimiter(config config.RateLimitConfig) func(http.Handler) http.Handler {
	middleware := NewRateLimiterMiddleware(config)
	return middleware.Handler
}