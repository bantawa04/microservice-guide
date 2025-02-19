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
	Login(c *fiber.Ctx) error
}

type authController struct {
	userService    service.UserService
	authService    service.AuthService
	validator      validator.RegisterValidator
	loginValidator validator.LoginValidator
}

func NewAuthController(userService service.UserService, authService service.AuthService) AuthController {
	return &authController{
		userService:    userService,
		authService:    authService,
		validator:      validator.NewRegisterationValidator(),
		loginValidator: validator.NewLoginValidator(),
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

func (ctrl *authController) Login(c *fiber.Ctx) error {
	reqData := new(request.LoginRequestData)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.loginValidator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.loginValidator.GenerateValidationResponse(errors))
	}

	token, err := ctrl.authService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		token, "Login successful")
}
