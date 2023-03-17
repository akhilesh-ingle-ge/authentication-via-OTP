package main

import (
	"net/http"

	"github.com/akhilesh-ge/jwt-email/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome")
	})

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/signin", controller.SignIn)
		userRoutes.POST("/verify", controller.VerifyOTP)
	}

	router.Run(":8080")
}
