package models

type Book struct {
	ISBN              string `gorm:"primary_key; size:14" binding:"required"`
	Title             string `json:"title" gorm:"size:255" binding:"required"`
	AuthorID          uint   `json:"authorID"`
	Author            Author `gorm:"references:AuthorID;constraint:OnDelete:CASCADE;"`
	YearOfPublication string `json:"yearOfPublication" gorm:"size:4"`
	AvailableCopies   uint   `json:"availableCopies" binding:"required"`
}
