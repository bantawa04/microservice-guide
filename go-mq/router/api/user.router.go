package api

import (
	"github.com/bantawa04/go-mq/app/controller"
	"github.com/bantawa04/go-mq/app/middleware"
	"github.com/bantawa04/go-mq/app/repository"
	"github.com/bantawa04/go-mq/app/service"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	app            *fiber.App
	userController controller.UserController
}

func NewUserRouter(app *fiber.App) *UserRouter {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &UserRouter{
		app:            app,
		userController: userController,
	}
}

func (r *UserRouter) Setup(api fiber.Router) {
	users := api.Group("/users")
	users.Get("", r.userController.GetUsers)
	users.Get("/:id", r.userController.GetUserByID)
	users.Post("", middleware.DBTransactionHandler(), r.userController.CreateUser)
	users.Put("/:id", middleware.DBTransactionHandler(), r.userController.UpdateUser)
	users.Delete("/:id", middleware.DBTransactionHandler(), r.userController.DeleteUser)
}
