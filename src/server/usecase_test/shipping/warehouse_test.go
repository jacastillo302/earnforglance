package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/shipping"
	test "earnforglance/server/usecase/shipping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestWarehouseUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := test.NewWarehouseUsecase(mockRepo, timeout)

	warehouseID := bson.NewObjectID().Hex()

	updatedWarehouse := domain.Warehouse{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Secondary Warehouse",
		AdminComment: "Backup storage facility.",
		AddressID:    bson.NewObjectID(),
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
	usecase := test.NewWarehouseUsecase(mockRepo, timeout)

	newWarehouse := &domain.Warehouse{
		Name:         "Main Warehouse",
		AdminComment: "Primary storage facility.",
		AddressID:    bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newWarehouse).Return(nil)

	err := usecase.Create(context.Background(), newWarehouse)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := test.NewWarehouseUsecase(mockRepo, timeout)

	updatedWarehouse := &domain.Warehouse{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Secondary Warehouse",
		AdminComment: "Backup storage facility.",
		AddressID:    bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedWarehouse).Return(nil)

	err := usecase.Update(context.Background(), updatedWarehouse)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := test.NewWarehouseUsecase(mockRepo, timeout)

	warehouseID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, warehouseID).Return(nil)

	err := usecase.Delete(context.Background(), warehouseID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWarehouseUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.WarehouseRepository)
	timeout := time.Duration(10)
	usecase := test.NewWarehouseUsecase(mockRepo, timeout)

	fetchedWarehouses := []domain.Warehouse{
		{
			ID:           bson.NewObjectID(),
			Name:         "Main Warehouse",
			AdminComment: "Primary storage facility.",
			AddressID:    bson.NewObjectID(),
		},
		{
			ID:           bson.NewObjectID(),
			Name:         "Secondary Warehouse",
			AdminComment: "Backup storage facility.",
			AddressID:    bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedWarehouses, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedWarehouses, result)
	mockRepo.AssertExpectations(t)
}
