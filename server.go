package main

/* Al useful imports */
import (
	"encoding/json"
	"log"
	"net/http"
	"oop/distributed"
)

func squareVector(rw http.ResponseWriter, req *http.Request) {

	// Decode the POST body
	decoder := json.NewDecoder(req.Body)
	var reqModel distributed.VectorModel
	err := decoder.Decode(&reqModel)
	if err != nil {
		panic(err)
	}

	// Square the Vector
	size := len(reqModel.Vector)

	c := make([]int, size)
	for i := 0; i < size; i++ {
		val := reqModel.Vector[i]
		c[i] = val * val
	}

	// Set the index to align the response object
	// Create the response model
	resModel := distributed.VectorModel{reqModel.RowIndex, c}

	// Parse the response model into a writable format
	resString, err := json.Marshal(resModel)

	rw.Write(resString)
}

func main() {
	http.HandleFunc("/api/v1/vector-square", squareVector)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
