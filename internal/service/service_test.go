package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/pasha1coil/testingmelvad/internal/models"
	mock_service "github.com/pasha1coil/testingmelvad/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrement_Good(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTask := mock_service.NewMockTasks(ctrl)
	expectedResult := int64(20)

	mockTask.EXPECT().Increment(&models.Incr{
		Key:   "Age",
		Value: 19,
	}).Return(expectedResult, nil)

	service := &TasksService{
		repo: mockTask,
	}
	result, err := service.Increment(&models.Incr{
		Key:   "Age",
		Value: 19,
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestIncrement_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTask := mock_service.NewMockTasks(ctrl)

	expectedError := errors.New("error")
	mockTask.EXPECT().Increment(&models.Incr{
		Key:   "Age",
		Value: 19,
	}).Return(int64(0), expectedError)

	service := &TasksService{
		repo: mockTask,
	}
	_, err := service.Increment(&models.Incr{
		Key:   "Age",
		Value: 19,
	})

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestCalculateHMAC_Good(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTask := mock_service.NewMockTasks(ctrl)
	expectedResult := "55b5bb82607f1c64d187d1089b6cc27e1667cb24de14bb26f8a0ebc3b3dc7595096290b35d7df477a57e69059f2a00946f1737d3eaef3e6c73fa29ac400b8bdb"

	mockTask.EXPECT().CalculateHMAC(&models.Hash{
		Text: "test",
		Key:  "test123",
	}).Return(expectedResult, nil)

	service := &TasksService{
		repo: mockTask,
	}
	result, err := service.CalculateHMAC(&models.Hash{
		Text: "test",
		Key:  "test123",
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestCalculateHMAC_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTask := mock_service.NewMockTasks(ctrl)
	expectedError := errors.New("error")

	mockTask.EXPECT().CalculateHMAC(&models.Hash{
		Text: "test",
		Key:  "test123",
	}).Return("", expectedError)

	service := &TasksService{
		repo: mockTask,
	}
	_, err := service.CalculateHMAC(&models.Hash{
		Text: "test",
		Key:  "test123",
	})

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestAddUser_Good(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTask := mock_service.NewMockTasks(ctrl)
	expectedResult := 1

	mockTask.EXPECT().AddUser(&models.Users{
		Name: "Alex",
		Age:  21,
	}).Return(expectedResult, nil)

	service := &TasksService{
		repo: mockTask,
	}
	result, err := service.AddUser(&models.Users{
		Name: "Alex",
		Age:  21,
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestAddUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTask := mock_service.NewMockTasks(ctrl)
	expectedError := errors.New("error")

	mockTask.EXPECT().AddUser(&models.Users{
		Name: "Alex",
		Age:  21,
	}).Return(0, expectedError)

	service := &TasksService{
		repo: mockTask,
	}
	_, err := service.AddUser(&models.Users{
		Name: "Alex",
		Age:  21,
	})

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
