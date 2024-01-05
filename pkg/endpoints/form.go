package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

type Form interface {
	Validate() error
}

func GetValidForm[T Form](ctx *fiber.Ctx, form T) (T, error) {
	if err := ctx.BodyParser(form); err != nil {
		return form, err
	}
	if err := form.Validate(); err != nil {
		ctx.Status(422).SendString(err.Error())
		return form, err
	}

	return form, nil
}
