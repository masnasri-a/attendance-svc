package main

import (
	"attendance-svc/src/routes"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		println("Request method: ", c.Request.URL.Path, " Status: ", c.Writer.Status())
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())

	routes.AuthRouter(router)

	router.Run(":9091")
}
