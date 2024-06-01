package handlers

import (
	"net/http"
	"portfolio/models"
	"portfolio/views"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.Contact()
	err := views.Layout(c, "Amirreza Ashouri", models.CONTACT_TAB).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering Contact template", http.StatusInternalServerError)
	}
}
