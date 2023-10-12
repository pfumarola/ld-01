package database

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	connStr := "postgres://postgres:lastingdynamics@localhost:5432/postgres?sslmode=disable"
	sqlDB, _ := sql.Open("pgx", connStr)

	DB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return DB
}
