package migrate

import (
	"log"

	"github.com/pfumarola/ld-01/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	log.Println("Migrating stuff...")
	db.AutoMigrate(
		&models.Author{},
		&models.Book{},
		&models.User{},
		&models.Loan{},
	)
	log.Println("Done.")
}
