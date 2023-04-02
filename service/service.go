package service

import "simpleapi/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	repository.BookRepo
}

func NewService(repo repository.RepoInterface) *Service {
	return &Service{repo: repo}
}
