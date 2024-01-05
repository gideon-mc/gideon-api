package gideon

import (
	"os"

	utils "github.com/gideon-mc/gideon-api/pkg"
	"github.com/gideon-mc/gideon-api/pkg/auth"
	"github.com/gideon-mc/gideon-api/pkg/db"
)

func ClaimDefaultRows(db *db.DB) {
	db.ClaimRow("discord", "discord_id='465886354941673473'", map[string]string{
		"discord_id": utils.SurroundWithQuotes("465886354941673473"),
		"slug":       utils.SurroundWithQuotes("bbfh"),
	})

	db.ClaimRow("user", "user_id='root'", map[string]string{
		"user_id": utils.SurroundWithQuotes("root"),
		"password_hash": utils.SurroundWithQuotes(
			auth.EncryptPassword(os.Getenv("ROOT_PASSWORD")),
		),
		"is_admin":      "true",
		"fk_discord_id": utils.SurroundWithQuotes("465886354941673473"),
	})
}
