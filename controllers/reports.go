package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportsController struct{}

func (p *ReportsController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *ReportsController) FindOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *ReportsController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *ReportsController) DeleteOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *ReportsController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
