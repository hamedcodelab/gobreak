package gobreak

import "time"

func (b *breaker) Execute(fn func() error) error {
	err := fn()
	if err != nil {
		b.handleError()
		return err
	}
	b.handleSuccess()
	return nil
}

func (b *breaker) handleError() {
	b.failureCount++
	if b.state == Closed {
		if b.failureCount >= b.failureThresholdAllow {
			b.state = HalfOpen
			b.failureCount = 0
		}
		return
	}
	if b.state == HalfOpen {
		if b.failureCount >= b.failureThresholdAllow {
			b.state = Open
			b.failureCount = 0
		}
	}
}

func (b *breaker) handleSuccess() {
	if b.state == HalfOpen {
		b.SuccessRequestCount++
		if b.SuccessRequestCount >= b.halfOpenMaxRequests {
			b.state = Closed
			b.SuccessRequestCount = 0
			b.failureCount = 0
		}
	}

	if b.state == Open {
		if b.timeToHalfOpen.Add(b.recoveryTimeToHalfOpen).After(time.Now()) {
			return
		}
		b.timeToHalfOpen = time.Now()
		b.state = HalfOpen
	}
}
