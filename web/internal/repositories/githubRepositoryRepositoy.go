package repositories

import (
	"coursework/web/internal/models"
	"github.com/jmoiron/sqlx"
)

type GithubRepository struct {
	db *sqlx.DB
}

func NewGithubRepository(db *sqlx.DB) *GithubRepository {
	return &GithubRepository{db: db}
}

func (g GithubRepository) GetAll() ([]models.GithubRepository, error) {
	var result []models.GithubRepository
	if err := g.db.Select(&result, "SELECT * FROM repositories"); err != nil {
		return nil, err
	}
	return result, nil
}

func (g GithubRepository) GetRepoDataFromLocalBase(id int) (models.GithubRepository, error) {
	var result models.GithubRepository
	if err := g.db.Get(&result, "select * from repositories where id=$1", id); err != nil {
		return models.GithubRepository{}, err
	}
	return result, nil
}

func (g GithubRepository) SetRepoData(repo models.GithubRepository) (int, error) {
	g.db.QueryRow("INSERT INTO repositories(id, fullName, isPrivate,  url, descr, canFork, created, updated, pushed, size,  language, forks, issues, watchers, subscribers) "+
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id", repo.Id, repo.FullName, repo.IsPrivate, repo.RepositoryUrl, repo.Description, repo.CanFork, repo.CreatedAt, repo.UpdatedAt, repo.PushedAt, repo.Size, repo.Language, repo.Forks, repo.Issues, repo.Watchers, repo.Subscribers)
	return 0, nil
}
