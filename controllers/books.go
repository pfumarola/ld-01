package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
)

type BooksController struct{}

func (p *BooksController) FindAll(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
	return
}

func (p *BooksController) FindOneByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	database.DB.Find(&book, id)
	c.JSON(http.StatusOK, book)
	return
}

func (p *BooksController) Update(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *BooksController) DeleteOneByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	database.DB.Delete(&book, id)
	c.Status(http.StatusOK)
	return
}

func (p *BooksController) Save(c *gin.Context) {
	var book models.Book
	c.Bind(&book)
	result := database.DB.Create(&book)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
