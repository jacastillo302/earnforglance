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

func TestStockQuantityChangeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.StockQuantityChangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewStockQuantityChangeUsecase(mockRepo, timeout)

	stockQuantityChangeID := primitive.NewObjectID().Hex()

	updatedStockQuantityChange := domain.StockQuantityChange{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		QuantityAdjustment: -5,
		StockQuantity:      95,
		Message:            "Stock reduced due to sale.",
		CreatedOnUtc:       time.Now(),
		ProductID:          primitive.NewObjectID(),
		CombinationID:      nil,
		WarehouseID:        primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, stockQuantityChangeID).Return(updatedStockQuantityChange, nil)

	result, err := usecase.FetchByID(context.Background(), stockQuantityChangeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedStockQuantityChange, result)
	mockRepo.AssertExpectations(t)
}

func TestStockQuantityChangeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.StockQuantityChangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewStockQuantityChangeUsecase(mockRepo, timeout)

	newStockQuantityChange := &domain.StockQuantityChange{
		QuantityAdjustment: 10,
		StockQuantity:      100,
		Message:            "Initial stock added.",
		CreatedOnUtc:       time.Now(),
		ProductID:          primitive.NewObjectID(),
		CombinationID:      nil,
		WarehouseID:        primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newStockQuantityChange).Return(nil)

	err := usecase.Create(context.Background(), newStockQuantityChange)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStockQuantityChangeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.StockQuantityChangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewStockQuantityChangeUsecase(mockRepo, timeout)

	updatedStockQuantityChange := &domain.StockQuantityChange{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		QuantityAdjustment: -5,
		StockQuantity:      95,
		Message:            "Stock reduced due to sale.",
		CreatedOnUtc:       time.Now(),
		ProductID:          primitive.NewObjectID(),
		CombinationID:      nil,
		WarehouseID:        primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedStockQuantityChange).Return(nil)

	err := usecase.Update(context.Background(), updatedStockQuantityChange)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStockQuantityChangeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.StockQuantityChangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewStockQuantityChangeUsecase(mockRepo, timeout)

	stockQuantityChangeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, stockQuantityChangeID).Return(nil)

	err := usecase.Delete(context.Background(), stockQuantityChangeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStockQuantityChangeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.StockQuantityChangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewStockQuantityChangeUsecase(mockRepo, timeout)

	fetchedStockQuantityChanges := []domain.StockQuantityChange{
		{
			ID:                 primitive.NewObjectID(),
			QuantityAdjustment: 10,
			StockQuantity:      100,
			Message:            "Initial stock added.",
			CreatedOnUtc:       time.Now().AddDate(0, 0, -10), // 10 days ago
			ProductID:          primitive.NewObjectID(),
			CombinationID:      nil,
			WarehouseID:        primitive.NewObjectID(),
		},
		{
			ID:                 primitive.NewObjectID(),
			QuantityAdjustment: -5,
			StockQuantity:      95,
			Message:            "Stock reduced due to sale.",
			CreatedOnUtc:       time.Now().AddDate(0, 0, -5), // 5 days ago
			ProductID:          primitive.NewObjectID(),
			CombinationID:      nil,
			WarehouseID:        primitive.NewObjectID(),
		},
	}
	fetchedStockQuantityChanges[1].StockQuantity = 100
	fetchedStockQuantityChanges[1].Message = "Stock reduced due to sale. update"

	mockRepo.On("Fetch", mock.Anything).Return(fetchedStockQuantityChanges, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedStockQuantityChanges, result)
	mockRepo.AssertExpectations(t)
}
