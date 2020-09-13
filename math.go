package pgo

import (
	"math/rand"
	"time"
)

// Rand returns a pseudo-random integer between min and max based on unix-nano time seed
// !! for random numbers suitable for security-sensitive work, use the crypto/rand package instead
func Rand(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
