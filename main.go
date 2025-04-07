package main

import (
	"restaurantuserservice/db"
	routers "restaurantuserservice/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	// Initialize Router
	router := gin.Default()

	// CORS Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	//Xử lý OPTIONS request để tránh bị block bởi preflight request
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204) // No Content
	})
	routers.SetUpRoutes(router)
	router.Run(":1703")
}
