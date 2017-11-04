package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MakeHandler(handlers.HomeHandler))
	http.HandleFunc("/home", handlers.MakeHandler(handlers.HomeHandler))
	http.HandleFunc("/contact", handlers.MakeHandler(handlers.ContactHandler))
	http.HandleFunc("/blog", handlers.MakeHandler(handlers.BlogHandler))
	http.HandleFunc("/projects", handlers.MakeHandler(handlers.ProjectHandler))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	log.Printf("main: starting HTTP server at port: 8081")
	http.HandleFunc("/404.html", handlers.NotFoundHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
