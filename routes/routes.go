package routes

import (
	"github.com/Al0kKumar/Go-Auth.git/controllers"
	"github.com/gorilla/mux"
)

func Authroutes(r *mux.Router){

	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login",controllers.Login).Methods("POST")
	
}