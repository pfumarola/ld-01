package main

import (
	"flag"

	"github.com/pfumarola/ld-01/database"
	"github.com/pfumarola/ld-01/server"
)

func main() {

	//CLI command
	var shouldMigrate = flag.Bool("m", false, "Migrate the db")
	flag.Parse()

	if *shouldMigrate {
		database.Migrate()
		return
	}

	database.Init()
	server.Init()
}
