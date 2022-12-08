package services

import (
	"coursework/web/internal/broker"
	"coursework/web/internal/models"
	"coursework/web/internal/repositories"
)

type RepositoriesService struct {
	repo   repositories.GithubRepositoryRepo
	broker *broker.Broker
}

func NewRepositoriesService(repo repositories.GithubRepositoryRepo, broker *broker.Broker) *RepositoriesService {
	return &RepositoriesService{repo: repo, broker: broker}
}

func (r RepositoriesService) GetAll() ([]models.GithubRepository, error) {
	return r.repo.GetAll()
}

func (r RepositoriesService) GetRepoDataFromLocalBase(id string) (models.GithubRepository, error) {
	row, err := r.repo.GetRepoDataFromLocalBase(id)
	if err != nil {
		return models.GithubRepository{}, err
	}
	if err := r.broker.Publish(map[string]interface{}{"repository": row.RepositoryUrl}); err != nil {
		return row, nil
	}
	return r.repo.GetRepoDataFromLocalBase(id)
}

func (r RepositoriesService) SetRepoData(repo models.GithubRepository) (int, error) {
	return r.repo.SetRepoData(repo)
}
