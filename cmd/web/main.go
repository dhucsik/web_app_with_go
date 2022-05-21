package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhucsik/web_app_with_go/pkg/config"
	"github.com/dhucsik/web_app_with_go/pkg/handlers"
	"github.com/dhucsik/web_app_with_go/pkg/render"
)

const portNumber = ":8080"

//main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
