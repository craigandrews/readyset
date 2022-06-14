package readyset

import (
	"sync"
	"time"
)

// ReadySet is the basic timer type of readyset.go
type ReadySet struct {
	start time.Time
	lap   time.Time
	mux   sync.Mutex
}

// Go starts a new timer
func Go() *ReadySet {
	now := time.Now()
	return &ReadySet{
		start: now,
		lap:   now,
	}
}

// Lap time as a Duration
func (r *ReadySet) Lap() time.Duration {
	r.mux.Lock()
	defer r.mux.Unlock()

	prevLap := r.lap
	r.lap = time.Now()
	return r.lap.Sub(prevLap)
}

// Stop the timer and return the total elapsed time as a Duration
func (r *ReadySet) Stop() time.Duration {
	return time.Now().Sub(r.start)
}
