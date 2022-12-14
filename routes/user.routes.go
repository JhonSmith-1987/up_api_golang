package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"up_api_golang/DTO"
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

func PostLoginUser(res http.ResponseWriter, req *http.Request) {
	var loginUser models.LoginUser
	json.NewDecoder(req.Body).Decode(&loginUser)
	fmt.Println("el usuario enviado es ", loginUser.Email+"password: ", loginUser.Password)
	var user models.Users
	db.DB.Where("email = ?", loginUser.Email).First(&user)
	fmt.Println(user)
	if user.ID == 0 {
		error := DTO.LoginDTO{
			Error: "Usuario no existe",
		}
		json.NewEncoder(res).Encode(&error)
		return
	}
	if user.Password != loginUser.Password {
		error := DTO.LoginDTO{
			Error: "Contraseña incorrecta",
		}
		json.NewEncoder(res).Encode(&error)
		return
	}
	response := DTO.LoginDTO{
		Respose: user,
	}
	json.NewEncoder(res).Encode(&response)
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
