package authservice

import (
	"attendance-svc/src/config"
	"attendance-svc/src/models"
	"attendance-svc/src/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// LoginService is a function that handles the login service
func LoginService(c *gin.Context) {
	var loginModel models.LoginModelInput
	c.ShouldBindJSON(&loginModel)
	if err := validator.New().Struct(loginModel); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	// Check if user exists
	conn, err := config.GetMongoClient()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	var responseData models.UserModelDB
	coll := conn.Collection("users")
	user := coll.FindOne(context.TODO(), map[string]string{"email": loginModel.Email})
	if err := user.Decode(&responseData); err != nil {
		c.JSON(400, gin.H{
			"message": "User not found",
		})
		return
	}

	// Compare password
	if !utils.ComparePassword(responseData.Password, loginModel.Password) {
		c.JSON(400, gin.H{
			"message": "Invalid password",
		})
		return
	}

	responseData.Password = ""

	// Generate JWT token
	token, err := utils.CreateToken(responseData.ID)
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": "Internal server error",
	// 	})
	// 	return
	// }

	// Return JWT token
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
		"data":    responseData,
	})
}
