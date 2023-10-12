package wipe

import (
	"log"

	"github.com/pfumarola/ld-01/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	log.Println("Wiping db...")
	db.Migrator().DropTable(
		&models.Author{},
		&models.Book{},
		&models.Customer{},
		&models.Transaction{},
	)
	log.Println("Done.")
}
