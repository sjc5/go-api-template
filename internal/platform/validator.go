package platform

import (
	"github.com/go-playground/validator/v10"
	"github.com/sjc5/kit/pkg/validate"
)

var Validate validate.Validate

func init() {
	Validate.Instance = validator.New(validator.WithRequiredStructEnabled())
}
