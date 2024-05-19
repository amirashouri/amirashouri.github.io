package handlers

import (
	"net/http"
	"portfolio/views"
)

type ProjectsHandler struct{}

func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{}
}

func (h *ProjectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.Projects()
	err := views.Layout(c, "Amirreza Ashouri").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering Projects template", http.StatusInternalServerError)
	}
}
