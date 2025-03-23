package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/logging"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/logging"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestLogUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.LogRepository)
	timeout := time.Duration(10)
	usecase := test.NewLogUsecase(mockRepo, timeout)

	logID := primitive.NewObjectID().Hex()
	updatedLog := domain.Log{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		LogLevelID:   2,
		ShortMessage: "Application error",
		FullMessage:  "An error occurred while processing the request.",
		IpAddress:    "192.168.1.2",
		CustomerID:   new(primitive.ObjectID),
		PageUrl:      "/error",
		ReferrerUrl:  "/home",
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		LogLevel:     1,
	}

	mockRepo.On("FetchByID", mock.Anything, logID).Return(updatedLog, nil)

	result, err := usecase.FetchByID(context.Background(), logID)

	assert.NoError(t, err)
	assert.Equal(t, updatedLog, result)
	mockRepo.AssertExpectations(t)
}

func TestLogUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.LogRepository)
	timeout := time.Duration(10)
	usecase := test.NewLogUsecase(mockRepo, timeout)

	newLog := &domain.Log{
		LogLevelID:   1,
		ShortMessage: "Application started",
		FullMessage:  "The application has successfully started.",
		IpAddress:    "192.168.1.1",
		CustomerID:   nil,
		PageUrl:      "/home",
		ReferrerUrl:  "",
		CreatedOnUtc: time.Now(),
		LogLevel:     2,
	}

	mockRepo.On("Create", mock.Anything, newLog).Return(nil)

	err := usecase.Create(context.Background(), newLog)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLogUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.LogRepository)
	timeout := time.Duration(10)
	usecase := test.NewLogUsecase(mockRepo, timeout)

	updatedLog := &domain.Log{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		LogLevelID:   2,
		ShortMessage: "Application error",
		FullMessage:  "An error occurred while processing the request.",
		IpAddress:    "192.168.1.2",
		CustomerID:   new(primitive.ObjectID),
		PageUrl:      "/error",
		ReferrerUrl:  "/home",
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		LogLevel:     1,
	}
	*updatedLog.CustomerID = primitive.NewObjectID()

	mockRepo.On("Update", mock.Anything, updatedLog).Return(nil)

	err := usecase.Update(context.Background(), updatedLog)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLogUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.LogRepository)
	timeout := time.Duration(10)
	usecase := test.NewLogUsecase(mockRepo, timeout)

	logID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, logID).Return(nil)

	err := usecase.Delete(context.Background(), logID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLogUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.LogRepository)
	timeout := time.Duration(10)
	usecase := test.NewLogUsecase(mockRepo, timeout)

	fetchedLogs := []domain.Log{
		{
			ID:           primitive.NewObjectID(),
			LogLevelID:   1,
			ShortMessage: "Application started",
			FullMessage:  "The application has successfully started.",
			IpAddress:    "192.168.1.1",
			CustomerID:   nil,
			PageUrl:      "/home",
			ReferrerUrl:  "",
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
			LogLevel:     2,
		},
		{
			ID:           primitive.NewObjectID(),
			LogLevelID:   2,
			ShortMessage: "Application error",
			FullMessage:  "An error occurred while processing the request.",
			IpAddress:    "192.168.1.2",
			CustomerID:   new(primitive.ObjectID),
			PageUrl:      "/error",
			ReferrerUrl:  "/home",
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
			LogLevel:     3,
		},
	}
	*fetchedLogs[1].CustomerID = primitive.NewObjectID()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedLogs, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedLogs, result)
	mockRepo.AssertExpectations(t)
}
