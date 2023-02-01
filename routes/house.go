package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func HouseRoute(r *mux.Router) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)
	r.HandleFunc("/houses", h.FindHouses).Methods("GET")
	r.HandleFunc("/house/{id}", h.GetHouse).Methods("GET")
	r.HandleFunc("/house", middleware.Auth(middleware.UploadFile(h.CreateHouse))).Methods("POST")
	r.HandleFunc("/house/{id}", h.UpdateHouse).Methods("PATCH")
	r.HandleFunc("/house/{id}", h.DeleteHouse).Methods("DELETE")
}
