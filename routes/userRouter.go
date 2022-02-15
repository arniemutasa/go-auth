package routes

import (
	controller "github.com/arniemutasa/go-auth/controllers"
	"github.com/arniemutasa/go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users/", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
	// incomingRoutes.POST("users/create", controller.CreateUser())
}
