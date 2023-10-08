package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/controllers"
)

func TransactionsRoutes(group *gin.RouterGroup) {
	transactions := new(controllers.TransactionsController)
	group.GET("/", transactions.All)
}
