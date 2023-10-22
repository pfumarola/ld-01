package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
)

func UsersRoutes(group *gin.RouterGroup) {
	users := new(controllers.UsersController)

	group.GET("/", users.FindAll)

	group.GET("/:id", users.FindOneByID)

	group.PUT("/:id", users.Update)

	group.DELETE("/:id", users.DeleteOneByID)

	group.POST("/", users.Save)

}
