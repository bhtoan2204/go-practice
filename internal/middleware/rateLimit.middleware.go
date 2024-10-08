package middleware

import (
	"simple_bank/package/response"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// client stores the rate limiter and the last time the client was seen.
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// rateLimiter holds the map of clients and the mutex to synchronize access.
type rateLimiter struct {
	clients map[string]*client
	mu      sync.Mutex
	// Configuration for rate limiter
	requestsPerSecond rate.Limit
	burstSize         int
}

// newRateLimiter initializes a new rateLimiter.
func NewRateLimiter(r rate.Limit, b int) *rateLimiter {
	rl := &rateLimiter{
		clients:           make(map[string]*client),
		requestsPerSecond: r,
		burstSize:         b,
	}

	// Start a background goroutine to remove old clients.
	go rl.cleanupClients()

	return rl
}

// getClient retrieves the rate limiter for a specific IP address.
func (rl *rateLimiter) getClient(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Check if the client already exists.
	if c, exists := rl.clients[ip]; exists {
		c.lastSeen = time.Now()
		return c.limiter
	}

	// Create a new rate limiter for the new client.
	limiter := rate.NewLimiter(rl.requestsPerSecond, rl.burstSize)
	rl.clients[ip] = &client{
		limiter:  limiter,
		lastSeen: time.Now(),
	}
	return limiter
}

// cleanupClients removes clients that haven't been seen for over 3 minutes.
func (rl *rateLimiter) cleanupClients() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, c := range rl.clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimitMiddleware is the Gin middleware that enforces rate limiting.
func RateLimitMiddleware(rl *rateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the client's IP address.
		ip := c.ClientIP()

		limiter := rl.getClient(ip)

		// Allow or reject the request based on the rate limiter.
		if !limiter.Allow() {
			response.ErrorTooManyRequestsResponse(c, 429)
			c.Abort()
			return
		}

		c.Next()
	}
}
