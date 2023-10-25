package repository

import (
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
	"testingMelvad/internal/models"
)

var ctx = context.Background()

type TasksRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewTasksRepo(db *sql.DB, rdb *redis.Client) *TasksRepo {
	return &TasksRepo{db: db, rdb: rdb}
}

func (r *TasksRepo) Increment(incr *models.Incr) (int, error) {
	return 0, nil
}

func (r *TasksRepo) CalculateHMAC(hmac *models.Hmac) (string, error) {
	return "", nil
}

func (r *TasksRepo) AddUser(user *models.Users) (string, error) {
	return "", nil
}
