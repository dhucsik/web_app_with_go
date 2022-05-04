package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

//main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
