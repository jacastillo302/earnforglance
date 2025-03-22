package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPermissionRecordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordUsecase(mockRepo, timeout)

	permissionRecordID := primitive.NewObjectID().Hex()

	updatedPermissionRecord := domain.PermissionRecord{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		Name:       "Manage Categories",
		SystemName: "manage_categories",
		Category:   "Catalog",
	}

	mockRepo.On("FetchByID", mock.Anything, permissionRecordID).Return(updatedPermissionRecord, nil)

	result, err := usecase.FetchByID(context.Background(), permissionRecordID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPermissionRecord, result)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordUsecase(mockRepo, timeout)

	newPermissionRecord := &domain.PermissionRecord{
		Name:       "Manage Products",
		SystemName: "manage_products",
		Category:   "Catalog",
	}

	mockRepo.On("Create", mock.Anything, newPermissionRecord).Return(nil)

	err := usecase.Create(context.Background(), newPermissionRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordUsecase(mockRepo, timeout)

	updatedPermissionRecord := &domain.PermissionRecord{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		Name:       "Manage Categories",
		SystemName: "manage_categories",
		Category:   "Catalog",
	}

	mockRepo.On("Update", mock.Anything, updatedPermissionRecord).Return(nil)

	err := usecase.Update(context.Background(), updatedPermissionRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordUsecase(mockRepo, timeout)

	permissionRecordID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, permissionRecordID).Return(nil)

	err := usecase.Delete(context.Background(), permissionRecordID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPermissionRecordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PermissionRecordRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewPermissionRecordUsecase(mockRepo, timeout)

	fetchedPermissionRecords := []domain.PermissionRecord{
		{
			ID:         primitive.NewObjectID(),
			Name:       "Manage Products",
			SystemName: "manage_products",
			Category:   "Catalog",
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Manage Orders",
			SystemName: "manage_orders",
			Category:   "Sales",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPermissionRecords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPermissionRecords, result)
	mockRepo.AssertExpectations(t)
}
