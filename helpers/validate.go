package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidatorHelpers struct {
	Validate *validator.Validate
}

func NewValidatorService() *ValidatorHelpers {
	return &ValidatorHelpers{
		Validate: validator.New(),
	}
}

func (validatorService *ValidatorHelpers) BuildAndGetValidationMessage(err error) string {
	var (
		validationMessage string
		validationError   = err.(validator.ValidationErrors)[0]
	)

	validationMessage = fmt.Sprintf("%s %s", validationError.Field(), validationError.Tag())
	return validationMessage
}

func (validatorService *ValidatorHelpers) ValidateIncomingRequest(request interface{}) error {
	return validatorService.Validate.Struct(request)
}
