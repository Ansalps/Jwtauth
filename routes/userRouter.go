package routes

import (
	"github.com/gin-gonic/gin"
	controller "jwt authentication/controllers"
	"jwt authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
