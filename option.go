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

/*
// Option to set timeout
func WithTimeout(duration time.Duration) Option {
	return func(b *breaker) {
		b.timeout = duration
	}
}

*/

// Option to set initial state
func WithInitialState(state State) Option {
	return func(b *breaker) {
		b.state = state
	}
}

/*
// Option to set last failure time (useful for testing or resetting)
func WithLastFailureTime(t time.Time) Option {
	return func(b *breaker) {
		b.lastFailureTime = t
	}
}

*/

// Option to set initial failure count (useful for restoring state)
func WithFailureCount(count int) Option {
	return func(b *breaker) {
		b.failureCount = count
	}
}

/*
// Option to set half-open success count
func WithHalfOpenSuccessCount(count int) Option {
	return func(b *breaker) {
		b.halfOpenSuccessCount = count
	}
}

*/
