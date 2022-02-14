package routes

import (
	controller "go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("users/", controller.GetUsers())
	incomingRoutes.GET("users/user_id", controller.GetUser())
	incomingRoutes.POST("users/create", controller.CreateUser())
}
