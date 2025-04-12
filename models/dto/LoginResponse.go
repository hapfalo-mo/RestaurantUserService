package dto

import (
	models "restaurantuserservice/models"
)

type LoginResponse struct {
	Data        models.User
	TokenString string `json:"token"`
}
