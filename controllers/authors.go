package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
)

type AuthorsController struct{}

func (p *AuthorsController) FindAll(c *gin.Context) {
	var authors []models.Author
	database.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
	return
}

func (p *AuthorsController) FindOneByID(c *gin.Context) {
	var author models.Author
	id := c.Param("id")
	database.DB.Find(&author, id)
	c.JSON(http.StatusOK, author)
	return
}

func (p *AuthorsController) Update(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *AuthorsController) DeleteOneByID(c *gin.Context) {
	var author models.Author
	id := c.Param("id")
	database.DB.Delete(&author, id)
	c.Status(http.StatusOK)
	return
}

func (p *AuthorsController) Save(c *gin.Context) {
	var author models.Author
	c.Bind(&author)
	result := database.DB.Create(&author)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
