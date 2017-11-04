package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func ContactHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, "contact")
}

func HomeHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, "home")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "404")
}

func BlogHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, "blog")
}

func ProjectHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, "projects")
}


func renderTemplate(w http.ResponseWriter, tmpl string) {
	var templates = template.Must(
    template.ParseFiles(
      "tmpl/home.html",
      "static/404.html",
      "tmpl/contact.html",
      "tmpl/blog.html",
      "tmpl/projects.html"))
	err := templates.ExecuteTemplate(w, tmpl+".html", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(home|contact|blog|projects|)$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(r.URL.Path)
		fmt.Println(m)
		if m == nil {
			http.Redirect(w, r, "/404.html", http.StatusFound)
			return
		}
		fn(w, r, m[1])
	}
}
