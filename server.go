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
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getjobs(w, r)
		} else if r.Method == http.MethodPost {
			postjobs(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	

	fmt.Printf("Starting Server on http://localhost:8080\n")
	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
