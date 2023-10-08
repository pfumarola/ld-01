package server

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	app := gin.Default()
	DefineRoutes(app)
	app.Run()
}
