package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	UserID    uint   `gorm:"primary_key"`
	FirstName string `json:"firstName" gorm:"size:255" binding:"required"`
	LastName  string `json:"lastName" gorm:"size:255" binding:"required"`
	Address   string `json:"address" gorm:"size:255"`
	Phone     string `json:"phone" gorm:"size:15"`
	Email     string `json:"email" gorm:"size:255, unique" binding:"required,email"`
	Password  string `json:"-" gorm:"size:255" binding:"required"`
	IsAdmin   bool   `json:"isAdmin" gorm:"default:false"`
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
