package core

import (
	"math/rand"
	"time"
)

// GetArray creates an Array of random integers of size n
// Where n is supplied by the requestor.
func GetArray(size int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	a := make([]int, size)

	for i := 0; i < size; i++ {
		a[i] = rand.Int()
	}

	return a
}
