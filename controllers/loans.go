package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
)

type LoansController struct{}

func (p *LoansController) FindAll(c *gin.Context) {
	var loans []models.Loan
	database.DB.Find(&loans)
	c.JSON(http.StatusOK, loans)
	return
}

func (p *LoansController) FindOneByID(c *gin.Context) {
	var loan models.Loan
	id := c.Param("id")
	database.DB.Find(&loan, id)
	if loan.LoanID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Id not found"})
		return
	}
	c.JSON(http.StatusOK, loan)
	return
}

func (p *LoansController) Update(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
}

func (p *LoansController) DeleteOneByID(c *gin.Context) {
	var loan models.Loan
	id := c.Param("id")
	database.DB.Delete(&loan, id)
	c.Status(http.StatusOK)
	return
}

func (p *LoansController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
	return
	/* TO DO:
	**	- check if the book is available
	**	-	decrese the book available quantity
	var loan models.Loan
	book := models.Book{}
	c.Bind(&loan)
	book.ISBN = loan.BookID
	database.DB.Find(&book)
	if book.ISBN == "" {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Book not found"})
		return
	}
	log.Println("Book", book, book.ISBN)
	loan.LoanDate = time.Now()
	loan.DueDate = time.Now().AddDate(0, 1, 0)
	loan.ActualReturn = time.Time{}

	result := database.DB.Create(&loan)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	return
	*/
}
