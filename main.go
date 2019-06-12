package main


import (
	"fmt"
	api "restapi"
	"os"
	"github.com/gorilla/mux"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage: ./restapi <address>")
	}

	domain := os.Args[1]

	router := mux.NewRouter()

	dataset = append(api.dataset, api.Data{ID: "1", Body: "First Post Body"})

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
