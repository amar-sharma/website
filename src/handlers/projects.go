package handlers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Project struct {
	Link  string
	Title string
	Icon  string
	Desc  string
}

type Projects struct {
	Projects []Project
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request, title string) {
	filename := "./projects.json"
	projList, err := ioutil.ReadFile(filename)
	var projs []Project
	json.Unmarshal(projList, &projs)
	var templates = template.Must(
		template.ParseFiles("tmpl/home.html",
			"tmpl/projects.html"))
	templates.ExecuteTemplate(w, "projects.html", &Projects{Projects: projs})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
