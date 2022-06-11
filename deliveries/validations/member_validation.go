package validations

import (
	_web "gin-bmg-restful/entities/web"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var userErrorMessages = map[string]string {
	"Username|required": "Username is required",
	"Password|required": "Password is required",
	"Name|required": "Name is required",
	"Email|required": "Email is required",
}

// ValidateUserRequest returns validation status
func ValidateUserRequest(validate *validator.Validate, request _web.UserRequest) error {
	errors := []_web.ValidationErrorItem{}
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(request).FieldByName(err.Field())
			errors = append(errors, _web.ValidationErrorItem{
				Field: field.Tag.Get("form"),
				Error: userErrorMessages[err.Field() + "|" + err.ActualTag()],
			})
		}
	}
	if len(errors) > 0 {
		return _web.ValidationError{
			Message: "Validation error",
			Code: 422,
			Errors: errors,
		}
	}
	return nil
}