package interfaces

import (
	custom "restaurantuserservice/custom"
	models "restaurantuserservice/models"
	dto "restaurantuserservice/models/dto"
)

type UserInterface interface {
	// Login(request *dto.LoginRequest) (response dto.LoginResponse, err error)
	LoginToken(request *dto.LoginRequest) (response custom.Data[dto.LoginResponse], err custom.Error)
	GetAllUser() (result []models.User, err error)
	IsAcceptUserAccess(tokenString string) (response bool, err error)
}
