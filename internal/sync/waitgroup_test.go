package sync

import (
	"testing"
	"time"
)

func TestWaitWithTimeout(t *testing.T) {
	testCases := []struct {
		name    string
		count   int
		d       time.Duration
		timeout bool
	}{
		{"Done", 3, time.Second * 10, false},
		{"Timeout", 3, time.Second / 2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wg := WaitGroup{}
			wg.Add(tc.count)
			for i := 0; i < tc.count; i++ {
				go func() {
					time.Sleep(time.Second)
					wg.Done()
				}()
			}
			actual := wg.WaitWithTimeout(tc.d)
			if actual != tc.timeout {
				t.Errorf("TestWaitWithTimeout error, expect timeout:%v but actual timeout:%v", tc.timeout, actual)
			}
		})
	}
}
