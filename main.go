package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inititalize middleware to route requests based on URL path to functions
	r := mux.NewRouter()
	// Register Hello function at '/' route
	r.HandleFunc("/", Hello)

	// Inject router into http server
	http.Handle("/", r)

	// Start the http server
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Hello prints message to the command line
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
