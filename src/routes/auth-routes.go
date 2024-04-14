package routes

import (
	authservice "attendance-svc/src/services/auth-service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	authRouter := router.Group("/v1/auth")
	{
		authRouter.POST("/login", authservice.LoginService)
		authRouter.POST("/create-workspace", authservice.CreateWorkspaceService)
		authRouter.POST("/register", authservice.CreateUser)
	}
}
