package handlers

import (
	"net/http"
	"portfolio/views"
)

type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.About()
	err := views.Layout(c, "Amirreza Ashouri").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering About template", http.StatusInternalServerError)
	}
}
