package main

import (
	"fmt"
	"github.com/khodemobin/go_web_app_mux/pkg/config"
	"github.com/khodemobin/go_web_app_mux/pkg/handlers"
	"github.com/khodemobin/go_web_app_mux/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	// register app templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	// register handler repositories
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)

	fmt.Println("Listen to port 8000")
	_ = http.ListenAndServe(":8000", nil)
}
