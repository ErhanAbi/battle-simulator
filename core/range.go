package core

import "math/rand"

// Range returns a random number between the given range
func Range(min, max float64) float64 {
	r := rand.Float64()
	v := (max-min)*r + min
	return v
}
