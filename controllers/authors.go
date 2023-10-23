package controllers

import (
	"net/http"
	"strconv"

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
	var author models.Author
	id, atoiError := strconv.Atoi(c.Param("id"))
	if atoiError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Id must be a number"})
		return
	}
	bindError := c.ShouldBind(&author)
	if bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid fields, malformed request"})
		return
	}
	author.AuthorID = uint(id)
	result := database.DB.Save(&author)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
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
	bindError := c.ShouldBind(&author)
	if bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid fields, malformed request"})
		return
	}
	result := database.DB.Create(&author)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
