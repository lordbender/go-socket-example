package main

import (
	"flag"
	"log"
	"oop/core"
	"oop/distributed"
	"oop/services"
	"time"
)

func report(algorithm, elapsed string, size int) {
	log.Printf("%s Sort took %s\n\tSize: %d\n\n", algorithm, elapsed, size)
}

// Citation: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
func main() {
	size := flag.Int("size", 1000, "Size of the test set.")
	all := flag.Bool("all", false, "If true, all reports will run")
	mergeSortReports := flag.Bool("nlogn", false, "If true, all merge sort actions are triggered, report run.")
	squaredReports := flag.Bool("nsquared", false, "If true, all merge sort actions are triggered, report run.")
	rocks := flag.Bool("rocks", false, "If true, run distributed memory programs.")

	flag.Parse()

	if (*all || *mergeSortReports) && !*rocks {
		a := core.GetArray(*size)
		start := time.Now()
		services.MergeSort(a)
		elapsed := time.Since(start)
		report("Sequential Merge Sort", elapsed.String(), *size)
	}

	if (*all || *mergeSortReports) && !*rocks {
		b := core.GetArray(*size)
		parallelMergeSortStart := time.Now()
		services.MergeSortParallel(b, 50)
		parallelMergeSortElapsed := time.Since(parallelMergeSortStart)
		report("Parallel Merge Sort", parallelMergeSortElapsed.String(), *size)
	}

	if (*all || *squaredReports) && !*rocks {
		c := core.GetMatrix(*size, *size, true)
		squareMatrixStart := time.Now()
		services.SquareMatrix(c)
		squareMatrixEnd := time.Since(squareMatrixStart)
		report("Sequential Square Matrix", squareMatrixEnd.String(), *size)
	}

	if (*all || *squaredReports) && !*rocks {
		d := core.GetMatrix(*size, *size, true)
		squareMatrixParStart := time.Now()
		services.SquareMatrixParallel(d)
		squareMatrixParEnd := time.Since(squareMatrixParStart)
		report("Parallel Square Matrix", squareMatrixParEnd.String(), *size)
	}

	if (*all || *mergeSortReports) && *rocks {
		rocksMergeSortStart := time.Now()
		distributed.Greet()
		rocksMergeSortElapsed := time.Since(rocksMergeSortStart)
		report("Distributed Merge Sort", rocksMergeSortElapsed.String(), *size)
	}
}
