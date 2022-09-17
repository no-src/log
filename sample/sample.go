package sample

import (
	"math/rand"
	"time"

	"github.com/no-src/log/internal/crand"
)

// SampleFunc the random sample function
type SampleFunc func(rate float64) bool

var random = crand.New(rand.NewSource(time.Now().UnixNano()))

// DefaultSampleFunc the default random sample function
var DefaultSampleFunc = func(rate float64) bool {
	if rate < 0 {
		rate = 0
	} else if rate > 1 {
		rate = 1
	}
	return random.CFloat64() <= rate
}
