package api

import (
	"github.com/bantawa04/go-mq/app/controller"
	"github.com/bantawa04/go-mq/app/middleware"
	"github.com/bantawa04/go-mq/app/repository"
	"github.com/bantawa04/go-mq/app/service"
	"github.com/gofiber/fiber/v2"
)

type BookRouter struct {
	app            *fiber.App
	bookController controller.BookController
}

func NewBookRouter(app *fiber.App) *BookRouter {
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	return &BookRouter{
		app:            app,
		bookController: bookController,
	}
}

func (r *BookRouter) Setup(api fiber.Router) {
	books := api.Group("/books")
	books.Post("", middleware.DBTransactionHandler(), r.bookController.CreateBook)
}
