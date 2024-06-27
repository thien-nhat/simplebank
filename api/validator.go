package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/thien-nhat/simplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	// currency := fl.Field().String()
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
	
}
