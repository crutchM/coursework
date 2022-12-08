package repositories

import (
	"coursework/web/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	CreateUser(user models.User) (string, error)
	GetUser(login, password string) (models.User, error)
}

type GithubRepositoryRepo interface {
	GetAll() ([]models.GithubRepository, error)
	GetRepoDataFromLocalBase(id string) (models.GithubRepository, error)
	SetRepoData(repo models.GithubRepository) (int, error)
}

type FavoritesRepo interface {
	GetAll(userId int) ([]models.GithubRepository, error)
	GetOne(id int) (models.GithubRepository, error)
	PutToFavorites(user int, id int) error
	RemoveFromFavorites(user int, id int) error
}

type Repository struct {
	AuthRepo
	GithubRepositoryRepo
	FavoritesRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepo:             NewAuthRepository(db),
		GithubRepositoryRepo: NewGithubRepository(db),
		FavoritesRepo:        NewFavoritesRepository(db),
	}
}
