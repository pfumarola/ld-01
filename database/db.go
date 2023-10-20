package database

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbuser := os.Getenv("DBUSER")
	bpassword := os.Getenv("DBPASSWORD")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", dbuser, bpassword, dbhost, dbport)
	sqlDB, _ := sql.Open("pgx", connStr)

	DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return DB
}
