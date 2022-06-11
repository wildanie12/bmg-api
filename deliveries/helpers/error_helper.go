package helpers

import (
	"gin-bmg-restful/entities/web"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// WebErrorResponse helps return correct response type
// Based on what kind of error that returned by services
func WebErrorResponse(c *gin.Context, err error) {
	if reflect.TypeOf(err).String() == "web.Error" {
		webErr := err.(web.Error)
		c.JSON(webErr.Code, web.ErrorResponse{
			Status: "ERROR",
			Code:   webErr.Code,
			Error:  webErr.Error(),
		})
		return
	} else if reflect.TypeOf(err).String() == "web.ValidationError" {
		valErr := err.(web.ValidationError)
		c.JSON(valErr.Code, web.ValidationErrorResponse{
			Status: "ERROR",
			Code:   valErr.Code,
			Error:  valErr.Error(),
			Errors: valErr.Errors,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, web.ErrorResponse{
		Status: "ERROR",
		Code: http.StatusInternalServerError,
		Error: "Server Error",
	})
	return
}