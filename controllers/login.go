package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct{}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginController) Run(c *gin.Context) {

	var loginCredentials = LoginRequestBody{}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		return
	}

	var customer models.Customer
	database.DB.Where("email = ?", loginCredentials.Email).Find(&customer)

	if customer.CustomerID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(loginCredentials.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Passwords doesn't match
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Wrong password"})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	jwt, err := generateToken(loginCredentials.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"JWT": jwt})
	return
}

func generateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["email"] = email
	jwtSecret := []byte(os.Getenv("JWTSECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}
