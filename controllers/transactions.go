package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct{}

func (p *TransactionsController) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
