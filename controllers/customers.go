package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomersController struct{}

func (p *CustomersController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *CustomersController) FindOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *CustomersController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *CustomersController) DeleteOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *CustomersController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
