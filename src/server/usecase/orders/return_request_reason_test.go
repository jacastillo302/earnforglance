package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReturnRequestReasonUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestReasonRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewReturnRequestReasonUsecase(mockRepo, timeout)

	returnRequestReasonID := primitive.NewObjectID().Hex()

	updatedReturnRequestReason := domain.ReturnRequestReason{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Wrong Item Delivered",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, returnRequestReasonID).Return(updatedReturnRequestReason, nil)

	result, err := usecase.FetchByID(context.Background(), returnRequestReasonID)

	assert.NoError(t, err)
	assert.Equal(t, updatedReturnRequestReason, result)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestReasonUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestReasonRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewReturnRequestReasonUsecase(mockRepo, timeout)

	newReturnRequestReason := &domain.ReturnRequestReason{
		Name:         "Defective Item",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newReturnRequestReason).Return(nil)

	err := usecase.Create(context.Background(), newReturnRequestReason)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestReasonUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestReasonRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewReturnRequestReasonUsecase(mockRepo, timeout)

	updatedReturnRequestReason := &domain.ReturnRequestReason{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Wrong Item Delivered",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedReturnRequestReason).Return(nil)

	err := usecase.Update(context.Background(), updatedReturnRequestReason)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestReasonUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestReasonRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewReturnRequestReasonUsecase(mockRepo, timeout)

	returnRequestReasonID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, returnRequestReasonID).Return(nil)

	err := usecase.Delete(context.Background(), returnRequestReasonID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestReasonUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestReasonRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewReturnRequestReasonUsecase(mockRepo, timeout)

	fetchedReturnRequestReasons := []domain.ReturnRequestReason{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Defective Item",
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Wrong Item Delivered",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedReturnRequestReasons, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedReturnRequestReasons, result)
	mockRepo.AssertExpectations(t)
}
