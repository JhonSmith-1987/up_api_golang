package DTO

import "up_api_golang/models"

type LoginDTO struct {
	Error   string       `json:"error"`
	Respose models.Users `json:"respose"`
}
