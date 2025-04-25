package interfaces

import (
	custom "github.com/hapfalo-mo/RestaurantUserService/custom"
	models "github.com/hapfalo-mo/RestaurantUserService/models"
	dto "github.com/hapfalo-mo/RestaurantUserService/models/dto"
)

type UserInterface interface {
	// Login(request *dto.LoginRequest) (response dto.LoginResponse, err error)
	LoginToken(request *dto.LoginRequest) (response custom.Data[dto.LoginResponse], err custom.Error)
	GetAllUser() (result []models.User, err error)
	IsAcceptUserAccess(tokenString string) (response bool, err error)
	GetUserByUserId(id int) (response custom.Data[dto.UserResponse], err custom.Error)
}
