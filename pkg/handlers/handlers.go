package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/khodemobin/go_web_app_mux/pkg/config"
	"github.com/khodemobin/go_web_app_mux/pkg/models"
	"github.com/khodemobin/go_web_app_mux/pkg/render"
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
	render.Template(w, r, "home.page", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["text"] = "Hello world"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_id")
	stringMap["remote_ip"] = remoteIp

	render.Template(w, r, "about.page", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "majors.page", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "make-reservation.page", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Header().Set("Content-type", "application/json")
	_, err := w.Write([]byte(fmt.Sprintf("start is %s and end is %s", start, end)))
	if err != nil {
		log.Fatal(err)
	}
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostAvailabilityJson(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	res := jsonResponse{
		OK:      true,
		Message: start + "-test" + end,
	}

	out, err := json.Marshal(res)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(out))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page", &models.TemplateData{})
}
