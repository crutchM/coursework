package services

import (
	"coursework/web/internal/broker"
	"coursework/web/internal/models"
	"coursework/web/internal/repositories"
)

type FavoriteService struct {
	broker *broker.Broker
	repo   repositories.FavoritesRepo
}

func NewFavoriteService(repo repositories.FavoritesRepo, rabbit *broker.Broker) *FavoriteService {
	return &FavoriteService{repo: repo, broker: rabbit}
}

func (f FavoriteService) GetAll(userId int) ([]models.GithubRepository, error) {
	return f.repo.GetAll(userId)
}

func (f FavoriteService) PutToFavorites(user int, id int) error {
	return f.repo.PutToFavorites(user, id)
}

func (f FavoriteService) RemoveFromFavorites(user int, id int) error {
	return f.repo.RemoveFromFavorites(user, id)
}
