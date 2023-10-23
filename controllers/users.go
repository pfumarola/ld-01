package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
)

type UsersController struct{}

func (p *UsersController) FindAll(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
	return
}

func (p *UsersController) FindOneByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	database.DB.Find(&user, id)
	c.JSON(http.StatusOK, user)
	return
}

func (p *UsersController) Update(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *UsersController) DeleteOneByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	database.DB.Delete(&user, id)
	c.Status(http.StatusOK)
	return
}

func (p *UsersController) Save(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
}
