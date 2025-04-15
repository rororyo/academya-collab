package route

import (
	"collab-be/internal/delivery/http"
	"collab-be/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	// users
	c.App.Post("/api/users/register", c.UserController.Register)
	c.App.Post("/api/users/login", c.UserController.Login)
	c.App.Get("/api/users/user/:id", c.UserController.Get)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	//authenticated users
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Post("/api/users/logout", c.UserController.Logout)

	// Admin-only
	adminOnly := c.App.Group("/api/admin", middleware.RequireRole("admin"))
	// users
	adminOnly.Get("/users", c.UserController.List)
	adminOnly.Delete("/users/:id", c.UserController.Delete)
}
