package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerCustomerRoleMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerCustomerRoleMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerCustomerRoleMappingUsecase(mockRepo, timeout)

	customerCustomerRoleMappingID := primitive.NewObjectID().Hex()

	updatedCustomerCustomerRoleMapping := domain.CustomerCustomerRoleMapping{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:     primitive.NewObjectID(),
		CustomerRoleID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, customerCustomerRoleMappingID).Return(updatedCustomerCustomerRoleMapping, nil)

	result, err := usecase.FetchByID(context.Background(), customerCustomerRoleMappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerCustomerRoleMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerCustomerRoleMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerCustomerRoleMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerCustomerRoleMappingUsecase(mockRepo, timeout)

	newCustomerCustomerRoleMapping := &domain.CustomerCustomerRoleMapping{
		CustomerID:     primitive.NewObjectID(),
		CustomerRoleID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newCustomerCustomerRoleMapping).Return(nil)

	err := usecase.Create(context.Background(), newCustomerCustomerRoleMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerCustomerRoleMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerCustomerRoleMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerCustomerRoleMappingUsecase(mockRepo, timeout)
	updatedCustomerCustomerRoleMapping := &domain.CustomerCustomerRoleMapping{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:     primitive.NewObjectID(),
		CustomerRoleID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedCustomerCustomerRoleMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerCustomerRoleMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerCustomerRoleMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerCustomerRoleMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerCustomerRoleMappingUsecase(mockRepo, timeout)

	customerCustomerRoleMappingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerCustomerRoleMappingID).Return(nil)

	err := usecase.Delete(context.Background(), customerCustomerRoleMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerCustomerRoleMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerCustomerRoleMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerCustomerRoleMappingUsecase(mockRepo, timeout)

	fetchedCustomerCustomerRoleMappings := []domain.CustomerCustomerRoleMapping{
		{
			ID:             primitive.NewObjectID(),
			CustomerID:     primitive.NewObjectID(),
			CustomerRoleID: primitive.NewObjectID(),
		},
		{
			ID:             primitive.NewObjectID(),
			CustomerID:     primitive.NewObjectID(),
			CustomerRoleID: primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerCustomerRoleMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerCustomerRoleMappings, result)
	mockRepo.AssertExpectations(t)
}
