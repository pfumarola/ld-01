package models

import (
	"time"
)

type Author struct {
	AuthorID    uint      `gorm:"primary_key"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastNamw"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Nationality string    `json:"nationality"`
	Biography   string    `json:"biography"`
}
