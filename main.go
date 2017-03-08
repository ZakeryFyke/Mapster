package main

import (
	"net/http"

	"github.com/RyanMcBerg/Mapster/controllers"
	"github.com/gorilla/mux"
)

func main() {
	staticController := controllers.NewStatic()

	r := mux.NewRouter()

	r.Handle("/home", staticController.Home).Methods("Get")
	http.ListenAndServe(":3000", r)

}
