package handler

import (
	"coursework/web/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
	"net/http"
)

func (s *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := s.service.AuthService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *Handler) signIn(c *gin.Context) {

	var input AuthInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := s.service.AuthService.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Header("Access-Control-Allow-Origin", "http://192.168.11.40:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	fmt.Println(token)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
