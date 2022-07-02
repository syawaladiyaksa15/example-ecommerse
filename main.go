package main

import (
	"e-Commerse/config"
	"e-Commerse/factory"
	"e-Commerse/routes"
)

func main() {
	// connection db
	dbConn := config.InitDB()

	// factory
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Logger.Fatal(e.Start(":8000"))
}
