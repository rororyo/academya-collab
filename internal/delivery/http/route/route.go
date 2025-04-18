package route

import (
	"collab-be/internal/delivery/http"
	"collab-be/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	JobController  *http.JobController
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
	c.App.Get("/api/jobs", c.JobController.List)
	c.App.Get("/api/jobs/:id", c.JobController.Get)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	//authenticated users
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Post("/api/users/logout", c.UserController.Logout)

	// Company-only
	companyOnly := c.App.Group("/api/company", middleware.RequireRole("company"))
	// jobs
	companyOnly.Post("/jobs", c.JobController.Create)
	companyOnly.Put("/jobs/:id", c.JobController.Update)
	companyOnly.Delete("/jobs/:id", c.JobController.Delete)

	// Admin-only
	adminOnly := c.App.Group("/api/admin", middleware.RequireRole("admin"))
	// users
	adminOnly.Get("/users", c.UserController.List)
	adminOnly.Delete("/users/:id", c.UserController.Delete)
}
