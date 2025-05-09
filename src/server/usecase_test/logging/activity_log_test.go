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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestActivityLogUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ActivityLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogUsecase(mockRepo, timeout)

	loggingID := bson.NewObjectID().Hex()

	updatedActivityLog := domain.ActivityLog{
		ID:                bson.NewObjectID(), // Existing ID of the record to update
		ActivityLogTypeID: bson.NewObjectID(),
		EntityID:          new(bson.ObjectID),
		EntityName:        "Product",
		CustomerID:        bson.NewObjectID(),
		Comment:           "Customer viewed a product.",
		CreatedOnUtc:      time.Now().AddDate(0, 0, -7), // Created 7 days ago
		IpAddress:         "192.168.1.2",
	}
	*updatedActivityLog.EntityID = bson.NewObjectID()

	mockRepo.On("FetchByID", mock.Anything, loggingID).Return(updatedActivityLog, nil)

	result, err := usecase.FetchByID(context.Background(), loggingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedActivityLog, result)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ActivityLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogUsecase(mockRepo, timeout)

	newActivityLog := &domain.ActivityLog{
		ActivityLogTypeID: bson.NewObjectID(),
		EntityID:          nil,
		EntityName:        "Order",
		CustomerID:        bson.NewObjectID(),
		Comment:           "Customer placed an order.",
		CreatedOnUtc:      time.Now(),
		IpAddress:         "192.168.1.1",
	}

	mockRepo.On("Create", mock.Anything, newActivityLog).Return(nil)

	err := usecase.Create(context.Background(), newActivityLog)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ActivityLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogUsecase(mockRepo, timeout)

	updatedActivityLog := &domain.ActivityLog{
		ID:                bson.NewObjectID(), // Existing ID of the record to update
		ActivityLogTypeID: bson.NewObjectID(),
		EntityID:          new(bson.ObjectID),
		EntityName:        "Product",
		CustomerID:        bson.NewObjectID(),
		Comment:           "Customer viewed a product.",
		CreatedOnUtc:      time.Now().AddDate(0, 0, -7), // Created 7 days ago
		IpAddress:         "192.168.1.2",
	}
	*updatedActivityLog.EntityID = bson.NewObjectID()

	mockRepo.On("Update", mock.Anything, updatedActivityLog).Return(nil)

	err := usecase.Update(context.Background(), updatedActivityLog)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ActivityLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogUsecase(mockRepo, timeout)

	loggingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, loggingID).Return(nil)

	err := usecase.Delete(context.Background(), loggingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ActivityLogRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogUsecase(mockRepo, timeout)

	fetchedActivityLogs := []domain.ActivityLog{
		{
			ID:                bson.NewObjectID(),
			ActivityLogTypeID: bson.NewObjectID(),
			EntityID:          nil,
			EntityName:        "Order",
			CustomerID:        bson.NewObjectID(),
			Comment:           "Customer placed an order.",
			CreatedOnUtc:      time.Now().AddDate(0, 0, -10), // Created 10 days ago
			IpAddress:         "192.168.1.1",
		},
		{
			ID:                bson.NewObjectID(),
			ActivityLogTypeID: bson.NewObjectID(),
			EntityID:          new(bson.ObjectID),
			EntityName:        "Product",
			CustomerID:        bson.NewObjectID(),
			Comment:           "Customer viewed a product.",
			CreatedOnUtc:      time.Now().AddDate(0, 0, -5), // Created 5 days ago
			IpAddress:         "192.168.1.2",
		},
	}
	*fetchedActivityLogs[1].EntityID = bson.NewObjectID()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedActivityLogs, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedActivityLogs, result)
	mockRepo.AssertExpectations(t)
}
