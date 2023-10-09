package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
)

func BooksRoutes(group *gin.RouterGroup) {
	books := new(controllers.BooksController)

	group.GET("/", books.FindAll)

	group.GET("/:id", books.FindOneByID)

	group.PUT("/:id", books.Update)

	group.DELETE("/:id", books.DeleteOneByID)

	group.POST("/", books.Save)
}
