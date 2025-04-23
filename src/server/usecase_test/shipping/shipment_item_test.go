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

func TestShipmentItemUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShipmentItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentItemUsecase(mockRepo, timeout)

	shipmentItemID := bson.NewObjectID().Hex()

	updatedShipmentItem := domain.ShipmentItem{
		ID:          bson.NewObjectID(), // Existing ID of the record to update
		ShipmentID:  bson.NewObjectID(),
		OrderItemID: bson.NewObjectID(),
		Quantity:    15,
		WarehouseID: bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, shipmentItemID).Return(updatedShipmentItem, nil)

	result, err := usecase.FetchByID(context.Background(), shipmentItemID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShipmentItem, result)
	mockRepo.AssertExpectations(t)
}

func TestShipmentItemUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShipmentItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentItemUsecase(mockRepo, timeout)

	newShipmentItem := &domain.ShipmentItem{
		ShipmentID:  bson.NewObjectID(),
		OrderItemID: bson.NewObjectID(),
		Quantity:    10,
		WarehouseID: bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newShipmentItem).Return(nil)

	err := usecase.Create(context.Background(), newShipmentItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipmentItemUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShipmentItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentItemUsecase(mockRepo, timeout)

	updatedShipmentItem := &domain.ShipmentItem{
		ID:          bson.NewObjectID(), // Existing ID of the record to update
		ShipmentID:  bson.NewObjectID(),
		OrderItemID: bson.NewObjectID(),
		Quantity:    15,
		WarehouseID: bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedShipmentItem).Return(nil)

	err := usecase.Update(context.Background(), updatedShipmentItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipmentItemUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShipmentItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentItemUsecase(mockRepo, timeout)

	shipmentItemID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shipmentItemID).Return(nil)

	err := usecase.Delete(context.Background(), shipmentItemID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipmentItemUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShipmentItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentItemUsecase(mockRepo, timeout)

	fetchedShipmentItems := []domain.ShipmentItem{
		{
			ID:          bson.NewObjectID(),
			ShipmentID:  bson.NewObjectID(),
			OrderItemID: bson.NewObjectID(),
			Quantity:    10,
			WarehouseID: bson.NewObjectID(),
		},
		{
			ID:          bson.NewObjectID(),
			ShipmentID:  bson.NewObjectID(),
			OrderItemID: bson.NewObjectID(),
			Quantity:    20,
			WarehouseID: bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShipmentItems, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShipmentItems, result)
	mockRepo.AssertExpectations(t)
}
