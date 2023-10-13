package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
)

type CustomersController struct{}

func (p *CustomersController) FindAll(c *gin.Context) {
	var customers []models.Customer
	database.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
	return
}

func (p *CustomersController) FindOneByID(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	database.DB.Find(&customer, id)
	c.JSON(http.StatusOK, customer)
	return
}

func (p *CustomersController) Update(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *CustomersController) DeleteOneByID(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	database.DB.Delete(&customer, id)
	c.Status(http.StatusOK)
	return
}

func (p *CustomersController) Save(c *gin.Context) {
	var customer models.Customer
	c.Bind(&customer)
	result := database.DB.Create(&customer)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
