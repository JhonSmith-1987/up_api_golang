package routes

import (
	"encoding/json"
	"net/http"
	"up_api_golang/db"
	"up_api_golang/models"
)

func GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("get users"))
}

func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("get user"))
}

func PostUserHandler(res http.ResponseWriter, req *http.Request) {
	var user models.Users
	json.NewDecoder(req.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		res.WriteHeader(http.StatusBadRequest) // 400
		res.Write([]byte(err.Error()))
	}
	json.NewEncoder(res).Encode(createdUser)
}

func DeleteUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete"))
}
