package readyset

import "time"

// ReadySet is the basic timer type of readyset.go
type ReadySet *time.Time

// Go starts a new timer
func Go() ReadySet {
	return &time.Now()
}

// Lap time as a Duration
func (r ReadySet) Lap() time.Duration {
	return time(*r).Sub(time.Now())
}

// Stop the timer and return the integer value rounded to the required precision (e.g. time.Millisecond)
func (r ReadySet) Stop(precision time.Duration) int64 {
	return int(r.Lap() / precision)
}
