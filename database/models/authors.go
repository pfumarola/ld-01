package models

import (
	"time"

	"gorm.io/gorm"
)

type author struct {
	gorm.Model
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastNamw"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Nationality string    `json:"nationality"`
	Biography   string    `json:"biography"`
}
