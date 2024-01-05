package forms

import (
	"errors"

	"github.com/gideon-mc/gideon-api/pkg/auth"
)

type RegisterForm struct {
	UserID      string `json:"user_id" form:"user_id"`
	Password    string `json:"password" form:"password"`
	DiscordID   string `json:"discord_id" form:"discord_id"`
	DiscordSlug string `json:"discord_slug" form:"discord_slug"`
}

func (form *RegisterForm) Validate() error {
	if !auth.IsPasswordSecure(form.Password) {
		return errors.New("Password is too weak.")
	}

	if len(form.UserID) == 0 {
		return errors.New("Username is too short.")
	}

	if len(form.DiscordSlug) == 0 {
		return errors.New("Discord slug is too short.")
	}

	if len(form.DiscordID) != 18 {
		return errors.New("Invalid discord id. It must be 18 characters long.")
	}

	return nil
}
