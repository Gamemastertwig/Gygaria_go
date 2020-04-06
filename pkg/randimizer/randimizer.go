// Package randimizer is a collections of functions for randimizing groups of data
package randimizer

import (
	"math/rand"
	"time"
)

// StringRandimizer is function that takes a string slice and returns a
// random string from it.
func StringRandimizer(inputs []string) string {
	// seed rand based on time
	rand.Seed(time.Now().UnixNano())
	// get a random int from 0 to length of string slice -1
	l := len(inputs)
	n := rand.Intn(l)
	// find string at index of random number
	// return that string
	return inputs[n]
}

// IntRandimizer is function that takes a min and max integers and returns a
// random integer between those number.
func IntRandimizer(min int, max int) int {
	// seed rand based on time
	rand.Seed(time.Now().UnixNano())
	// get a random int from min to max
	n := rand.Intn(max-min) + min
	// return random int
	return n
}
