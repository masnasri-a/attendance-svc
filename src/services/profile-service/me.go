package profileservice

import (
	"attendance-svc/src/config"
	"attendance-svc/src/models"
	"attendance-svc/src/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func MeService(c *gin.Context) {
	// Get Data from cache
	cacheData, isExist := utils.GetCacheData(c, true)
	if isExist {
		c.JSON(200, gin.H{
			"message": "Success",
			"data":    cacheData,
		})
		return
	}
	// Token
	authHeader := c.GetHeader("Authorization")
	userId, _ := utils.VerifyToken(authHeader)

	// Get user profile
	userProfile, err := getUserProfile(userId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	userProfile.Password = ""
	utils.SetCacheData(c, userProfile, 15*time.Minute, true)
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    userProfile,
	})
}

func getUserProfile(userId string) (*models.UserModelDB, error) {
	conn, err := config.GetMongoClient()
	if err != nil {
		return nil, err
	}
	var responseData models.UserModelDB
	filter := bson.D{{Key: "_id", Value: userId}}
	coll := conn.Collection("users")
	coll.FindOne(context.TODO(), filter).Decode(&responseData)
	return &responseData, nil
}
