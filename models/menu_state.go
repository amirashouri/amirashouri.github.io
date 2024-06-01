package models

const (
	HOME_TAB     = "home"
	ABOUT_TAB    = "about"
	CONTACT_TAB  = "contact"
	PROJECTS_TAB = "projects"
)

type MenuState struct {
	IsOpen       bool
	SelectedPage string
}

func NewMenuState(isOpen bool, selectedState string) *MenuState {
	return &MenuState{IsOpen: isOpen, SelectedPage: selectedState}
}
