package auth

import (
	"fmt"
	"log"
	"strings"

	"github.com/gideon-mc/gideon-api/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func ValidateToken(db *db.DB, ctx *fiber.Ctx, predicate string) bool {
	token, ok := ctx.GetReqHeaders()["Auth-Token"]
	if !ok {
		ctx.Status(401).SendString("Auth-Token header is not present in the request header.")
		return false
	}

	partitions := strings.Split(token, "@")
	if len(partitions) != 2 {
		ctx.Status(401).SendString("Auth-Token has invalid format. Required is: 'NAME@HASH'")
		return false
	}

	var condition string
	if len(predicate) != 0 {
		condition = fmt.Sprintf("AND %s", predicate)
	}
	rows, err := db.Query(fmt.Sprintf("SELECT password_hash FROM user WHERE user_id=%q %s;", partitions[0], condition))
	if err != nil {
		log.Panicf("(Database) %s", err)
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		ctx.Status(401).SendString("User is not found")
		return false
	}
	var password_hash string
	if err := rows.Scan(&password_hash); err != nil {
		log.Panicln(err)
	}

	return true
}
