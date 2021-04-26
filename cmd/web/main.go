package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/khodemobin/go_web_app_mux/pkg/config"
	"github.com/khodemobin/go_web_app_mux/pkg/handlers"
	"github.com/khodemobin/go_web_app_mux/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = "8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.UseCache = false
	app.InProduction = false

	registerSession(&app)
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

func registerSession(app *config.AppConfig){
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
}