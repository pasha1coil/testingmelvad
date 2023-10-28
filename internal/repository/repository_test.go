package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pasha1coil/testingmelvad/internal/models"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestCalculateHMAC(t *testing.T) {
	db, _, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	redisClient := redis.NewClient(&redis.Options{})
	repo := NewTasksRepo(db, redisClient)
	hash := &models.Hash{
		Text: "testtext",
		Key:  "testkey",
	}

	_, err := repo.CalculateHMAC(hash)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

//func TestCreateTable(t *testing.T) {
//	db, mock, _ := sqlmock.New()
//	mock.ExpectExec("CREATE TABLE IF NOT EXISTS").WillReturnResult(sqlmock.NewResult(1, 1))
//
//	redisClient := redis.NewClient(&redis.Options{})
//	repo := NewTasksRepo(db, redisClient)
//
//	err := repo.CreateTable()
//	if err != nil {
//		t.Errorf("Expected no error, got %v", err)
//	}
//}
