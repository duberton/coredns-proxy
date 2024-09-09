package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("lovely folks")
	_, _ = fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/path", handler) // Handle all HTTP requests

	// Listen and serve on port 8080
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Println(fmt.Sprintf("Starting HTTP server on %s", port))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
