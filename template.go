package main

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
		return
	}
}
