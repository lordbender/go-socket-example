package services

import (
	"fmt"
	"oop/core"
	"sync"
)

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

// SquareMatrixParallel takes a matrix and returns
// a matrix of the same size, but with every
// index squared.
//
// bit overflow possible when squaring 32bit integers: be careful.
func SquareMatrixParallel(a [][]int) [][]int {
	size := len(a)
	c := core.GetMatrix(size, size, false)

	respond := make(chan []int, size)
	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go squareIt(respond, &wg, a[i])
	}

	wg.Wait()
	close(respond)

	for queryResp := range respond {
		fmt.Printf("Got Response:\t %s\n", queryResp)
	}
	return c
}

func squareIt(respond chan<- []int, wg *sync.WaitGroup, a []int) {
	defer wg.Done()
	size := len(a)
	c := make([]int, size)
	for j := 0; j < size; j++ {
		c[j] = a[j] ^ 2
	}
	respond <- c
}
