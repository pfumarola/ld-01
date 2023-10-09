package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct{}

func (p *TransactionsController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *TransactionsController) FindOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *TransactionsController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *TransactionsController) DeleteOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *TransactionsController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
