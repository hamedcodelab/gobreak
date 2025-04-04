package gobreak

import (
	"sync"
	"time"
)

type Breaker interface {
}

// NAME: breaker
// This struct includes:
// mu: A mutex to ensure thread-safe access to the circuit breaker.
// state: The current state of the circuit breaker (Closed, Open, or HalfOpen).
// failureCount: The current count of consecutive(successive) failures.
// lastFailureTime: The timestamp of the last failure.
// halfOpenSuccessCount: The number of successful requests in the HalfOpen state.
// failureThreshold: The number of consecutive failures allowed before opening the circuit.
// recoveryTime: The cool-down period before the circuit breaker transitions from Open to HalfOpen.
// halfOpenMaxRequests: The maximum number of successful requests needed to close the circuit.
// timeout: The maximum duration to wait for a request to complete.

type breaker struct {
	mu                   sync.Mutex
	state                State
	failureCount         int
	lastFailureTime      time.Time
	halfOpenSuccessCount int
	failureThreshold     int
	recoveryTime         time.Duration
	halfOpenMaxRequests  int
	timeout              time.Duration
}

func NewBreaker(opts ...Option) Breaker {
	b := &breaker{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}
