package models

type User struct {
	UserID       uint   `gorm:"primary_key"`
	FirstName    string `json:"firstName" gorm:"size:255"`
	LastName     string `json:"lastName" gorm:"size:255"`
	Address      string `json:"address" gorm:"size:255"`
	Phone        string `json:"phone" gorm:"size:15"`
	Email        string `json:"email" gorm:"size:255"`
	PasswordHash string `json:"-" gorm:"size:255"`
	IsAdmin      bool   `json:"isAdmin"`
}
