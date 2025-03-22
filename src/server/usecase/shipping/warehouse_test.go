package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/shipping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestWarehouseUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := NewWarehouseUsecase(mockRepo, timeout)

	warehouseID := primitive.NewObjectID().Hex()

	updatedWarehouse := domain.Warehouse{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Secondary Warehouse",
		AdminComment: "Backup storage facility.",
		AddressID:    primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, warehouseID).Return(updatedWarehouse, nil)

	result, err := usecase.FetchByID(context.Background(), warehouseID)

	assert.NoError(t, err)
	assert.Equal(t, updatedWarehouse, result)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := NewWarehouseUsecase(mockRepo, timeout)

	newWarehouse := &domain.Warehouse{
		Name:         "Main Warehouse",
		AdminComment: "Primary storage facility.",
		AddressID:    primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newWarehouse).Return(nil)

	err := usecase.Create(context.Background(), newWarehouse)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := NewWarehouseUsecase(mockRepo, timeout)

	updatedWarehouse := &domain.Warehouse{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Secondary Warehouse",
		AdminComment: "Backup storage facility.",
		AddressID:    primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedWarehouse).Return(nil)

	err := usecase.Update(context.Background(), updatedWarehouse)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := NewWarehouseUsecase(mockRepo, timeout)

	warehouseID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, warehouseID).Return(nil)

	err := usecase.Delete(context.Background(), warehouseID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := NewWarehouseUsecase(mockRepo, timeout)

	fetchedWarehouses := []domain.Warehouse{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Main Warehouse",
			AdminComment: "Primary storage facility.",
			AddressID:    primitive.NewObjectID(),
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Secondary Warehouse",
			AdminComment: "Backup storage facility.",
			AddressID:    primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedWarehouses, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedWarehouses, result)
	mockRepo.AssertExpectations(t)
}
