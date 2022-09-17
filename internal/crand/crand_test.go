package crand

import (
	"math/rand"
	"testing"
	"time"

	"github.com/no-src/log/internal/sync"
)

func TestCFloat64(t *testing.T) {
	count := 3
	timeout := time.Second
	wg := sync.WaitGroup{}
	wg.Add(count)
	cr := New(rand.NewSource(100))
	for i := 0; i < count; i++ {
		go func() {
			cr.CFloat64()
			wg.Done()
		}()
	}

	if wg.WaitWithTimeout(timeout) {
		t.Errorf("TestCFloat64 timeout for %s", timeout.String())
	}
}
