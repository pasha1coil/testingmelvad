package repository

import (
	"testingMelvad/internal/models"
)

type Tasks interface {
	Increment(incr *models.Incr) (int, error)
	CalculateHMAC(hmac *models.Hmac) (string, error)
	AddUser(user *models.Users) (string, error)
}

type Repository struct {
	tasks Tasks
}

func NewRepository(t Tasks) *Repository {
	return &Repository{
		tasks: t,
	}
}
