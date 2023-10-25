package service

import (
	"github.com/pasha1coil/testingmelvad/internal/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Tasks interface {
	Increment(incr *models.Incr) (int64, error)
	CalculateHMAC(hmac *models.Hash) (string, error)
	AddUser(user *models.Users) (int, error)
}

func NewTasksService(repo Tasks) *TasksService {
	return &TasksService{repo: repo}
}

type TasksService struct {
	repo Tasks
}

func (s *TasksService) Increment(incr *models.Incr) (int64, error) {
	return s.repo.Increment(incr)
}

func (s *TasksService) CalculateHMAC(hmac *models.Hash) (string, error) {
	return s.repo.CalculateHMAC(hmac)
}

func (s *TasksService) AddUser(user *models.Users) (int, error) {
	return s.repo.AddUser(user)
}
