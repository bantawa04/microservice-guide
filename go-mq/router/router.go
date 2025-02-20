package router

import (
	"github.com/bantawa04/go-mq/router/api"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app           *fiber.App
	healthRouter  *api.HealthRouter
	userRouter    *api.UserRouter
	todoRouter    *api.TodoRouter
	authRouter    *api.AuthRouter
	productRouter *api.ProductRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:           app,
		healthRouter:  api.NewHealthRouter(app),
		userRouter:    api.NewUserRouter(app),
		todoRouter:    api.NewTodoRouter(app),
		authRouter:    api.NewAuthRouter(app),
		productRouter: api.NewProductRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)
	app.Stack()

	// Setup API routes with rate limiter
	apiRoute := app.Group("/api")

	// Setup individual route groups
	router.healthRouter.Setup(apiRoute)
	router.userRouter.Setup(apiRoute)
	router.todoRouter.Setup(apiRoute)
	router.authRouter.Setup(apiRoute)
	router.productRouter.Setup(apiRoute)
}
