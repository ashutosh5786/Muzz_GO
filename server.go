package main

import (
	"fmt"
	"net/http"
)

func getjobs(w http.ResponseWriter, r *http.Request) {
	// Placeholder for job retrieval logic
	fmt.Fprintf(w, "Job retrieval logic goes here")
}

func postjobs(w http.ResponseWriter, r *http.Request) {
	// Placeholder for job posting logic
	fmt.Fprintf(w, "Job posting logic goes here")
}



func main() {
	fmt.Printf("Starting Server on http://localhost:8080\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
