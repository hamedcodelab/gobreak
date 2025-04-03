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
// failureCount: The current count of consecutive failures.
// lastFailureTime: The timestamp of the last failure.
// halfOpenSuccessCount: The number of successful requests in the HalfOpen state.
// failureThreshold: The number of consecutive failures allowed before opening the circuit.
// recoveryTime: The cool-down period before the circuit breaker transitions from Open to HalfOpen.
// halfOpenMaxRequests: The maximum number of successful requests needed to close the circuit.
// timeout: The maximum duration to wait for a request to complete.

type breaker struct {
	mu                   sync.Mutex    // Guards the circuit breaker state
	state                string        // Current state of the circuit breaker
	failureCount         int           // Number of consecutive failures
	lastFailureTime      time.Time     // Time of the last failure
	halfOpenSuccessCount int           // Successful requests in half-open state
	failureThreshold     int           // Failures to trigger open state
	recoveryTime         time.Duration // Wait time before half-open
	halfOpenMaxRequests  int           // Requests allowed in half-open state
	timeout              time.Duration // Timeout for requests
}
