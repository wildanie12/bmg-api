package validations

import (
	_web "gin-bmg-restful/entities/web"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var referralErrorMessages = map[string]string {
	"referral|required": "Referral is required",
}

// ValidateReferralRequest returns validation status
func ValidateReferralRequest(validate *validator.Validate, request _web.ReferralRequest) error {
	errors := []_web.ValidationErrorItem{}
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(request).FieldByName(err.Field())
			errors = append(errors, _web.ValidationErrorItem{
				Field: field.Tag.Get("form"),
				Error: referralErrorMessages[err.Field() + "|" + err.ActualTag()],
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