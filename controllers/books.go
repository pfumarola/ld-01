package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BooksController struct{}

func (p *BooksController) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
