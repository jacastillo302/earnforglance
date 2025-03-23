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

func TestCustomerPermanentlyDeletedUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerPermanentlyDeletedRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerPermanentlyDeletedUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	updatedCustomerPermanentlyDeleted := domain.CustomerPermanentlyDeleted{
		CustomerID: primitive.NewObjectID(), // Existing CustomerID to update
		Email:      "updated_deleted_customer@example.com",
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedCustomerPermanentlyDeleted, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerPermanentlyDeleted, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPermanentlyDeletedUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerPermanentlyDeletedRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerPermanentlyDeletedUsecase(mockRepo, timeout)
	newCustomerPermanentlyDeleted := &domain.CustomerPermanentlyDeleted{
		CustomerID: primitive.NewObjectID(),
		Email:      "deleted_customer@example.com",
	}

	mockRepo.On("Create", mock.Anything, newCustomerPermanentlyDeleted).Return(nil)

	err := usecase.Create(context.Background(), newCustomerPermanentlyDeleted)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPermanentlyDeletedUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerPermanentlyDeletedRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerPermanentlyDeletedUsecase(mockRepo, timeout)

	updatedCustomerPermanentlyDeleted := &domain.CustomerPermanentlyDeleted{
		CustomerID: primitive.NewObjectID(), // Existing CustomerID to update
		Email:      "updated_deleted_customer@example.com",
	}

	mockRepo.On("Update", mock.Anything, updatedCustomerPermanentlyDeleted).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerPermanentlyDeleted)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPermanentlyDeletedUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerPermanentlyDeletedRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerPermanentlyDeletedUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerPermanentlyDeletedUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerPermanentlyDeletedRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerPermanentlyDeletedUsecase(mockRepo, timeout)

	fetchedCustomerPermanentlyDeleted := []domain.CustomerPermanentlyDeleted{
		{
			CustomerID: primitive.NewObjectID(),
			Email:      "deleted_customer1@example.com",
		},
		{
			CustomerID: primitive.NewObjectID(),
			Email:      "deleted_customer2@example.com",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerPermanentlyDeleted, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerPermanentlyDeleted, result)
	mockRepo.AssertExpectations(t)
}
