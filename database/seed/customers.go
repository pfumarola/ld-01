package seed

import (
	"github.com/pfumarola/ld-01/database/models"
	"gorm.io/gorm"
)

func seedCustomers(db *gorm.DB) {
	db.Create(customers)
}

var customers = []models.Customer{
	{
		FirstName: "John",
		LastName:  "Doe",
		Address:   "123 Main St",
		Phone:     "555-123-4567",
		Email:     "john.doe@email.com",
		Score:     10,
	},
	{
		FirstName: "Jane",
		LastName:  "Smith",
		Address:   "456 Elm St",
		Phone:     "555-987-6543",
		Email:     "jane.smith@email.com",
		Score:     10,
	},
	{
		FirstName: "Alice",
		LastName:  "Johnson",
		Address:   "789 Oak St",
		Phone:     "555-555-5555",
		Email:     "alice.johnson@email.com",
		Score:     10,
	},
	{
		FirstName: "Bob",
		LastName:  "Brown",
		Address:   "101 Pine St",
		Phone:     "555-222-3333",
		Email:     "bob.brown@email.com",
		Score:     10,
	},
	{
		FirstName: "Eve",
		LastName:  "Wilson",
		Address:   "222 Cedar St",
		Phone:     "555-888-9999",
		Email:     "eve.wilson@email.com",
		Score:     10,
	},
}
