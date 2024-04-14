package authservice

import (
	"attendance-svc/src/config"
	"attendance-svc/src/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateWorkspaceService is a function that handles the create workspace service
func CreateWorkspaceService(c *gin.Context) {
	var workspaceModel models.WorkspaceModelInput
	c.ShouldBindJSON(&workspaceModel)
	if err := validator.New().Struct(workspaceModel); err != nil {
		println(err.Error())
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	roles := map[string]int{
		"trial":      7,
		"basic":      30,
		"premium":    90,
		"enterprise": 365,
	}

	ExpiredAt := time.Now().AddDate(0, 0, roles[workspaceModel.Subscription]).Unix()

	insertedData := models.WorkspaceModelDB{
		ID:            primitive.NewObjectID().Hex(),
		Name:          workspaceModel.Name,
		CreatorUserId: workspaceModel.CreatorUserId,
		MaxUsers:      workspaceModel.MaxUsers,
		Subscription:  workspaceModel.Subscription,
		ExpiredAt:     ExpiredAt,
		CreatedAt:     time.Now().Unix(),
	}
	coll, err := config.GetMongoClient()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	_, err = coll.Collection("workspaces").InsertOne(c, insertedData)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	responseData := models.WorkspaceModelResponse{
		Name:          insertedData.Name,
		CreatorUserId: insertedData.CreatorUserId,
		Subscription:  insertedData.Subscription,
		MaxUsers:      insertedData.MaxUsers,
	}
	c.JSON(200, gin.H{
		"message": "Workspace created",
		"data":    responseData,
	})
}
