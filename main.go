package main

import (
	"net/http"

	"github.com/RyanMcBerg/Mapster/controllers"
	"github.com/gorilla/mux"
)

func main() {

	// c, err := maps.NewClient(maps.WithAPIKey("AIzaSyAGGxruyyZKhj9fzWk-hTDohDsU8cfIi3s"))
	// if err != nil {
	// 	panic(err)
	// }
	//
	// mappy := &maps.DirectionsRequest{
	// 	Origin:      "Lubbock",
	// 	Destination: "Dallas",
	// }
	//
	// resp, _, err := c.Directions(context.Background(), mappy)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(resp)

	staticController := controllers.NewStatic()

	r := mux.NewRouter()

	r.Handle("/", staticController.Home).Methods("Get")
	r.Handle("/home", staticController.Home).Methods("Get")
	http.ListenAndServe(":3000", r)

}
