package distributed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const root string = "127.0.0.1"

func hosts() []string {
	return []string{
		"compute-0-2",
		"compute-0-3",
		"compute-0-4",
		"compute-0-5",
		"compute-0-6",
		"compute-0-7",
		"compute-0-8",
		"compute-0-9",
		"compute-0-10",
		"compute-0-11",
		"compute-0-12",
	}
}

type rowHelper struct {
	row      []int
	position int
}

// SquareMatrix will return a copy of a matrix, with every index's value squared.
func SquareMatrix(a [][]int, hostsfile bool) {
	size := len(a[0])

	// Here is where the result will go
	result := make([][]int, size)

	burstLimit := 50
	respond := make(chan rowHelper, burstLimit)

	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go squareIt(respond, &wg, rowHelper{row: a[i], position: i})
	}

	wg.Wait()
	close(respond)
	for queryResp := range respond {
		result[queryResp.position] = queryResp.row
	}
	fmt.Print(result)
}

func squareIt(respond chan<- rowHelper, wg *sync.WaitGroup, a rowHelper) {
	defer wg.Done()
	size := len(a.row)
	c := make([]int, size)

	for j := 0; j < size; j++ {
		jsonData := VectorModel{a.position, a.row}
		jsonValue, _ := json.Marshal(jsonData)

		response, err := http.Post("http://localhost:8080/api/v1/vector-square", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			// Marshal the results into an object
			var obj VectorModel
			if err := json.Unmarshal(data, &obj); err != nil {
				panic(err)
			}
			c = obj.Vector
		}
	}
	respond <- rowHelper{row: c, position: a.position}
}
