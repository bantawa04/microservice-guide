package controller

import (
	"github.com/bantawa04/go-mq/app/middleware"
	"github.com/bantawa04/go-mq/app/request"
	"github.com/bantawa04/go-mq/app/response"
	"github.com/bantawa04/go-mq/app/service"
	"github.com/bantawa04/go-mq/app/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookController interface {
	CreateBook(c *fiber.Ctx) error
}

type bookController struct {
	bookService service.BookService
	validator   validator.ProductValidator
}

func NewBookController(bookService service.BookService) BookController {
	return &bookController{
		bookService: bookService,
		validator:   validator.NewProductValidator(),
	}
}

func (ctrl *bookController) CreateBook(c *fiber.Ctx) error {
    tx := c.Locals(middleware.DBTransaction).(*gorm.DB)

    // Extract user ID from context
    userContext := c.Locals(middleware.UserContextKey).(map[string]interface{})
    userID := userContext["user_id"].(string)

    reqData := new(request.CreateBookRequestData)
    if err := c.BodyParser(reqData); err != nil {
        return err
    }

    if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
        return response.ValidationErrorResponse(c,
            ctrl.validator.GenerateValidationResponse(errors))
    }

    bookModel := reqData.ToModel(userID)

    err := ctrl.bookService.WithTrx(tx).CreateBook(bookModel)
    if err != nil {
        return err
    }

    return response.SuccessResponse(c, fiber.StatusCreated, "Book Created Successfully")
}
