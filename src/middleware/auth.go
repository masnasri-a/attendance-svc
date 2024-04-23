package middleware

import (
	"attendance-svc/src/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// Your code here
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		if authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		println("Auth Header: ", authHeader)

		_, err := utils.VerifyToken(authHeader)
		if err != nil {
			c.JSON(403, gin.H{
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
