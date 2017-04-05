package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZakeryFyke/Mapster/Mapster/views"
)

//NewUser returns a User.
func NewUser() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

//Users contains a pointer to a view
type Users struct {
	NewView *views.View
}

//SignupForm contains schema for email and password
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//New renders a view for a User object
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Create fills the form and applies it to a user.
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}
