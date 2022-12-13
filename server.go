package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"up_api_golang/db"
	"up_api_golang/models"
	"up_api_golang/routes"
)

func main() {

	db.DBConnection()

	err := db.DB.AutoMigrate(models.Users{})
	if err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	router.HandleFunc("/", routes.HomeHandler)
	api.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	api.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	api.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	api.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)

}
