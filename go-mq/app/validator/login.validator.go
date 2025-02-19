package validator

import (
	"fmt"
	"regexp"

	"github.com/bantawa04/go-mq/app/response"
	"github.com/go-playground/validator/v10"
)

type LoginValidator struct {
	*validator.Validate
}

func NewLoginValidator() LoginValidator {
	v := validator.New()

	_ = v.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		if fl.Field().String() != "" {
			match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, fl.Field().String())
			return match
		}
		return true
	})
	return LoginValidator{
		Validate: v,
	}
}

func (cv LoginValidator) generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	case "email":
		return fmt.Sprintf("Field '%s' is not valid.", field)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func (cv LoginValidator) GenerateValidationResponse(err error) []response.ValidationError {
	var validations []response.ValidationError
	for _, value := range err.(validator.ValidationErrors) {
		field, rule := value.Field(), value.Tag()
		validation := response.ValidationError{Field: field, Message: cv.generateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	return validations
}
