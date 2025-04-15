package helper

import (
	"collab-be/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
