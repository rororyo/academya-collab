package config

import (
	"collab-be/internal/delivery/http"
	"collab-be/internal/delivery/http/middleware"
	"collab-be/internal/delivery/http/route"
	"collab-be/internal/repository"
	"collab-be/internal/usecase"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	//setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	//setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	//setup controllers
	userController := http.NewUserController(userUseCase, config.Log)

	//setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
