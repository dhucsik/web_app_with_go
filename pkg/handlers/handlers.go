package handlers

import (
	"fmt"
	"net/http"

	"github.com/dhucsik/web_app_with_go/pkg/render"
)

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func Ts() {
	fmt.Println("the fuck is that")
}
