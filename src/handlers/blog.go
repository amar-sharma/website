package handlers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Blog struct {
	Link   string
	Title  string
	Posted string
	Desc   string
}

type Page struct {
	Blogs []Blog
}

func BlogHandler(w http.ResponseWriter, r *http.Request, title string) {
	filename := "./blogs.json"
	blogLists, err := ioutil.ReadFile(filename)
	var blogs []Blog
	json.Unmarshal(blogLists, &blogs)
	var templates = template.Must(
		template.ParseFiles("tmpl/home.html",
			"tmpl/blog.html"))
	templates.ExecuteTemplate(w, "blog.html", &Page{Blogs: blogs})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
