package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
)

type TransactionsController struct{}

func (p *TransactionsController) FindAll(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(http.StatusOK, transactions)
	return
}

func (p *TransactionsController) FindOneByID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")
	database.DB.Find(&transaction, id)
	c.JSON(http.StatusOK, transaction)
	return
}

func (p *TransactionsController) Update(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *TransactionsController) DeleteOneByID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")
	database.DB.Delete(&transaction, id)
	c.Status(http.StatusOK)
	return
}

func (p *TransactionsController) Save(c *gin.Context) {
	var transaction models.Transaction
	c.Bind(&transaction)
	result := database.DB.Create(&transaction)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
