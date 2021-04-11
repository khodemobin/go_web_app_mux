package main

import (
	"fmt"
	"github.com/khodemobin/go_web_app_mux/pkg/config"
	"github.com/khodemobin/go_web_app_mux/pkg/handlers"
	"github.com/khodemobin/go_web_app_mux/pkg/render"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	var app config.AppConfig
	app.UseCache = false

	registerTemplates(&app)
	registerRepositories(&app)
	registerRoutes(&app)
}

func registerTemplates(app *config.AppConfig) {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache", err)
	}
	app.TemplateCache = tc
	render.NewTemplates(app)
}

func registerRepositories(app *config.AppConfig) {
	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)
}

func registerRoutes(app *config.AppConfig) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routes(app),
	}

	fmt.Println("Listen to port ", port)
	err := srv.ListenAndServe()
	log.Fatalln(err)
}
