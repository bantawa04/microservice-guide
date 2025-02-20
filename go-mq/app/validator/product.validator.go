package validator

import (
	"fmt"
	"github.com/bantawa04/go-mq/app/response"
	"github.com/go-playground/validator/v10"
)

type ProductValidator struct {
	*validator.Validate
}

func NewProductValidator() ProductValidator {
	v := validator.New()
	return ProductValidator{
		Validate: v,
	}
}

func (cv ProductValidator) generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func (cv ProductValidator) GenerateValidationResponse(err error) []response.ValidationError {
	var validations []response.ValidationError
	for _, value := range err.(validator.ValidationErrors) {
		field, rule := value.Field(), value.Tag()
		validation := response.ValidationError{Field: field, Message: cv.generateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	return validations
}
