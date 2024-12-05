package internal

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type WeatherHandler struct {
	logger    *Logger
	templates *template.Template
}

func NewWeatherHandler(logger *Logger) *WeatherHandler {
	templates := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))
	return &WeatherHandler{
		logger:    logger,
		templates: templates,
	}
}

func (h *WeatherHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	h.templates.ExecuteTemplate(w, "home.html", nil)
}

func (h *WeatherHandler) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	h.logger.Printf("Запрос погоды для города: %s", city)

	data := struct {
		City string
	}{
		City: city,
	}

	h.templates.ExecuteTemplate(w, "weather.html", data)
}
