package handlers

import (
	"net/http"
	"portfolio/models"
	"portfolio/views"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.Index()
	err := views.Layout(c, "Amirreza Ashouri", models.HOME_TAB).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering home template", http.StatusInternalServerError)
	}
}
