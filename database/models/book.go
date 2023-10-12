package models

type Book struct {
	ISBN              string `gorm:"primary_key"`
	Title             string `json:"title"`
	AuthorID          uint   `json:"authorID"`
	YearOfPublication string `json:"yearOfPublication"`
	AvailableCopies   uint   `json:"availableCopies"`
	TotalCopies       uint   `json:"totalCopies"`
}
