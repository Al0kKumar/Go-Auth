package main

import (
	"log"
	"net/http"

	"github.com/Al0kKumar/Go-Auth.git/db"
	"github.com/Al0kKumar/Go-Auth.git/models"
	"github.com/Al0kKumar/Go-Auth.git/routes"
	"github.com/gorilla/mux"
)

func main() {
    
	// connect to database
	db.Connect()

	db.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()
	routes.Authroutes(r)

	log.Println("server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000",r))


}
