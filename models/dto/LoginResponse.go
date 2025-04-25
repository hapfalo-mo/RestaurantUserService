package dto

import (
	models "github.com/hapfalo-mo/RestaurantUserService/models"
)

type LoginResponse struct {
	Data        models.User
	TokenString string `json:"token"`
}
