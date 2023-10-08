package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportsController struct{}

func (p *ReportsController) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
