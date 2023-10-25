package repository

import (
	"github.com/pasha1coil/testingmelvad/internal/models"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Tasks interface {
	Increment(incr *models.Incr) (int64, error)
	CalculateHMAC(hmac *models.Hash) (string, error)
	AddUser(user *models.Users) (int, error)
}

type Repository struct {
	tasks Tasks
}
