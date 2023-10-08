package server

import (
	"github.com/gin-gonic/gin"

	"github.com/pfumarola/ld-01/server/routes"
)

func DefineRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	authorsGroup := router.Group("authors")
	routes.AuthorsRoutes(authorsGroup)

	booksGroup := router.Group("books")
	routes.BooksRoutes(booksGroup)

	customersGroup := router.Group("customers")
	routes.CustomersRoutes(customersGroup)

	transactionsGroup := router.Group(("transactions"))
	routes.TransactionsRoutes(transactionsGroup)

}
