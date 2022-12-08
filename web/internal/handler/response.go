package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func NewError(message string) *Error {
	return &Error{Message: message}
}
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, NewError(message))
}

func sendJsonResponse(c *gin.Context, code int, key string, value interface{}) {
	c.JSON(code, map[string]interface{}{key: value})
}
