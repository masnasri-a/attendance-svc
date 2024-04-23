package routes

import (
	attendanceservice "attendance-svc/src/services/attendance-service"

	"github.com/gin-gonic/gin"
)

func AttendanceProfile(router *gin.Engine) {
	attendanceRouter := router.Group("/v1/attendance")
	{
		attendanceRouter.POST("/clock", attendanceservice.Clocks)
	}

}
