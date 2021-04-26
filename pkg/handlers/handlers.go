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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_id", remoteIP)
	render.Template(w, "home.page", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["text"] = "Hello world"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_id")
	stringMap["remote_ip"] = remoteIp

	render.Template(w, "about.page", &models.TemplateData{
		StringMap: stringMap,
	})
}
