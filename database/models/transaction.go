package models

import "time"

type Transaction struct {
	TransactionID uint      `json:"transactionID" gorm:"primary_key"`
	CustomerID    uint      `json:"customerID"`
	Customer      Customer  `gorm:"references:CustomerID;constraint:OnDelete:CASCADE;"`
	BookID        uint      `json:"bookID"`
	Book          Book      `gorm:"references:ISBN"`
	LoanDate      time.Time `json:"loanDate"`
	DueDate       time.Time `json:"dueDate"`
	ActualReturn  time.Time `json:"actualReturn"`
}
