package usecase

import (
	"context"
	"testing"
	"time"

	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReturnRequestActionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestActionRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestActionUsecase(mockRepo, timeout)

	returnRequestActionID := primitive.NewObjectID().Hex()

	updatedReturnRequestAction := domain.ReturnRequestAction{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Refund Item",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, returnRequestActionID).Return(updatedReturnRequestAction, nil)

	result, err := usecase.FetchByID(context.Background(), returnRequestActionID)

	assert.NoError(t, err)
	assert.Equal(t, updatedReturnRequestAction, result)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestActionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestActionRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestActionUsecase(mockRepo, timeout)

	newReturnRequestAction := &domain.ReturnRequestAction{
		Name:         "Replace Item",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newReturnRequestAction).Return(nil)

	err := usecase.Create(context.Background(), newReturnRequestAction)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestActionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestActionRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestActionUsecase(mockRepo, timeout)

	updatedReturnRequestAction := &domain.ReturnRequestAction{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Refund Item",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedReturnRequestAction).Return(nil)

	err := usecase.Update(context.Background(), updatedReturnRequestAction)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestActionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestActionRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestActionUsecase(mockRepo, timeout)

	returnRequestActionID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, returnRequestActionID).Return(nil)

	err := usecase.Delete(context.Background(), returnRequestActionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReturnRequestActionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ReturnRequestActionRepository)
	timeout := time.Duration(10)
	usecase := NewReturnRequestActionUsecase(mockRepo, timeout)

	fetchedReturnRequestActions := []domain.ReturnRequestAction{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Replace Item",
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Refund Item",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedReturnRequestActions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedReturnRequestActions, result)
	mockRepo.AssertExpectations(t)
}
