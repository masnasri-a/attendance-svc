package routes

import (
	"attendance-svc/src/middleware"
	profileservice "attendance-svc/src/services/profile-service"

	"github.com/gin-gonic/gin"
)

// ProfileRouter is a function that defines the routes for the profile service
func ProfileRouter(router *gin.Engine) {
	// Your code here

	// var cacheStrategy = cache.CacheByRequestURI(cacheConfig.Store, cacheConfig.DefaultCacheTime,
	// cache,
	// )

	profileRouter := router.Group("/v1/profile")
	{
		profileRouter.Use(middleware.AuthMiddleware())
		profileRouter.GET("/me",
			profileservice.MeService)
	}

}
