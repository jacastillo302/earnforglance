package usecase

import (
	"context"
	"testing"
	"time"

	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPermissionRecordCustomerRoleMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	mappingID := primitive.NewObjectID().Hex()

	updatedPermissionRecordCustomerRoleMapping := domain.PermissionRecordCustomerRoleMapping{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: primitive.NewObjectID(),
		CustomerRoleID:     primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, mappingID).Return(updatedPermissionRecordCustomerRoleMapping, nil)

	result, err := usecase.FetchByID(context.Background(), mappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPermissionRecordCustomerRoleMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	newPermissionRecordCustomerRoleMapping := &domain.PermissionRecordCustomerRoleMapping{
		PermissionRecordID: primitive.NewObjectID(),
		CustomerRoleID:     primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newPermissionRecordCustomerRoleMapping).Return(nil)

	err := usecase.Create(context.Background(), newPermissionRecordCustomerRoleMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	updatedPermissionRecordCustomerRoleMapping := &domain.PermissionRecordCustomerRoleMapping{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: primitive.NewObjectID(),
		CustomerRoleID:     primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedPermissionRecordCustomerRoleMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedPermissionRecordCustomerRoleMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	mappingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, mappingID).Return(nil)

	err := usecase.Delete(context.Background(), mappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	fetchedPermissionRecordCustomerRoleMappings := []domain.PermissionRecordCustomerRoleMapping{
		{
			ID:                 primitive.NewObjectID(),
			PermissionRecordID: primitive.NewObjectID(),
			CustomerRoleID:     primitive.NewObjectID(),
		},
		{
			ID:                 primitive.NewObjectID(),
			PermissionRecordID: primitive.NewObjectID(),
			CustomerRoleID:     primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPermissionRecordCustomerRoleMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPermissionRecordCustomerRoleMappings, result)
	mockRepo.AssertExpectations(t)
}
