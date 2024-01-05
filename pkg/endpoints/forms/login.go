package forms

import "errors"

type LoginForm struct {
	UserID   string `json:"user_id" form:"user_id"`
	Password string `json:"password" form:"password"`
}

func (form *LoginForm) Validate() error {
	if len(form.UserID) == 0 {
		return errors.New("Username is too short.")
	}

	if len(form.Password) == 0 {
		return errors.New("Password is too short.")
	}

	return nil
}
