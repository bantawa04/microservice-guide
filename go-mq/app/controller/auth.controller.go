package controller

import (
	"github.com/bantawa04/go-mq/app/middleware"
	"github.com/bantawa04/go-mq/app/request"
	"github.com/bantawa04/go-mq/app/response"
	"github.com/bantawa04/go-mq/app/service"
	"github.com/bantawa04/go-mq/app/validator"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	RegisterUser(c *fiber.Ctx) error
}

type authController struct {
	userService service.UserService
	validator   validator.AuthValidator
}

func NewAuthController(userService service.UserService) AuthController {
	return &authController{
		userService: userService,
		validator:   validator.NewAuthValidator(),
	}
}

func (ctrl *authController) RegisterUser(c *fiber.Ctx) error {
	// Get transaction from context
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)

	reqData := new(request.RegisterUserRequestData)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	userModel := reqData.ToModel()

	err := ctrl.userService.WithTrx(tx).CreateUser(userModel)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, fiber.StatusCreated, "User Created Successfully")
}
