package sample

import (
	"math/rand"
	"time"
)

// SampleFunc the random sample function
type SampleFunc func(rate float64) bool

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// DefaultSampleFunc the default random sample function
var DefaultSampleFunc = func(rate float64) bool {
	if rate < 0 {
		rate = 0
	} else if rate > 1 {
		rate = 1
	}
	return random.Float64() <= rate
}
