package models

import "time"

type Transaction struct {
	TransactionID uint      `json:"authorID" gorm:"primary_key"`
	CustomerID    uint      `json:"customerID"`
	BookID        uint      `json:"bookID"`
	LoanDate      time.Time `json:"loanDate"`
	DueDate       time.Time `json:"dueDate"`
	ActualReturn  time.Time `json:"actualReturn"`
}
