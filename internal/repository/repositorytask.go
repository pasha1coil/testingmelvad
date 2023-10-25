package repository

import (
	"context"
	"crypto/hmac"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/pasha1coil/testingmelvad/internal/models"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

var ctx = context.Background()

type TasksRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewTasksRepo(db *sql.DB, rdb *redis.Client) *TasksRepo {
	return &TasksRepo{db: db, rdb: rdb}
}

func (r *TasksRepo) Increment(incr *models.Incr) (int64, error) {
	log.Infof("Starting Redis increment for key: %s with increment value: %d", incr.Key, incr.Value)
	//не совсем понятно зачем при value 19, увеличивать его на 1, чтобы в ответе было 20
	value, err := r.rdb.IncrBy(ctx, incr.Key, incr.Value+1).Result()
	if err != nil {
		errMsg := fmt.Sprintf("Error incrementing Redis key: '%s' with value: %d. Error: %s", incr.Key, incr.Value, err.Error())
		log.Error(errMsg)
		return 0, errors.New(errMsg)
	}
	return value, nil
}

func (r *TasksRepo) CalculateHMAC(hash *models.Hash) (string, error) {
	log.Infof("Starting Calculate HMACSHA512 for text: %s and key: %s", hash.Text, hash.Key)
	h := hmac.New(sha3.New512, []byte(hash.Key))
	_, err := h.Write([]byte(hash.Text))
	if err != nil {
		errMsg := fmt.Sprintf("Error writing data for HMAC calculation: %s", err.Error())
		return "", errors.New(errMsg)
	}
	signature := h.Sum(nil)
	return hex.EncodeToString(signature), nil
}

func (r *TasksRepo) AddUser(user *models.Users) (int, error) {
	log.Infof("Starting add user with name - %s and age - %d", user.Name, user.Age)
	err := r.CreateTable()
	if err != nil {
		return 0, err
	}
	var id int
	err = r.db.QueryRow("INSERT INTO users (name,age) values ($1,$2) RETURNING id", user.Name, user.Age).Scan(&id)
	if err != nil {
		errMsg := fmt.Sprintf("Error insert user in in the database")
		log.Errorf(errMsg)
		return 0, err
	}
	return id, nil
}

func (r *TasksRepo) CreateTable() error {
	log.Infoln("Start check or creating a table in the database")
	_, err := r.db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT)")
	if err != nil {
		errMsg := fmt.Sprintf("Error creating table users in db:%s", err.Error())
		log.Errorf(errMsg)
		return err
	}
	log.Infoln("A table was successfully verified or created")
	return nil
}
