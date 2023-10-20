package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/migrate"
	"github.com/pfumarola/ld-01/database/seed"
	"github.com/pfumarola/ld-01/database/wipe"
	"github.com/pfumarola/ld-01/server"
)

func main() {

	godotenv.Load()

	db := database.Init()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	//CLI command
	shouldMigrate := flag.Bool("m", false, "Migrate the db")
	shouldSeed := flag.Bool("s", false, "Seed the db")
	shouldWipe := flag.Bool("w", false, "Wipe data in the db")
	flag.Parse()

	switch {
	case *shouldMigrate:
		{
			migrate.Run(db)
			return
		}
	case *shouldSeed:
		{
			seed.Run(db)
			return
		}
	case *shouldWipe:
		{
			wipe.Run(db)
			return
		}
	}

	server.Init()
}
