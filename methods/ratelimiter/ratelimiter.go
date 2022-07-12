package ratelimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter_ interface {
	SetLimit()
	UpdateLimit()
	DeleteLimit()
	IsLimitReached()
}

type limit struct {
	lastCall time.Time
	*rate.Limiter
}

type RateLimiter struct {
	mu          *sync.Mutex
	rate, burst int
	limit       map[string]*limit
}

func New(rate, burst int) *RateLimiter {
	return &RateLimiter{
		&sync.Mutex{},
		rate, burst,
		make(map[string]*limit),
	}
}

func (r *RateLimiter) SetLimit(ip string, rate, burst int) {
	r.mu.Lock()
	_, ok := r.limit[ip]
	if ok {
	}
}

func (r *RateLimiter) UpdateLimit(ip string, rate, burst int) {
	r.mu.Lock()
}

func (r *RateLimiter) DeleteLimit(ip string) {
	rate.Limiter{}.Allow()
	r.mu.Lock()
}

func (r *RateLimiter) IsLimitReached(ip string) bool {
	r.mu.Lock()
	return false
}
