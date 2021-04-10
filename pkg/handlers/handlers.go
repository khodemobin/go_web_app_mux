package handlers

import (
	"github.com/khodemobin/go_web_app_mux/pkg/config"
	"github.com/khodemobin/go_web_app_mux/pkg/models"
	"github.com/khodemobin/go_web_app_mux/pkg/render"
	"net/http"
)



var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello world"

	render.Template(w, "about.page", &models.TemplateData{
		StringMap: stringMap,
	})
}
