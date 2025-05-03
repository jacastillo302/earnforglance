package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/attributes"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/attributes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPermisionRecordAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeValueUsecase(mockRepo, timeout)

	customerID := bson.NewObjectID().Hex()

	updatedPermisionRecordAttributeValue := domain.PermisionRecordAttributeValue{
		ID:                         bson.NewObjectID(), // Existing ID of the record to update
		PermisionRecordAttributeID: bson.NewObjectID(),
		Value:                      "Preferred Currency",
		IsPreSelected:              false,
		DisplayOrder:               2,
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedPermisionRecordAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPermisionRecordAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeValueUsecase(mockRepo, timeout)

	newPermisionRecordAttributeValue := &domain.PermisionRecordAttributeValue{
		PermisionRecordAttributeID: bson.NewObjectID(),
		Value:                      "Preferred Language",
		IsPreSelected:              true,
		DisplayOrder:               1,
	}

	mockRepo.On("Create", mock.Anything, newPermisionRecordAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newPermisionRecordAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeValueUsecase(mockRepo, timeout)

	updatedPermisionRecordAttributeValue := &domain.PermisionRecordAttributeValue{
		ID:                         bson.NewObjectID(), // Existing ID of the record to update
		PermisionRecordAttributeID: bson.NewObjectID(),
		Value:                      "Preferred Currency",
		IsPreSelected:              false,
		DisplayOrder:               2,
	}

	mockRepo.On("Update", mock.Anything, updatedPermisionRecordAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedPermisionRecordAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeValueUsecase(mockRepo, timeout)

	customerID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermisionRecordAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PermisionRecordAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermisionRecordAttributeValueUsecase(mockRepo, timeout)

	fetchedPermisionRecordAttributeValues := []domain.PermisionRecordAttributeValue{
		{
			ID:                         bson.NewObjectID(),
			PermisionRecordAttributeID: bson.NewObjectID(),
			Value:                      "Preferred Language",
			IsPreSelected:              true,
			DisplayOrder:               1,
		},
		{
			ID:                         bson.NewObjectID(),
			PermisionRecordAttributeID: bson.NewObjectID(),
			Value:                      "Preferred Currency",
			IsPreSelected:              false,
			DisplayOrder:               2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPermisionRecordAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPermisionRecordAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
