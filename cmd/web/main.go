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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Fatal(srv.ListenAndServe())
}
