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
	var id int
	if g.contains(repo) {
		return 99999, nil
	}
	row := g.db.QueryRow("INSERT INTO repositories(id, full_name, is_private,  url, descr, can_fork, created, updated, pushed, size,  language, forks, issues, watchers, subscribers) "+
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id", repo.Id, repo.FullName, repo.IsPrivate, repo.RepositoryUrl, repo.Description, repo.CanFork, repo.CreatedAt, repo.UpdatedAt, repo.PushedAt, repo.Size, repo.Language, repo.Forks, repo.Issues, repo.Watchers, repo.Subscribers)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (g GithubRepository) contains(repo models.GithubRepository) bool {
	var res []models.GithubRepository
	err := g.db.Select(&res, "SELECT * FROM repositories")
	if err != nil {
		return false
	}
	for _, val := range res {
		if val.Id == repo.Id {
			return true
		}
	}

	return false
}

func (g GithubRepository) GetByUrl(url string) (int, error) {
	var id int
	err := g.db.Get(&id, "SELECT id FROM repositories WHERE url=$1", url)
	if err != nil {
		return 0, err
	}

	return id, nil
}
