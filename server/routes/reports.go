package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
)

func ReportsRoutes(group *gin.RouterGroup) {
	reports := new(controllers.ReportsController)

	group.GET("/", reports.FindAll)

	group.GET("/:id", reports.FindOneByID)

	group.PUT("/:id", reports.Update)

	group.DELETE("/:id", reports.DeleteOneByID)

	group.POST("/", reports.Save)

}
