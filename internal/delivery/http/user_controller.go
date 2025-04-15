package http

import (
	"collab-be/internal/delivery/http/middleware"
	"collab-be/internal/model"
	"collab-be/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	Usecase *usecase.UserUseCase
}

func NewUserController(usecase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		Usecase: usecase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}
	c.Log.Debugf("Parsed request: %s", request)
	response, err := c.Usecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create user: %v", err)
		return err
	}
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.Usecase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user: %v", err)
		return err
	}
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.Usecase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Get(ctx *fiber.Ctx) error {
	request := &model.GetUserRequest{
		ID: ctx.Params("id"),
	}
	response, err := c.Usecase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.Usecase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}

func (c *UserController) List(ctx *fiber.Ctx) error {

	request := &model.SearchUserRequest{
		Name:  ctx.Query("name"),
		Email: ctx.Query("email"),
		Page:  ctx.QueryInt("page"),
		Size:  ctx.QueryInt("size"),
	}

	responses, total, err := c.Usecase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to search user")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.UserResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteUserRequest{
		ID: ctx.Params("id"),
	}
	response, err := c.Usecase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to delete user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
