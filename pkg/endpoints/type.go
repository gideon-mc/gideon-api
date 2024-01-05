package endpoints

import (
	"github.com/gideon-mc/gideon-api/pkg/db"
	"github.com/gofiber/fiber/v2"
)

type Endpoint struct {
	Group fiber.Router
    Database *db.DB
}
