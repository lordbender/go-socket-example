package main

import (
	"log"
	"oop/core"
	"oop/services"
	"time"
)

// Citation: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
func main() {
	a := core.GetArray(100000000)

	start := time.Now()
	services.MergeSort(a)
	elapsed := time.Since(start)
	log.Printf("Sequential Merge Sort took %s", elapsed)
}
