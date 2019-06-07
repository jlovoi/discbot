package main

import (
	"os"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
)

// Data struct with just an id and a body tag
// This assumes that we will be encoding our data using json
type Data struct {
	ID string `json:"id"`
	Body string `json:"body"`
}

var dataset []Data


func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataset)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parameters := mux.Vars(r)

	for _, item := range dataset {
		if item.ID == parameters["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// I believe this should "return" an empty Data struct since this will only be reached
	// in the case of the specified data not being found
	json.NewEncoder(w).Encode(&Data{})
}


func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(data)
	data.ID = strconv.Itoa(rand.Intn(100000))
	dataset = append(dataset, data)
	json.NewEncoder(w).Encode(&data)
}


func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range dataset {
		if item.ID == parameters["id"] {
			// Remove the old version of this item so we can update it
			dataset = append(dataset[:index], dataset[index+1:]...)

			var data Data
			_ = json.NewDecoder(r.Body).Decode(&data)
			data.ID = parameters["id"]
			dataset = append(dataset, data)
			break
		}
	}
	// Otherwise, there is nothing to update... just "return" the OG dataset
	json.NewEncoder(w).Encode(dataset)
}


func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range dataset {
		if item.ID == parameters["id"] {
			// We will delete by just creating a new slice containing everything but item
			dataset = append(dataset[:index], dataset[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(dataset)
}


func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage: ./restapi <address>")
	}

	domain := os.Args[1]

	router := mux.NewRouter()

	dataset = append(dataset, Data{ID: "1", Body: "First Post Body"})

	// Here are our 5 basic endpoints for the data.
	// If an incoming request URL matches one of the paths, the corresponding handler
	// is called passing (http.ResponseWriter, *http.Request) as parameters.
	router.HandleFunc("/dataset", getPosts).Methods("GET")
	router.HandleFunc("/dataset", createPost).Methods("POST")
	router.HandleFunc("/dataset/{id}", getPost).Methods("GET")
	router.HandleFunc("/dataset/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/dataset/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(domain, router)

}