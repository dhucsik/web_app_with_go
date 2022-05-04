package handlers

import (
	"net/http"
)

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.html")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.html")
}
