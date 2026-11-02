package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	return &Limiter{mu: sync.Mutex{}, clock: clock, rate: ratePerSec, burst: burst, tokens: float64(burst)}
}
func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.clock == nil || l.burst <= 0 || n <= 0 || n > l.burst {
		return false
	}
	now := l.clock.Now()
	last := l.last
	var diff time.Duration
	if now.After(last) {
		diff = now.Sub(last)
		l.last = now
	}
	l.tokens = min(float64(l.burst), l.tokens+diff.Seconds()*max(l.rate, 0))
	if float64(n) <= l.tokens {
		l.tokens = l.tokens - float64(n)
		return true
	}
	return false
}
