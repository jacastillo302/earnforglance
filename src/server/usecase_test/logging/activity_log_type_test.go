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

func TestActivityLogTypeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ActivityLogTypeRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogTypeUsecase(mockRepo, timeout)

	activityLogTypeID := bson.NewObjectID().Hex()

	updatedActivityLogType := domain.ActivityLogType{
		ID:            bson.NewObjectID(), // Existing ID of the record to update
		SystemKeyword: "customer_registration",
		Name:          "Customer Registration",
		Enabled:       false,
	}

	mockRepo.On("FetchByID", mock.Anything, activityLogTypeID).Return(updatedActivityLogType, nil)

	result, err := usecase.FetchByID(context.Background(), activityLogTypeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedActivityLogType, result)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogTypeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ActivityLogTypeRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogTypeUsecase(mockRepo, timeout)

	newActivityLogType := &domain.ActivityLogType{
		SystemKeyword: "customer_login",
		Name:          "Customer Login",
		Enabled:       true,
	}

	mockRepo.On("Create", mock.Anything, newActivityLogType).Return(nil)

	err := usecase.Create(context.Background(), newActivityLogType)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogTypeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ActivityLogTypeRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogTypeUsecase(mockRepo, timeout)

	updatedActivityLogType := &domain.ActivityLogType{
		ID:            bson.NewObjectID(), // Existing ID of the record to update
		SystemKeyword: "customer_registration",
		Name:          "Customer Registration",
		Enabled:       false,
	}

	mockRepo.On("Update", mock.Anything, updatedActivityLogType).Return(nil)

	err := usecase.Update(context.Background(), updatedActivityLogType)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogTypeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ActivityLogTypeRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogTypeUsecase(mockRepo, timeout)

	activityLogTypeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, activityLogTypeID).Return(nil)

	err := usecase.Delete(context.Background(), activityLogTypeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActivityLogTypeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ActivityLogTypeRepository)
	timeout := time.Duration(10)
	usecase := test.NewActivityLogTypeUsecase(mockRepo, timeout)

	fetchedActivityLogTypes := []domain.ActivityLogType{
		{
			ID:            bson.NewObjectID(),
			SystemKeyword: "customer_login",
			Name:          "Customer Login",
			Enabled:       true,
		},
		{
			ID:            bson.NewObjectID(),
			SystemKeyword: "customer_registration",
			Name:          "Customer Registration",
			Enabled:       false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedActivityLogTypes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedActivityLogTypes, result)
	mockRepo.AssertExpectations(t)
}
