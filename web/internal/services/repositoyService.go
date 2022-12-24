package services

import (
	"coursework/web/internal/broker"
	"coursework/web/internal/models"
	"coursework/web/internal/repositories"
	"github.com/siruspen/logrus"
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

func (r RepositoriesService) GetRepoDataFromLocalBase(id int) (models.GithubRepository, error) {
	row, _ := r.repo.GetRepoDataFromLocalBase(id)
	if err := r.broker.Publish(map[string]interface{}{"repository": row.RepositoryUrl}); err != nil {
		return row, nil
	}
	return r.repo.GetRepoDataFromLocalBase(id)
}

func (r RepositoriesService) SetRepoData(repo models.GithubRepository) (int, error) {
	return r.repo.SetRepoData(repo)
}

func (r RepositoriesService) AddRepoData(url string) {
	if err := r.broker.Publish(map[string]interface{}{"repository": url}); err != nil {
		logrus.Info(err)
	}
}
