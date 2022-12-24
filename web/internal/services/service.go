package services

import (
	"coursework/web/internal/broker"
	"coursework/web/internal/models"
	"coursework/web/internal/repositories"
)

type AuthService interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type GithubRepositoryService interface {
	GetAll() ([]models.GithubRepository, error)
	GetRepoDataFromLocalBase(id int) (models.GithubRepository, error)
	SetRepoData(repo models.GithubRepository) (int, error)
	AddRepoData(url string)
}

type FavoritesService interface {
	GetAll(userId int) ([]models.GithubRepository, error)
	PutToFavorites(user int, id int) error
	RemoveFromFavorites(user int, id int) error
}

type Service struct {
	AuthService
	GithubRepositoryService
	FavoritesService
}

func NewService(repository *repositories.Repository, rabbit *broker.Broker) *Service {
	return &Service{
		AuthService:             NewAuthService(repository.AuthRepo),
		GithubRepositoryService: NewRepositoriesService(repository.GithubRepositoryRepo, rabbit),
		FavoritesService:        NewFavoriteService(repository.FavoritesRepo, rabbit),
	}
}
