package authservice

import (
	"attendance-svc/src/config"
	"attendance-svc/src/models"
	"attendance-svc/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser is a function that handles the create user service
func CreateUser(c *gin.Context) {
	var userModel models.UserModelInput
	c.ShouldBindJSON(&userModel)
	if err := validator.New().Struct(userModel); err != nil {
		println(err.Error())
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	coll, err := config.GetMongoClient()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	isUserExist := coll.Collection("users").FindOne(c, map[string]string{"email": userModel.Email})
	if isUserExist.Err() == nil {
		c.JSON(400, gin.H{
			"message": "User already exist",
		})
		return
	}

	password := utils.HashPassword(userModel.Password)

	insertedData := models.UserModelDB{
		ID:        primitive.NewObjectID().Hex(),
		Email:     userModel.Email,
		FullName:  userModel.FullName,
		Password:  password,
		CreatedAt: time.Now().Unix(),
	}

	_, err = coll.Collection("users").InsertOne(c, insertedData)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	responseData := models.UserModelResponse{
		Email:    insertedData.Email,
		FullName: userModel.FullName,
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"data":    responseData,
	})
}
