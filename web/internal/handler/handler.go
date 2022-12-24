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
	router.POST("/data", s.putRepoData)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", s.signIn)
		auth.POST("sign-up", s.signUp)
	}
	//oauth := router.Group("/oauth")
	//{
	//
	//}

	api := router.Group("/api", s.userIdentity)
	{
		repos := api.Group("/repos")
		{
			repos.GET("/:id", s.getRepoData)
			repos.POST("/", s.addNewRepo)
			repos.GET("/favs", s.getFavs)
			repos.GET("/", s.getAll)
			repos.DELETE("/:id", s.deleteRepoFromFavs)
		}
	}
	return router
}
