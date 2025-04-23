package usecase_test

import (
	"context"
	"testing"
	"time"

	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"
	test "earnforglance/server/usecase/security"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPermissionRecordCustomerRoleMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	mappingID := bson.NewObjectID().Hex()

	updatedPermissionRecordCustomerRoleMapping := domain.PermissionRecordCustomerRoleMapping{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: bson.NewObjectID(),
		CustomerRoleID:     bson.NewObjectID(),
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
	usecase := test.NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	newPermissionRecordCustomerRoleMapping := &domain.PermissionRecordCustomerRoleMapping{
		PermissionRecordID: bson.NewObjectID(),
		CustomerRoleID:     bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newPermissionRecordCustomerRoleMapping).Return(nil)

	err := usecase.Create(context.Background(), newPermissionRecordCustomerRoleMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	updatedPermissionRecordCustomerRoleMapping := &domain.PermissionRecordCustomerRoleMapping{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		PermissionRecordID: bson.NewObjectID(),
		CustomerRoleID:     bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedPermissionRecordCustomerRoleMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedPermissionRecordCustomerRoleMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	mappingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, mappingID).Return(nil)

	err := usecase.Delete(context.Background(), mappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordCustomerRoleMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordCustomerRoleMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPermissionRecordCustomerRoleMappingUsecase(mockRepo, timeout)

	fetchedPermissionRecordCustomerRoleMappings := []domain.PermissionRecordCustomerRoleMapping{
		{
			ID:                 bson.NewObjectID(),
			PermissionRecordID: bson.NewObjectID(),
			CustomerRoleID:     bson.NewObjectID(),
		},
		{
			ID:                 bson.NewObjectID(),
			PermissionRecordID: bson.NewObjectID(),
			CustomerRoleID:     bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPermissionRecordCustomerRoleMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPermissionRecordCustomerRoleMappings, result)
	mockRepo.AssertExpectations(t)
}
