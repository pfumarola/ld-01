package main

import (
	"flag"

	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/database/migrate"
	"github.com/pfumarola/ld-01/database/seed"
	"github.com/pfumarola/ld-01/server"
)

func main() {

	db := database.Init()

	//CLI command
	shouldMigrate := flag.Bool("m", false, "Migrate the db")
	shouldSeed := flag.Bool("s", false, "Seed the db")
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
	}

	server.Init()
}
