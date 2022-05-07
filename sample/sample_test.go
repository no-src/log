package sample

import (
	"testing"
)

func TestDefaultSampleFunc(t *testing.T) {
	testCases := []struct {
		name  string
		total int
		rate  float64
	}{
		{"sample rate -1", 1000, -1},
		{"sample rate 0", 1000, 0},
		{"sample rate 0.1", 1000, 0.1},
		{"sample rate 0.2", 1000, 0.2},
		{"sample rate 0.3", 1000, 0.3},
		{"sample rate 0.4", 1000, 0.4},
		{"sample rate 0.5", 1000, 0.5},
		{"sample rate 0.6", 1000, 0.6},
		{"sample rate 0.7", 1000, 0.7},
		{"sample rate 0.8", 1000, 0.8},
		{"sample rate 0.9", 1000, 0.9},
		{"sample rate 1", 1000, 1},
		{"sample rate 2", 1000, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			y := 0
			n := 0
			for i := 0; i < tc.total; i++ {
				if DefaultSampleFunc(tc.rate) {
					y++
				} else {
					n++
				}
			}
			t.Logf("[DefaultSampleFunc] total:%d rate:%.2f hit rate:%.2f hit:%d miss:%d", tc.total, tc.rate, float64(y)/float64(tc.total), y, n)
		})
	}
}
