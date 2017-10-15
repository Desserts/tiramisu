package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "title",
		"Content": "aaaa",
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}
