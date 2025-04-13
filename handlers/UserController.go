package handlers

import (
	"net/http"
	"restaurantuserservice/custom"
	errorList "restaurantuserservice/error"
	"restaurantuserservice/interfaces"
	dto "restaurantuserservice/models/dto"
	service "restaurantuserservice/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service interfaces.UserInterface
}

func NewUserController(service interfaces.UserInterface) *UserController {
	if service == nil {
		panic("User Service does not work!")
	}
	return &UserController{service: service}
}
func (u *UserController) GetAllUser(c *gin.Context) {
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service does not work!"})
		return
	}
	result, err := u.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
}

// func (u *UserController) Login(c *gin.Context) {
// 	var request *dto.LoginRequest
// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, custom.Error{
// 			Message:    errorList.LoginInvalidJSONError,
// 			ErrorField: err,
// 			Field:      "Handler Layer - Body",
// 		})
// 		return
// 	}
// 	if u.service == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service does not work"})
// 		return
// 	}
// 	ok, err := u.service.Login(request)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in something"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": ok})
// }

func (u *UserController) LoginToken(c *gin.Context) {
	var request *dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom.Error{
			Message:    errorList.LoginInvalidJSONError.Error(),
			ErrorField: err.Error(),
			Field:      "Handler Layer - Body",
		})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, custom.Error{
			Message:    errorList.ServiceError.Error(),
			ErrorField: "",
			Field:      "Handler - Service",
		})
		return
	}
	ok, errorResponse := u.service.LoginToken(request)
	statusCode := http.StatusInternalServerError
	if errorResponse.Field != "" {
		if errorResponse.Message == errorList.InvalidPhoneNumber.Error() {
			statusCode = http.StatusBadRequest
		} else if errorResponse.Message == errorList.TokenGenerateError.Error() {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, errorResponse)
		return
	}
	loginResp := ok.Data
	c.SetCookie("token", loginResp.TokenString, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"data": ok})
}

func (u *UserController) IsUserVerifyAccess(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		var response = custom.Error{
			Message:    errorList.ErrCreatingToken.Error(),
			ErrorField: "Empty AuthHeader",
			Field:      "Authen-Token",
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		var response2 = custom.Error{
			Message:    errorList.ErrInvalidToken.Error(),
			ErrorField: "Invalid Token",
			Field:      "Authen-Token",
		}
		c.JSON(http.StatusUnauthorized, response2)
		c.Abort()
		return
	}
	if u.service == nil {
		var response3 = custom.Error{
			Message:    errorList.ServiceError.Error(),
			ErrorField: "nil service",
			Field:      "Authen-Token",
		}
		c.JSON(http.StatusInternalServerError, response3)
		return
	}
	_, err := u.service.IsAcceptUserAccess(tokenString)
	if err != nil {
		var response4 = custom.Error{
			Message:    errorList.ErrCreatingToken.Error(),
			ErrorField: err.Error(),
			Field:      "Authen-Token",
		}
		c.JSON(http.StatusUnauthorized, response4)
		c.Abort()
		return
	}
	userToken, err := service.ParseToken(tokenString)
	data, err2 := u.service.GetUserByUserId(userToken.Id)
	if data.Data == (dto.UserResponse{}) {
		c.JSON(http.StatusNotFound, err2)
		return
	}
	type VerifyResponse struct {
		IsVerify bool        `json:"isVerify"`
		Data     interface{} `json:"data"`
	}
	c.JSON(http.StatusOK, VerifyResponse{
		IsVerify: true,
		Data:     data,
	})
}
