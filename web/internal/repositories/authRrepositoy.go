package repositories

import (
	"coursework/web/internal/models"
	"github.com/jmoiron/sqlx"
	"sync"
)

type AuthRepository struct {
	sync.RWMutex
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a AuthRepository) CreateUser(user models.User) (int, error) {
	var id int
	row := a.db.QueryRow("INSERT INTO users(id, login, password) values ($1,$2,$3) RETURNING id", user.Id, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a AuthRepository) GetUser(login, password string) (models.User, error) {
	var user models.User
	if err := a.db.Get(&user, "SELECT * FROM users WHERE login=$1 and password=$2", login, password); err != nil {
		return models.User{}, err
	}
	return user, nil
}
