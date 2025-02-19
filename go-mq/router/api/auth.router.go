package api

import (
	"github.com/bantawa04/go-mq/app/controller"
	"github.com/bantawa04/go-mq/app/middleware"
	"github.com/bantawa04/go-mq/app/repository"
	"github.com/bantawa04/go-mq/app/service"
	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	app            *fiber.App
	authController controller.AuthController
}

func NewAuthRouter(app *fiber.App) *AuthRouter {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	authController := controller.NewAuthController(userService)

	return &AuthRouter{
		app:            app,
		authController: authController,
	}
}

func (r *AuthRouter) Setup(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/register", middleware.DBTransactionHandler(), r.authController.RegisterUser)
}
