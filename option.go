package gobreak

import "time"

type Option func(*breaker)

// Option to set failure threshold
func WithFailureThreshold(threshold int) Option {
	return func(b *breaker) {
		b.failureThresholdAllow = threshold
	}
}

// Option to set recovery time
func WithRecoveryTime(duration time.Duration) Option {
	return func(b *breaker) {
		b.recoveryTimeToHalfOpen = duration
	}
}

// Option to set max half-open requests
func WithHalfOpenMaxRequests(max int) Option {
	return func(b *breaker) {
		b.halfOpenMaxRequests = max
	}
}
