package models

type Book struct {
	ISBN              string `gorm:"primary_key; size:14"`
	Title             string `json:"title" gorm:"size:255"`
	AuthorID          uint   `json:"authorID"`
	YearOfPublication string `json:"yearOfPublication" gorm:"size:4"`
	AvailableCopies   uint   `json:"availableCopies"`
	TotalCopies       uint   `json:"totalCopies"`
}
