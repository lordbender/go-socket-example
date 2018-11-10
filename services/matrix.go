package services

import "oop/core"

// SquareMatrix takes a matrix and returns
// a matrix of the same size, but with every
// index squared.
//
// bit overflow possible when squaring 32bit integers: be careful.
func SquareMatrix(a [][]int) [][]int {
	size := len(a)
	c := core.GetMatrix(size, size, false)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c[i][j] = a[i][j] ^ 2
		}
	}

	return c
}
