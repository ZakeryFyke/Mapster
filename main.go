package main

import (
	"net/http"

	"github.com/RyanMcBerg/Mapster/controllers"
	"github.com/gorilla/mux"
)

func main() {

	staticController := controllers.NewStatic()
	userController := controllers.NewUser()

	r := mux.NewRouter()

	r.Handle("/home",
		staticController.Home).Methods("Get")
	r.Handle("/", staticController.Home).Methods("Get")
	r.HandleFunc("/signup", userController.New).Methods("GET")
	r.HandleFunc("/signup", userController.Create).Methods("post")
	http.ListenAndServe(":3000", r)

}
