package main

import "oop/core"
import "oop/services"

// Citation: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
func main() {

	a := core.GetArray(120)
	b := core.GetArray(1000000)

	services.MergeSort(a)
	services.MergeSort(b)
}
