package models

import (
	"time"
)

type Author struct {
	AuthorID    uint      `json:"authorID" gorm:"primary_key"`
	FirstName   string    `json:"firstName" gorm:"size:255" binding:"required"`
	LastName    string    `json:"lastName" gorm:"size:255" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" gorm:"type:date"`
	Nationality string    `json:"nationality" gorm:"size:255"`
	Biography   string    `json:"biography"`
}
