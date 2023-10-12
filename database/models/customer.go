package models

type Customer struct {
	CustomerID uint   `gorm:"primary_key"`
	FirstName  string `json:"firstName" gorm:"size:255"`
	LastName   string `json:"lastName" gorm:"size:255"`
	Address    string `json:"address" gorm:"size:255"`
	Phone      string `json:"phone" gorm:"size:15"`
	Email      string `json:"email" gorm:"size:255"`
	Score      uint   `json:"score"`
}
