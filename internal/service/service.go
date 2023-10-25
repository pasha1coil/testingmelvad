package service

import (
	"testingMelvad/internal/models"
	"testingMelvad/internal/repository"
)

type Tasks interface {
	Increment(incr *models.Incr) (int, error)
	CalculateHMAC(hmac *models.Hmac) (string, error)
	AddUser(user *models.Users) (string, error)
}

type Service struct {
	tasks Tasks
}

func NewService(t Tasks) *Service {
	return &Service{
		tasks: t,
	}
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{repo: repo}
}

type TasksService struct {
	repo repository.Tasks
}

func (s *TasksService) Increment(incr *models.Incr) (int, error) {
	return s.repo.Increment(incr)
}

func (s *TasksService) CalculateHMAC(hmac *models.Hmac) (string, error) {
	return s.repo.CalculateHMAC(hmac)
}

func (s *TasksService) AddUser(user *models.Users) (string, error) {
	return s.repo.AddUser(user)
}
