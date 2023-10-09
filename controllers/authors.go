package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorsController struct{}

func (p *AuthorsController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *AuthorsController) FindOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *AuthorsController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *AuthorsController) DeleteOneByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *AuthorsController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}
