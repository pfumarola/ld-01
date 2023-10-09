package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
)

func CustomersRoutes(group *gin.RouterGroup) {
	customers := new(controllers.CustomersController)

	group.GET("/", customers.FindAll)

	group.GET("/:id", customers.FindOneByID)

	group.PUT("/:id", customers.Update)

	group.DELETE("/:id", customers.DeleteOneByID)

	group.POST("/", customers.Save)

}
