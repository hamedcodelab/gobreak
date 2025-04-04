package gobreak

import (
	"sync"
	"time"
)

type Breaker interface {
	Execute(func() error)
}

// NAME: breaker
// This struct includes:
// mu: A mutex to ensure thread-safe access to the circuit breaker.
// state: The current state of the circuit breaker (Closed, Open, or HalfOpen).
// failureCount: The current count of consecutive(successive) failures.
// failureThresholdAllow: The number of consecutive failures allowed before opening and half opening the circuit.
// recoveryTimeToHalfOpen: The cool-down period before the circuit breaker transitions from Open to HalfOpen.
// halfOpenMaxRequests: The maximum number of successful requests needed to close the circuit.

type breaker struct {
	mu                     sync.Mutex
	state                  State
	failureCount           int
	SuccessRequestCount    int
	failureThresholdAllow  int
	recoveryTimeToHalfOpen time.Duration
	timeToHalfOpen         time.Time
	halfOpenMaxRequests    int
}

func NewBreaker(opts ...Option) Breaker {
	b := &breaker{
		failureCount:           0,
		recoveryTimeToHalfOpen: 10 * time.Second,
		halfOpenMaxRequests:    1,
		failureThresholdAllow:  3,
		state:                  Closed,
		timeToHalfOpen:         time.Now(),
	}

	for _, opt := range opts {
		opt(b)
	}
	return b
}
