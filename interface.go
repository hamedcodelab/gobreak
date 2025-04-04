package gobreak

import (
	"sync"
	"time"
)

type Breaker interface {
	Execute(func() error) error
}

// NAME: breaker
// This struct includes:
// mu: A mutex to ensure thread-safe access to the circuit breaker.
// state: The current state of the circuit breaker (Closed, Open, or HalfOpen).
// failureCount: The current count of consecutive(successive) failures.
// failureThresholdAllow: The number of consecutive failures allowed before opening the circuit.
// recoveryTimeToHalfOpen: The cool-down period before the circuit breaker transitions from Open to HalfOpen.
// halfOpenMaxRequests: The maximum number of successful requests needed to close the circuit.

/*
// lastFailureTime: The timestamp of the last failure.
// halfOpenSuccessCount: The number of successful requests in the HalfOpen state.
// timeout: The maximum duration to wait for a request to complete.
*/
type breaker struct {
	mu                  sync.Mutex
	state               State
	failureCount        int
	SuccessRequestCount int
	//lastFailureTime      time.Time
	//halfOpenSuccessCount int
	failureThresholdAllow  int
	recoveryTimeToHalfOpen time.Duration
	timeToHalfOpen         time.Time
	halfOpenMaxRequests    int
	//timeout              time.Duration
}

func NewBreaker(opts ...Option) Breaker {
	b := &breaker{
		state: Closed,
	}
	for _, opt := range opts {
		opt(b)
	}

	if b.failureCount == 0 {
		b.failureCount = 1
	}
	if b.recoveryTimeToHalfOpen == 0 {
		b.recoveryTimeToHalfOpen = 10 * time.Second
	}
	if b.halfOpenMaxRequests == 0 {
		b.halfOpenMaxRequests = 1
	}
	if b.failureThresholdAllow == 0 {
	}
	return b
}
