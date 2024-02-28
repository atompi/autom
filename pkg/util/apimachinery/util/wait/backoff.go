package wait

import "github.com/atompi/autom/pkg/util/clock"

// BackoffManager manages backoff with a particular scheme based on its underlying implementation.
type BackoffManager interface {
	// Backoff returns a shared clock.Timer that is Reset on every invocation. This method is not
	// safe for use from multiple threads. It returns a timer for backoff, and caller shall backoff
	// until Timer.C() drains. If the second Backoff() is called before the timer from the first
	// Backoff() call finishes, the first timer will NOT be drained and result in undetermined
	// behavior.
	Backoff() clock.Timer
}
