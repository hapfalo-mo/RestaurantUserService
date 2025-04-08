package handlers

import (
	"net/http"
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

func (u *UserController) Login(c *gin.Context) {
	var request *dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service does not work"})
		return
	}
	ok, err := u.service.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in something"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ok})
}

func (u *UserController) LoginToken(c *gin.Context) {
	var request *dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ok, err := u.service.LoginToken(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
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
