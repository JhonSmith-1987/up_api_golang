package routes

import "net/http"

func GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("get users"))
}
func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("get user"))
}
func PostUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("post"))
}
func DeleteUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete"))
}
