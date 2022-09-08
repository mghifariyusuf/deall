package helper

import (
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

// Validate data form body
func Validate(data interface{}) (err error) {
	var validate *validator.Validate
	validate = validator.New()

	if err := validate.Struct(data); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
