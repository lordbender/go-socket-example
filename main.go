package main

import (
	"log"
	"oop/core"
	"oop/services"
	"os"
	"strconv"
	"time"
)

// Citation: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
func main() {
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("Invalid Command Line Arguments, expected int representing test size.")
		return
	}

	a := core.GetArray(size)
	start := time.Now()
	services.MergeSort(a)
	elapsed := time.Since(start)
	log.Printf("Sequential Merge Sort took %s", elapsed)

	b := core.GetArray(size)
	parallelMergeSortStart := time.Now()
	services.MergeSortParallel(b, 50)
	parallelMergeSortElapsed := time.Since(parallelMergeSortStart)
	log.Printf("Parallel Merge Sort took %s", parallelMergeSortElapsed)

	c := core.GetMatrix(size, size, true)
	squareMatrixStart := time.Now()
	services.SquareMatrix(c)
	squareMatrixEnd := time.Since(squareMatrixStart)
	log.Printf("Sequential Square Matrix took %s", squareMatrixEnd)
}
