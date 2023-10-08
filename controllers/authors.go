package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorsController struct{}

func (p *AuthorsController) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
