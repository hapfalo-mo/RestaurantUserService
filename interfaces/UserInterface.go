package interfaces

import (
	models "restaurantuserservice/models"
	dto "restaurantuserservice/models/dto"
)

type UserInterface interface {
	Login(request *dto.LoginRequest) (response dto.LoginResponse, err error)
	LoginToken(request *dto.LoginRequest) (token string, err error)
	GetAllUser() (result []models.User, err error)
	IsAcceptUserAccess(tokenString string) (response bool, err error)
}
