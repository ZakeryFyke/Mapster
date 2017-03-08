package controllers

import "github.com/RyanMcBerg/Mapster/views"

//NewStatic returns a new static
func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"),
	}
}

// Static woo
type Static struct {
	Home *views.View
}
