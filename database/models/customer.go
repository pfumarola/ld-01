package models

type Customer struct {
	CustomerID uint   `gorm:"primary_key"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Score      uint   `json:"score"`
}
