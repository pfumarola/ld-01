package seed

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {

	log.Println("Seeding...")

	seedAuthors(db)

	seedBooks(db)

	seedCustomers(db)

	log.Println("Done.")

}
