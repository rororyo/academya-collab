package middleware

import (
	"collab-be/internal/helper"
	"collab-be/internal/model"
	"collab-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUsecase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userUsecase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUsecase.Verify(ctx.UserContext(), request)

		if err != nil {
			userUsecase.Log.Warnf("Failed to verify user: %v", err)
			return fiber.ErrUnauthorized
		}
		userUsecase.Log.Debugf("User: %v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return helper.GetUser(ctx)
}

func RequireRole(role string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := helper.GetUser(ctx) // Extract the user info

		if auth == nil || auth.Role != role {
			return fiber.NewError(fiber.StatusForbidden, "You do not have access to this resource")
		}

		return ctx.Next()
	}
}
