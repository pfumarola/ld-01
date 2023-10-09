package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
)

func AuthorsRoutes(group *gin.RouterGroup) {
	authors := new(controllers.AuthorsController)

	group.GET("/", authors.FindAll)

	group.GET("/:id", authors.FindOneByID)

	group.PUT("/:id", authors.Update)

	group.DELETE("/:id", authors.DeleteOneByID)

	group.POST("/", authors.Save)

}
