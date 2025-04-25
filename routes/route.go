package routes

import (
	"github.com/hapfalo-mo/RestaurantUserService/handlers"
	service "github.com/hapfalo-mo/RestaurantUserService/repository"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	userService := &service.UserService{}
	userController := handlers.NewUserController(userService)

	v1 := r.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/get-all-user", userController.GetAllUser)
			// users.POST("/login", userController.Login)
			users.POST("token-login", userController.LoginToken)
			users.POST("verify-user-access", userController.IsUserVerifyAccess)
		}
	}
}
