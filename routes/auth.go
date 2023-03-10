package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(UserRepository)
	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}
