package seed

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {

	log.Println("Seeding...")

	seedAuthors(db)

	seedBooks(db)

	seedUsers(db)

	log.Println("Done.")

}
