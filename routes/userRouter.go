package routes

import (
	controller "github.com/arniemutasa/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("users/", controller.GetAll())
	incomingRoutes.GET("users/user_id", controller.GetUser())
	incomingRoutes.POST("users/create", controller.CreateUser())
}
