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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShipmentUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShipmentRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentUsecase(mockRepo, timeout)

	shipmentID := primitive.NewObjectID().Hex()

	updatedShipment := domain.Shipment{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		OrderID:               primitive.NewObjectID(),
		TrackingNumber:        "UPDATEDTRACK67890",
		TotalWeight:           new(float64),
		ShippedDateUtc:        new(time.Time),
		DeliveryDateUtc:       new(time.Time),
		ReadyForPickupDateUtc: new(time.Time),
		AdminComment:          "Shipment details updated.",
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, shipmentID).Return(updatedShipment, nil)

	result, err := usecase.FetchByID(context.Background(), shipmentID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShipment, result)
	mockRepo.AssertExpectations(t)
}

func TestShipmentUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShipmentRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentUsecase(mockRepo, timeout)

	newShipment := &domain.Shipment{
		OrderID:               primitive.NewObjectID(),
		TrackingNumber:        "TRACK12345",
		TotalWeight:           new(float64),
		ShippedDateUtc:        new(time.Time),
		DeliveryDateUtc:       new(time.Time),
		ReadyForPickupDateUtc: new(time.Time),
		AdminComment:          "Shipment created successfully.",
		CreatedOnUtc:          time.Now(),
	}
	*newShipment.TotalWeight = 15.5
	*newShipment.ShippedDateUtc = time.Now().AddDate(0, 0, -1)       // Shipped 1 day ago
	*newShipment.DeliveryDateUtc = time.Now().AddDate(0, 0, 2)       // Delivery in 2 days
	*newShipment.ReadyForPickupDateUtc = time.Now().AddDate(0, 0, 1) // Ready for pickup in 1 day

	mockRepo.On("Create", mock.Anything, newShipment).Return(nil)

	err := usecase.Create(context.Background(), newShipment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipmentUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShipmentRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentUsecase(mockRepo, timeout)

	updatedShipment := &domain.Shipment{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		OrderID:               primitive.NewObjectID(),
		TrackingNumber:        "UPDATEDTRACK67890",
		TotalWeight:           new(float64),
		ShippedDateUtc:        new(time.Time),
		DeliveryDateUtc:       new(time.Time),
		ReadyForPickupDateUtc: new(time.Time),
		AdminComment:          "Shipment details updated.",
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}
	*updatedShipment.TotalWeight = 20.0
	*updatedShipment.ShippedDateUtc = time.Now().AddDate(0, 0, -2)       // Shipped 2 days ago
	*updatedShipment.DeliveryDateUtc = time.Now().AddDate(0, 0, 3)       // Delivery in 3 days
	*updatedShipment.ReadyForPickupDateUtc = time.Now().AddDate(0, 0, 2) // Ready for pickup in 2 days

	mockRepo.On("Update", mock.Anything, updatedShipment).Return(nil)

	err := usecase.Update(context.Background(), updatedShipment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipmentUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShipmentRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentUsecase(mockRepo, timeout)

	shipmentID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shipmentID).Return(nil)

	err := usecase.Delete(context.Background(), shipmentID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipmentUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShipmentRepository)
	timeout := time.Duration(10)
	usecase := test.NewShipmentUsecase(mockRepo, timeout)

	fetchedShipments := []domain.Shipment{
		{
			ID:                    primitive.NewObjectID(),
			OrderID:               primitive.NewObjectID(),
			TrackingNumber:        "TRACK12345",
			TotalWeight:           new(float64),
			ShippedDateUtc:        new(time.Time),
			DeliveryDateUtc:       new(time.Time),
			ReadyForPickupDateUtc: new(time.Time),
			AdminComment:          "Shipment created successfully.",
			CreatedOnUtc:          time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:                    primitive.NewObjectID(),
			OrderID:               primitive.NewObjectID(),
			TrackingNumber:        "TRACK67890",
			TotalWeight:           new(float64),
			ShippedDateUtc:        new(time.Time),
			DeliveryDateUtc:       new(time.Time),
			ReadyForPickupDateUtc: new(time.Time),
			AdminComment:          "Shipment updated successfully.",
			CreatedOnUtc:          time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}
	*fetchedShipments[0].TotalWeight = 15.5
	*fetchedShipments[0].ShippedDateUtc = time.Now().AddDate(0, 0, -9)        // Shipped 9 days ago
	*fetchedShipments[0].DeliveryDateUtc = time.Now().AddDate(0, 0, -7)       // Delivered 7 days ago
	*fetchedShipments[0].ReadyForPickupDateUtc = time.Now().AddDate(0, 0, -8) // Ready for pickup 8 days ago

	*fetchedShipments[1].TotalWeight = 20.0
	*fetchedShipments[1].ShippedDateUtc = time.Now().AddDate(0, 0, -4)        // Shipped 4 days ago
	*fetchedShipments[1].DeliveryDateUtc = time.Now().AddDate(0, 0, -2)       // Delivered 2 days ago
	*fetchedShipments[1].ReadyForPickupDateUtc = time.Now().AddDate(0, 0, -3) // Ready for pickup 3 days ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShipments, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShipments, result)
	mockRepo.AssertExpectations(t)
}
