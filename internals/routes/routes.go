package routes

import (
	controller "github.com/FeelDat/go-testtask/internals/auth"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
