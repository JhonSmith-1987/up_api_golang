package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"up_api_golang/db"
	"up_api_golang/models"
)

func GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	var users []models.Users
	db.DB.Find(&users)
	json.NewEncoder(res).Encode(&users)
}

func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	var user models.Users
	params := mux.Vars(req)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		res.WriteHeader(http.StatusNotFound) // 404
		res.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(res).Encode(&user)
}

func PostUserHandler(res http.ResponseWriter, req *http.Request) {
	var user models.Users
	json.NewDecoder(req.Body).Decode(&user)
	fmt.Println("nueva conexion de usuario", user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		res.WriteHeader(http.StatusBadRequest) // 400
		res.Write([]byte(err.Error()))
	}
	json.NewEncoder(res).Encode(createdUser)
}

func DeleteUsersHandler(res http.ResponseWriter, req *http.Request) {
	var user models.Users
	params := mux.Vars(req)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		res.WriteHeader(http.StatusNotFound) // 404
		res.Write([]byte("User not found"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	res.WriteHeader(http.StatusOK)
}
