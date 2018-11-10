package main

import (
	"flag"
	"log"
	"oop/core"
	"oop/rocks"
	"oop/services"
	"time"
)

func report(algorithm, elapsed string, size int) {
	log.Printf("%s Sort took %s\n\tSize: %d\n\n", algorithm, elapsed, *size)
}

// Citation: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
func main() {
	size := flag.Int("--size", 1000, "Size of the test set.")
	flag.Parse()

	a := core.GetArray(*size)
	start := time.Now()
	services.MergeSort(a)
	elapsed := time.Since(start)
	report("Sequential Merge Sort", elapsed.String(), *size)

	b := core.GetArray(*size)
	parallelMergeSortStart := time.Now()
	services.MergeSortParallel(b, 50)
	parallelMergeSortElapsed := time.Since(parallelMergeSortStart)
	report("Parallel Merge Sort", parallelMergeSortElapsed.String(), *size)

	c := core.GetMatrix(*size, *size, true)
	squareMatrixStart := time.Now()
	services.SquareMatrix(c)
	squareMatrixEnd := time.Since(squareMatrixStart)
	report("Sequential Square Matrix", squareMatrixEnd.String(), *size)

	d := core.GetMatrix(*size, *size, true)
	squareMatrixParStart := time.Now()
	services.SquareMatrixParallel(d)
	squareMatrixParEnd := time.Since(squareMatrixParStart)
	report("Parallel Square Matrix", squareMatrixParEnd.String(), *size)

	rocksMergeSortStart := time.Now()
	rocks.Greet()
	rocksMergeSortElapsed := time.Since(rocksMergeSortStart)
	report("Distributed Merge Sort", rocksMergeSortElapsed.String(), *size)
}
