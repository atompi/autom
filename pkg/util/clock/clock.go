package clock

import "time"

// PassiveClock allows for injecting fake or real clocks into code
// that needs to read the current time but does not support scheduling
// activity in the future.
type PassiveClock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

// Clock allows for injecting fake or real clocks into code that
// needs to do arbitrary things based on time.
type Clock interface {
	PassiveClock
	// After returns the channel of a new Timer.
	// This method does not allow to free/GC the backing timer before it fires. Use
	// NewTimer instead.
	After(d time.Duration) <-chan time.Time
	// NewTimer returns a new Timer.
	NewTimer(d time.Duration) Timer
	// Sleep sleeps for the provided duration d.
	// Consider making the sleep interruptible by using 'select' on a context channel and a timer channel.
	Sleep(d time.Duration)
	// Tick returns the channel of a new Ticker.
	// This method does not allow to free/GC the backing ticker. Use
	// NewTicker from WithTicker instead.
	Tick(d time.Duration) <-chan time.Time
}

// Timer allows for injecting fake or real timers into code that
// needs to do arbitrary things based on time.
type Timer interface {
	C() <-chan time.Time
	Stop() bool
	Reset(d time.Duration) bool
}
