package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BooksController struct{}

func (p *BooksController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *BooksController) FindOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *BooksController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *BooksController) DeleteOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *BooksController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
