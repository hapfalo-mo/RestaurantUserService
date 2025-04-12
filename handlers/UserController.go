package handlers

import (
	"net/http"
	"restaurantuserservice/custom"
	errorList "restaurantuserservice/error"
	"restaurantuserservice/interfaces"
	dto "restaurantuserservice/models/dto"

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
	loginResp, okCast := ok.Data.(dto.LoginResponse)
	if !okCast {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}
	c.SetCookie("token", loginResp.TokenString, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"data": ok})
}

func (u *UserController) IsUserVerifyAccess(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not verify user"})
		c.Abort()
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service does not work!"})
		return
	}
	ok, err := u.service.IsAcceptUserAccess(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ok})
}
