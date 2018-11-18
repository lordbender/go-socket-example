package distributed

import (
	"sync"
)

// Merge merges left and right slice into newly created slice
func Merge(model MergeRequestModel) MergeResponseModel {
	left := model.Left
	right := model.Right
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	response := MergeResponseModel{slice}
	return response
}

// MergeSortDistributed algorithm with deferment.
// Reference https://github.com/duffleit/golang-parallel-mergesort/blob/master/mergesort/mergesort.go
func MergeSortDistributed(list []int, threshold int, hostsfile bool) []int {

	useThreshold := !(threshold < 0)

	size := len(list)
	middle := size / 2

	if size <= 1 {
		return list
	}

	var left, right []int

	sortInNewRoutine := size > threshold && useThreshold

	if !sortInNewRoutine {
		left = MergeSortDistributed(list[:middle], threshold, hostsfile)
		right = MergeSortDistributed(list[middle:], threshold, hostsfile)
	} else {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer func() { wg.Done() }()
			left = MergeSortDistributed(list[:middle], threshold, hostsfile)

		}()

		go func() {
			defer func() { wg.Done() }()
			right = MergeSortDistributed(list[middle:], threshold, hostsfile)
		}()

		wg.Wait()
	}

	// Distribute here
	request := MergeRequestModel{left, right}
	response := Merge(request)

	return response.Merged
}

// func merge(respond chan<- mergeModel, a mergeModel, url string) {
// 	size := len(a.row)
// 	c := make([]int, size)

// 	for j := 0; j < size; j++ {
// 		jsonData := VectorModel{a.position, a.row}
// 		jsonValue, _ := json.Marshal(jsonData)

// 		response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
// 		if err != nil {
// 			fmt.Printf("The HTTP request failed with error %s\n", err)
// 		} else {
// 			data, _ := ioutil.ReadAll(response.Body)
// 			// Marshal the results into an object
// 			var obj VectorModel
// 			if err := json.Unmarshal(data, &obj); err != nil {
// 				panic(err)
// 			}
// 			c = obj.Vector
// 		}
// 	}
// 	respond <- rowHelper{row: c, position: a.position}
// }
