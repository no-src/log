package crand

import (
	"math/rand"
	"sync"
)

// CRand A CRand is a source of random numbers and supports some concurrency function
type CRand struct {
	*rand.Rand

	mu sync.Mutex
}

// New returns a new CRand that uses random values from src to generate other random values
func New(src rand.Source) *CRand {
	return &CRand{
		Rand: rand.New(src),
	}
}

// CFloat64 support concurrency for Float64 function
func (cr *CRand) CFloat64() float64 {
	defer cr.mu.Unlock()
	cr.mu.Lock()
	return cr.Rand.Float64()
}
