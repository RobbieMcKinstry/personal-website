package handlers

import (
	"net/http"
)

// The IndexHandler points to the root of the application.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl", nil)
}

func ResearchHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Research", "research.md")
}

func EngineeringHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Engineering", "engineering.md")
}

func TeachingHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Teaching", "teaching.md")
}

func AboutMeHandler(w http.ResponseWriter, r *http.Request) {
	var title = "About Me"
	renderMarkdown(w, r, title, garbage)
}

var garbage = `
# About me

**My name:** Robbie McKinstry

__Likes:__` + "`coding`" + `

dislikes: shitty web frameworks`
