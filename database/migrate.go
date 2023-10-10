package database

import (
	"log"

	"github.com/pfumarola/ld-01/database/models"
)

func Migrate() {
	db := Init()
	log.Println("Migrating stuff...")
	db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Customer{}, &models.Transaction{})
	log.Println("Done.")
}
