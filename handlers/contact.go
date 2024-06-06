package handlers

import (
	"fmt"
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

func (h *ContactHandler) Submit(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	message := r.FormValue("message")
	email := r.FormValue("email")
	fmt.Println(name, lastName, message, email)
	c := views.SuccessfulMessage()

	err := c.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering Contact template", http.StatusInternalServerError)
	}
}
