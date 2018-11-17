package main

/* Al useful imports */
import (
	"encoding/json"
	"log"
	"net/http"
	"oop/distributed"
)

func squareVector(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t distributed.VectorModel
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.RowIndex)
}

func main() {
	http.HandleFunc("/api/v1/vector-square", squareVector)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
