package transport

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()

	if err := validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		date := fl.Field().String()
		_, err := time.Parse("2006-01-02", date)
		return err == nil
	}); err != nil {
		return errors.New(formatValidationErrors(err))
	}

	err := validate.Struct(s)
	if err != nil {
		var errMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errMessages = append(errMessages, err.Error())
		}
		return errors.New(formatValidationErrors(err))
	}

	return nil
}

func formatValidationErrors(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var errMessages []string
		for _, e := range errs {
			errMessages = append(errMessages, fmt.Sprintf("Field '%s' %s", e.Field(), e.Tag()))
		}
		return strings.Join(errMessages, ", ")
	}
	return err.Error()
}
