package gobreak

import "time"

func (b *breaker) Execute(fn func() error) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	err := fn()
	if err != nil {
		b.failureCount++
		b.handleError()
		return err
	}

	b.SuccessRequestCount++
	if b.state == HalfOpen && b.SuccessRequestCount >= b.halfOpenMaxRequests {
		b.state = Closed
		b.SuccessRequestCount = 0
		b.failureCount = 0
	}
	return nil
}

func (b *breaker) handleError() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.state == Closed && b.failureCount >= b.failureThresholdAllow {
		b.state = HalfOpen
		b.failureCount = 0
		return
	}
	if b.state == HalfOpen && b.failureCount >= b.failureThresholdAllow {
		b.state = Open
		b.failureCount = 0
		return
	}

	if b.state == Open && time.Now().After(b.timeToHalfOpen.Add(b.recoveryTimeToHalfOpen)) {
		b.timeToHalfOpen = time.Now()
		b.state = HalfOpen
	}
}
