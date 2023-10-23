package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
	"github.com/pfumarola/ld-01/middlewares"
)

func UsersRoutes(group *gin.RouterGroup) {
	users := new(controllers.UsersController)

	group.Use(middlewares.AdminMiddleware)

	group.GET("/", users.FindAll)

	group.GET("/:id", users.FindOneByID)

	group.PUT("/:id", users.Update)

	group.DELETE("/:id", users.DeleteOneByID)

	group.POST("/", users.Save)

}
