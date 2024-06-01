package handlers

import (
	"net/http"
	"portfolio/models"
	"portfolio/views"
)

type ProjectsHandler struct{}

func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{}
}

func (h *ProjectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.Projects()
	err := views.Layout(c, "Amirreza Ashouri", models.PROJECTS_TAB).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering Projects template", http.StatusInternalServerError)
	}
}
