package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/gdpr"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/gdpr"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGdprLogUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.GdprLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewGdprLogUsecase(mockRepo, timeout)

	gdprID := primitive.NewObjectID().Hex()

	updatedGdprLog := domain.GdprLog{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:     primitive.NewObjectID(),
		ConsentID:      primitive.NewObjectID(),
		CustomerInfo:   "Jane Doe, jane.doe@example.com",
		RequestTypeID:  2,
		RequestDetails: "Request to export personal data.",
		CreatedOnUtc:   time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, gdprID).Return(updatedGdprLog, nil)

	result, err := usecase.FetchByID(context.Background(), gdprID)

	assert.NoError(t, err)
	assert.Equal(t, updatedGdprLog, result)
	mockRepo.AssertExpectations(t)
}

func TestGdprLogUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.GdprLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewGdprLogUsecase(mockRepo, timeout)

	newGdprLog := &domain.GdprLog{
		CustomerID:     primitive.NewObjectID(),
		ConsentID:      primitive.NewObjectID(),
		CustomerInfo:   "John Doe, john.doe@example.com",
		RequestTypeID:  1,
		RequestDetails: "Request to delete personal data.",
		CreatedOnUtc:   time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newGdprLog).Return(nil)

	err := usecase.Create(context.Background(), newGdprLog)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprLogUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.GdprLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewGdprLogUsecase(mockRepo, timeout)

	updatedGdprLog := &domain.GdprLog{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:     primitive.NewObjectID(),
		ConsentID:      primitive.NewObjectID(),
		CustomerInfo:   "Jane Doe, jane.doe@example.com",
		RequestTypeID:  2,
		RequestDetails: "Request to export personal data.",
		CreatedOnUtc:   time.Now().AddDate(0, 0, -7), // Created 7 days ago

	}

	mockRepo.On("Update", mock.Anything, updatedGdprLog).Return(nil)

	err := usecase.Update(context.Background(), updatedGdprLog)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprLogUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.GdprLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewGdprLogUsecase(mockRepo, timeout)

	gdprID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, gdprID).Return(nil)

	err := usecase.Delete(context.Background(), gdprID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprLogUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.GdprLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewGdprLogUsecase(mockRepo, timeout)

	fetchedGdprLogs := []domain.GdprLog{
		{
			ID:             primitive.NewObjectID(),
			CustomerID:     primitive.NewObjectID(),
			ConsentID:      primitive.NewObjectID(),
			CustomerInfo:   "John Doe, john.doe@example.com",
			RequestTypeID:  1,
			RequestDetails: "Request to delete personal data.",
			CreatedOnUtc:   time.Now().AddDate(0, 0, -10), // Created 10 days ago

		},
		{
			ID:             primitive.NewObjectID(),
			CustomerID:     primitive.NewObjectID(),
			ConsentID:      primitive.NewObjectID(),
			CustomerInfo:   "Jane Doe, jane.doe@example.com",
			RequestTypeID:  2,
			RequestDetails: "Request to export personal data.",
			CreatedOnUtc:   time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedGdprLogs, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedGdprLogs, result)
	mockRepo.AssertExpectations(t)
}
