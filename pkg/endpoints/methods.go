package endpoints

import (
	"fmt"

	utils "github.com/gideon-mc/gideon-api/pkg"
	"github.com/gideon-mc/gideon-api/pkg/auth"
	"github.com/gideon-mc/gideon-api/pkg/endpoints/forms"
	"github.com/gofiber/fiber/v2"
)

func (e *Endpoint) Index(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	ctx.JSON(fiber.Map{
		"version": "1.0.0",
	})

	return nil
}

func (e *Endpoint) AuthRegister(ctx *fiber.Ctx) error {
	if !auth.ValidateToken(e.Database, ctx, "is_admin=true") {
		return nil
	}

	form, err := GetValidForm(ctx, new(forms.RegisterForm))
	if err != nil {
		return err
	}

	e.Database.ClaimRow("discord", fmt.Sprintf("discord_id=%q", form.DiscordID), map[string]string{
		"discord_id": utils.SurroundWithQuotes(form.DiscordID),
		"slug":       utils.SurroundWithQuotes(form.DiscordSlug),
	})

	password_hash := auth.EncryptPassword(form.Password)
	e.Database.ClaimRow("user", fmt.Sprintf("user_id=%q", form.UserID), map[string]string{
		"user_id": utils.SurroundWithQuotes(form.UserID),
		"password_hash": utils.SurroundWithQuotes(
			password_hash,
		),
		"fk_discord_id": utils.SurroundWithQuotes(form.DiscordID),
	})

	ctx.SendString(password_hash)
	return nil
}

func (e *Endpoint) AuthLogin(ctx *fiber.Ctx) error {
	form, err := GetValidForm(ctx, new(forms.LoginForm))
	if err != nil {
		return err
	}

	rows, err := e.Database.Query(fmt.Sprintf("SELECT password_hash FROM user WHERE user_id=%q;", form.UserID))
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		ctx.Status(404).SendString("User is not found")
		return nil
	}
	var password_hash string
	if err := rows.Scan(&password_hash); err != nil {
		return err
	}

	if !auth.CompareHashAndPassword(password_hash, form.Password) {
		ctx.Status(401).SendString("Wrong password.")
		return nil
	}

	ctx.SendString(password_hash)
	return nil
}
