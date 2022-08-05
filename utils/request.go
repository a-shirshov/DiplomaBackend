package utils

import (
	"Diploma/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

func GetAUFromContext(c *gin.Context) (*AccessDetails, error) {
	ctxau, ok := c.Get("access_details")
	if !ok {
		return nil, errors.ErrNoTokenInContext
	}

	au, ok := ctxau.(AccessDetails)
	if !ok {
		return nil, errors.ErrNoTokenInContext
	}
	
	return &au, nil
}

func Validate(object interface{}) error {
	validator := validator.New()
	return validator.Struct(object)
}