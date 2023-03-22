package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const timestampFormat = "20060102T150405Z"

type StructValidatorInt interface {
	Validate(structForCheck interface{}) error
}

type StructValidator struct {
	goValidator *validator.Validate
}

func NewStructValidator(goValidator *validator.Validate) *StructValidator {
	_ = goValidator.RegisterValidation("validateTimestampFormat", validateTimestampFormat)
	_ = goValidator.RegisterValidation("validateTimezone", validateTimezone)

	return &StructValidator{
		goValidator: goValidator,
	}
}

func (sv *StructValidator) Validate(structForCheck interface{}) error {
	err := sv.goValidator.Struct(structForCheck)
	if err != nil {
		return &StructValidatorError{
			ErrorMessage: err.Error(),
		}
	}

	return nil
}

func validateTimestampFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "<invalid Value>" {
		return false
	}

	_, err := time.Parse(timestampFormat, fl.Field().String())
	if err != nil {
		return false
	}

	return true
}

func validateTimezone(fl validator.FieldLevel) bool {
	if fl.Field().String() == "<invalid Value>" {
		return false
	}

	_, err := time.LoadLocation(fl.Field().String())
	if err != nil {
		return false
	}

	return true
}

type StructValidatorError struct {
	ErrorMessage string
}

func (sve *StructValidatorError) Error() string {
	return sve.ErrorMessage
}
