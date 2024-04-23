package attendanceservice

import (
	"attendance-svc/src/config"
	"attendance-svc/src/models"
	"attendance-svc/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Clocks(c *gin.Context) {
	//checking jwt token
	var inputModel models.AttendanceModelInput
	c.ShouldBindJSON(&inputModel)
	fmt.Println("Input Model: ", inputModel)

	if err := validator.New().Struct(inputModel); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	// Token
	authHeader := c.GetHeader("Authorization")
	fmt.Println("Auth Header: ", authHeader)
	userId, err := utils.VerifyToken(authHeader)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	conn, err := config.GetMongoClient()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// Get the collection
	collection := conn.Collection("attendances")

	rawData := models.AttendanceModelDB{
		ID:             utils.GenerateUUID(),
		UserId:         userId,
		AttendanceType: inputModel.AttendanceType,
		Longitute:      inputModel.Longitute,
		Latitude:       inputModel.Latitude,
		CreatedAt:      utils.GetTimeNow(),
	}

	_, err = collection.InsertOne(c, rawData)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": inputModel.AttendanceType + " Success",
	})

}
