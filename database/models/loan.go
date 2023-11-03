package models

import "time"

type Loan struct {
	LoanID       uint      `json:"loanID" gorm:"primary_key"`
	UserID       uint      `json:"userID" binding:"required"`
	User         User      `gorm:"references:UserID;constraint:OnDelete:CASCADE;"`
	BookID       string    `json:"bookID" binding:"required"`
	Book         Book      `gorm:"references:ISBN"`
	LoanDate     time.Time `json:"loanDate"`
	DueDate      time.Time `json:"dueDate"`
	ActualReturn time.Time `json:"actualReturn"`
}
