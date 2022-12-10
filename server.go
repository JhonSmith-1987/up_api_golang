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

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)

}
