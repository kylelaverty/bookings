package handlers

import (
	"fmt"
	"net/http"

	"github.com/kylelaverty/bookings/pkg/config"
	"github.com/kylelaverty/bookings/pkg/models"
	"github.com/kylelaverty/bookings/pkg/render"
)

// Repo is the repository used by the handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	fmt.Println(m.App.Session.GetString(r.Context(), "remote_ip"))
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateDate{})
}

// About is the about page handler.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic.
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	fmt.Println(m.App.Session.Get(r.Context(), "remote_ip"))

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Send data to the template.
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateDate{
		StringMap: stringMap,
	})
}
