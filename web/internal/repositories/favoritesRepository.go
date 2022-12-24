package repositories

import (
	"coursework/web/internal/models"
	"github.com/jmoiron/sqlx"
)

type FavoritesRepository struct {
	db *sqlx.DB
}

func NewFavoritesRepository(db *sqlx.DB) *FavoritesRepository {
	return &FavoritesRepository{db: db}
}

func (f FavoritesRepository) GetAll(userId int) ([]models.GithubRepository, error) {
	var result []models.GithubRepository
	var ids []int
	err := f.db.Select(&ids, "SELECT * favorites where userId =$1", userId)
	if err != nil {
		return nil, err
	}
	for _, value := range ids {
		var item models.GithubRepository
		err := f.db.Get(&item, "SELECT * FROM repositories where id=$1", value)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func (f FavoritesRepository) PutToFavorites(user int, id int) error {
	f.db.QueryRow("INSERT INTO favorites(user_id, repo_id) VALUES ($1,$2)", user, id)
	return nil
}

func (f FavoritesRepository) RemoveFromFavorites(user int, id int) error {
	f.db.QueryRow("DELETE FROM favorites WHERE repo_id=$1 and user_id=$2", id, user)
	return nil
}
