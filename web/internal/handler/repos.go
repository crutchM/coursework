package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Input struct {
	Url string `json:"url"`
}

func (s *Handler) addNewRepo(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	s.service.GithubRepositoryService.AddRepoData(input.Url)
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	s.service.FavoritesService.PutToFavorites(userId, 0)
	sendJsonResponse(c, 200, "message", "ok")
}

func (s *Handler) deleteRepoFromFavs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	err = s.service.FavoritesService.RemoveFromFavorites(userId, id)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}
	sendJsonResponse(c, 200, "message", "ok")

}

func (s *Handler) getRepoData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}

	repo, err := s.service.GithubRepositoryService.GetRepoDataFromLocalBase(id)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	sendJsonResponse(c, 200, "repository", repo)
}
