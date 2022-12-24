package handler

import (
	"coursework/web/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

	time.Sleep(3 * time.Second)
	id, err := s.service.GithubRepositoryService.GetByUrl(input.Url)
	if err != nil {
		return
	}
	s.service.FavoritesService.PutToFavorites(userId, id)
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

func (s *Handler) putRepoData(c *gin.Context) {
	var repo models.GithubRepository

	if err := c.BindJSON(&repo); err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}

	id, err := s.service.SetRepoData(repo)
	if id == 99999 {
		return
	}
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	sendJsonResponse(c, 200, "id", id)
}

func (s *Handler) getFavs(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, 400, err.Error())
	}

	res, err := s.service.FavoritesService.GetAll(id)

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	sendJsonResponse(c, 200, "repos", res)
}

func (s *Handler) getAll(c *gin.Context) {
	res, err := s.service.GithubRepositoryService.GetAll()
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	sendJsonResponse(c, 200, "repos", res)
}
