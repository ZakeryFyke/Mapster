package controllers

import "github.com/ZakeryFyke/Mapster/Mapster/views"

//Controller for all of our static pages, which will likely be most of our pages.

// NewStatic controller  for all of our static pages, which will likely be most of our pages.
func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"),
	}
}

// Static Struct to hold our static views, add to this if we have more static pages.
type Static struct {
	Home *views.View
}
