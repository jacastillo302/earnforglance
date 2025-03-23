package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductWarehouseInventoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductWarehouseInventoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductWarehouseInventoryUsecase(mockRepo, timeout)

	productWarehouseInventoryID := primitive.NewObjectID().Hex()

	updatedProductWarehouseInventory := domain.ProductWarehouseInventory{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:        primitive.NewObjectID(),
		WarehouseID:      primitive.NewObjectID(),
		StockQuantity:    150,
		ReservedQuantity: 20,
	}

	mockRepo.On("FetchByID", mock.Anything, productWarehouseInventoryID).Return(updatedProductWarehouseInventory, nil)

	result, err := usecase.FetchByID(context.Background(), productWarehouseInventoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductWarehouseInventory, result)
	mockRepo.AssertExpectations(t)
}

func TestProductWarehouseInventoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductWarehouseInventoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductWarehouseInventoryUsecase(mockRepo, timeout)

	newProductWarehouseInventory := &domain.ProductWarehouseInventory{
		ProductID:        primitive.NewObjectID(),
		WarehouseID:      primitive.NewObjectID(),
		StockQuantity:    100,
		ReservedQuantity: 10,
	}

	mockRepo.On("Create", mock.Anything, newProductWarehouseInventory).Return(nil)

	err := usecase.Create(context.Background(), newProductWarehouseInventory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductWarehouseInventoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductWarehouseInventoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductWarehouseInventoryUsecase(mockRepo, timeout)

	updatedProductWarehouseInventory := &domain.ProductWarehouseInventory{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:        primitive.NewObjectID(),
		WarehouseID:      primitive.NewObjectID(),
		StockQuantity:    150,
		ReservedQuantity: 20,
	}

	mockRepo.On("Update", mock.Anything, updatedProductWarehouseInventory).Return(nil)

	err := usecase.Update(context.Background(), updatedProductWarehouseInventory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductWarehouseInventoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductWarehouseInventoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductWarehouseInventoryUsecase(mockRepo, timeout)

	productWarehouseInventoryID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productWarehouseInventoryID).Return(nil)

	err := usecase.Delete(context.Background(), productWarehouseInventoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductWarehouseInventoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductWarehouseInventoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductWarehouseInventoryUsecase(mockRepo, timeout)

	fetchedProductWarehouseInventories := []domain.ProductWarehouseInventory{
		{
			ID:               primitive.NewObjectID(),
			ProductID:        primitive.NewObjectID(),
			WarehouseID:      primitive.NewObjectID(),
			StockQuantity:    100,
			ReservedQuantity: 10,
		},
		{
			ID:               primitive.NewObjectID(),
			ProductID:        primitive.NewObjectID(),
			WarehouseID:      primitive.NewObjectID(),
			StockQuantity:    200,
			ReservedQuantity: 30,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductWarehouseInventories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductWarehouseInventories, result)
	mockRepo.AssertExpectations(t)
}
