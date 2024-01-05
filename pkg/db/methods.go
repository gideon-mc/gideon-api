package db

import (
	"fmt"
	"log"
	"strings"

	utils "github.com/gideon-mc/gideon-api/pkg"
	_ "github.com/go-sql-driver/mysql"
)

func (db *DB) SafeExec(cmd string) {
	_, err := db.Exec(cmd)
	if err != nil {
		log.Panicf("(Database) %s", err)
	}
}

func (db *DB) HasRows(query string) bool {
	rows, err := db.Query(query)
	if err != nil {
		log.Panicf("(Database) %s", err)
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}

	return false
}

func (db *DB) ClaimTable(name string, fields []string) {
	if db.HasRows(fmt.Sprintf("SHOW TABLES LIKE '%s';", name)) {
		return
	}

	db.SafeExec(
		fmt.Sprintf(
			"CREATE TABLE %s (%s)",
			name,
			strings.Join(fields, ","),
		),
	)
	log.Printf("(Database) Created table %s", name)
}

func (db *DB) ClaimRow(name string, predicate string, fields map[string]string) bool {
	if db.HasRows(fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		strings.Split(predicate, "=")[0],
		name,
		predicate,
	)) {
		return false
	}

	keys, values := utils.MapToKeysAndValues(fields)
	db.SafeExec(fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		name,
		strings.Join(keys, ","),
		strings.Join(values, ","),
	))

	log.Printf("(Database) Created row %q in %s", predicate, name)
	return true
}
