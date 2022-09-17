package sync

import (
	"sync"
	"time"
)

// WaitGroup A WaitGroup waits for a collection of goroutines to finish and support wait timeout.
type WaitGroup struct {
	sync.WaitGroup
}

// WaitWithTimeout blocks until the WaitGroup counter is zero or timeout
func (wg *WaitGroup) WaitWithTimeout(d time.Duration) (timeout bool) {
	select {
	case <-wait(&wg.WaitGroup):
		return false
	case <-time.After(d):
		return true
	}
}

func wait(wg *sync.WaitGroup) <-chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}
