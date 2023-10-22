package server

import (
	"github.com/gin-gonic/gin"

	"github.com/pfumarola/ld-01/controllers"
	"github.com/pfumarola/ld-01/middlewares"
	"github.com/pfumarola/ld-01/server/routes"
)

func DefineRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	loginController := controllers.LoginController{}
	router.POST("/login", loginController.Run)

	router.Use(middlewares.AuthMiddleware)

	authorsGroup := router.Group("authors")
	routes.AuthorsRoutes(authorsGroup)

	booksGroup := router.Group("books")
	routes.BooksRoutes(booksGroup)

	usersGroup := router.Group("users")
	routes.UsersRoutes(usersGroup)

	loansGroup := router.Group(("loans"))
	routes.LoansRoutes(loansGroup)

}
