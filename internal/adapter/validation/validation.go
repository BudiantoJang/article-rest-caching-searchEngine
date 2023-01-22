package validation

import (
	"context"
	"fmt"
	"jang-article/internal/model"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type RequestValidator struct {
	validate *validator.Validate
}

func New() *RequestValidator {
	r := &RequestValidator{
		validate: validator.New(),
	}

	r.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return r
}

func (r *RequestValidator) ValidateRequest(ctx context.Context, req interface{}) error {
	if err := r.validate.Struct(req); err != nil {
		var validationResponse model.ErrorDetail

		if _, ok := err.(*validator.InvalidValidationError); ok {
			validationResponse = model.ErrorDetail{
				Code:        model.ResponseCode99,
				Description: err,
			}
			fmt.Println("1")
			return validationResponse
		}
		validationResponse.Description = err

		return validationResponse
	}
	return nil
}
