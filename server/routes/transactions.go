package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
	"github.com/pfumarola/ld-01/middlewares"
)

func LoansRoutes(group *gin.RouterGroup) {
	loans := new(controllers.LoansController)

	group.GET("/", loans.FindAll)

	group.GET("/:id", loans.FindOneByID)

	group.POST("/", loans.Save)

	group.Use(middlewares.AdminMiddleware)

	group.PUT("/:id", loans.Update)

	group.DELETE("/:id", loans.DeleteOneByID)

}
