package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhucsik/web_app_with_go/pkg/handlers"
)

const portNumber = ":8080"

//main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
