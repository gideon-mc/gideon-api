package main

import (
	"log"

	"github.com/gideon-mc/gideon-api/internal/gideon"
	"github.com/gideon-mc/gideon-api/pkg/db"
	"github.com/gideon-mc/gideon-api/pkg/endpoints"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Could not load '.env' file.")
	}

	db := db.ConnectToDB()
	defer db.Close()

	gideon.ClaimTables(db)
	gideon.ClaimDefaultRows(db)

	app := gideon.SetupApp()
	endpoints.AssignRoutesTo(&endpoints.Endpoint{Group: app.Group("/api"), Database: db})

	log.Fatal(app.Listen(":5000"))
}
