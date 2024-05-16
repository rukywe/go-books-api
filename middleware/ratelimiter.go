package middleware

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type rateLimiter struct {
    limiter  *rate.Limiter
    lastSeen time.Time
}

var clients = make(map[string]*rateLimiter)
var mu sync.Mutex

func getClient(ip string) *rateLimiter {
    mu.Lock()
    defer mu.Unlock()

    if limiter, exists := clients[ip]; exists {
        limiter.lastSeen = time.Now()
        return limiter
    }

    limiter := &rateLimiter{
        limiter:  rate.NewLimiter(1, 5), // 1 request per second with a burst of 5
        lastSeen: time.Now(),
    }
    clients[ip] = limiter
    return limiter
}

func cleanupClients() {
    for {
        time.Sleep(time.Minute)
        mu.Lock()
        for ip, limiter := range clients {
            if time.Since(limiter.lastSeen) > 3*time.Minute {
                delete(clients, ip)
            }
        }
        mu.Unlock()
    }
}

func RateLimit(next http.Handler) http.Handler {
    go cleanupClients()

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        client := getClient(ip)

        if !client.limiter.Allow() {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
