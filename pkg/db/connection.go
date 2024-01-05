package db

import (
	"database/sql"
	"log"
	"os"
)

func ConnectToDB() *DB {
	instance, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("(Database) Failed to connect: %v", err)
	}

	if err := instance.Ping(); err != nil {
		instance.Close()
		log.Fatalf("(Database) Failed to ping: %v", err)
	}

	log.Println("(Database) Successfully connected to database.")
	return &DB{instance}
}
