package distributed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SquareMatrix will return a copy of a matrix, with every index's value squared.
func SquareMatrix(matrix [][]int, hostsfile bool) {

	// Get the POST Body together
	b := []int{1, 2, 3, 4, 5}
	jsonData := VectorModel{1, b}
	jsonValue, _ := json.Marshal(jsonData)

	// Call the server(s) to multiply the vectors
	response, err := http.Post("http://localhost:8080/api/v1/vector-square", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("Hum, did it work => %s\n\n", string(data))
	}
}
