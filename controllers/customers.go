package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomersController struct{}

func (p *CustomersController) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
