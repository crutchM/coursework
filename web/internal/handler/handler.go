package handler

import (
	"coursework/web/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (s *Handler) InitRoues() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("sign-up")
	}
	//oauth := router.Group("/oauth")
	//{
	//
	//}

	api := router.Group("/api", s.userIdentity)
	{
		repos := api.Group("/repos", s.userIdentity)
		{
			repos.GET("/")
			repos.POST("/")
			repos.DELETE("/")
		}
	}
	return router
}
