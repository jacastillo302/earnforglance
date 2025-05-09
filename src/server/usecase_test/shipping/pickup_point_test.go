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

func TestPickupPointUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PickupPointRepository)
	timeout := time.Duration(10)
	usecase := test.NewPickupPointUsecase(mockRepo, timeout)

	pickupPointID := bson.NewObjectID().Hex()

	updatedPickupPoint := domain.PickupPoint{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		Name:               "Downtown Pickup",
		Description:        "Pickup point located downtown.",
		ProviderSystemName: "UPS",
		Address:            "456 Downtown Ave",
		City:               "Los Angeles",
		County:             "Los Angeles County",
		StateAbbreviation:  "CA",
		CountryCode:        "US",
		ZipPostalCode:      "90001",
		Latitude:           new(float64),
		Longitude:          new(float64),
		PickupFee:          7.50,
		OpeningHours:       "10:00 AM - 6:00 PM",
		DisplayOrder:       2,
		TransitDays:        new(int),
	}
	*updatedPickupPoint.Latitude = 34.0522
	*updatedPickupPoint.Longitude = -118.2437
	*updatedPickupPoint.TransitDays = 3
	mockRepo.On("FetchByID", mock.Anything, pickupPointID).Return(updatedPickupPoint, nil)

	result, err := usecase.FetchByID(context.Background(), pickupPointID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPickupPoint, result)
	mockRepo.AssertExpectations(t)
}

func TestPickupPointUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PickupPointRepository)
	timeout := time.Duration(10)
	usecase := test.NewPickupPointUsecase(mockRepo, timeout)

	newPickupPoint := &domain.PickupPoint{
		Name:               "Main Street Pickup",
		Description:        "Pickup point located on Main Street.",
		ProviderSystemName: "FedEx",
		Address:            "123 Main Street",
		City:               "New York",
		County:             "New York County",
		StateAbbreviation:  "NY",
		CountryCode:        "US",
		ZipPostalCode:      "10001",
		Latitude:           new(float64),
		Longitude:          new(float64),
		PickupFee:          5.00,
		OpeningHours:       "9:00 AM - 5:00 PM",
		DisplayOrder:       1,
		TransitDays:        new(int),
	}
	*newPickupPoint.Latitude = 40.7128
	*newPickupPoint.Longitude = -74.0060
	*newPickupPoint.TransitDays = 2

	mockRepo.On("Create", mock.Anything, newPickupPoint).Return(nil)

	err := usecase.Create(context.Background(), newPickupPoint)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPickupPointUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PickupPointRepository)
	timeout := time.Duration(10)
	usecase := test.NewPickupPointUsecase(mockRepo, timeout)

	updatedPickupPoint := &domain.PickupPoint{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		Name:               "Downtown Pickup",
		Description:        "Pickup point located downtown.",
		ProviderSystemName: "UPS",
		Address:            "456 Downtown Ave",
		City:               "Los Angeles",
		County:             "Los Angeles County",
		StateAbbreviation:  "CA",
		CountryCode:        "US",
		ZipPostalCode:      "90001",
		Latitude:           new(float64),
		Longitude:          new(float64),
		PickupFee:          7.50,
		OpeningHours:       "10:00 AM - 6:00 PM",
		DisplayOrder:       2,
		TransitDays:        new(int),
	}
	*updatedPickupPoint.Latitude = 34.0522
	*updatedPickupPoint.Longitude = -118.2437
	*updatedPickupPoint.TransitDays = 3

	mockRepo.On("Update", mock.Anything, updatedPickupPoint).Return(nil)

	err := usecase.Update(context.Background(), updatedPickupPoint)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPickupPointUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PickupPointRepository)
	timeout := time.Duration(10)
	usecase := test.NewPickupPointUsecase(mockRepo, timeout)

	pickupPointID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pickupPointID).Return(nil)

	err := usecase.Delete(context.Background(), pickupPointID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPickupPointUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PickupPointRepository)
	timeout := time.Duration(10)
	usecase := test.NewPickupPointUsecase(mockRepo, timeout)

	fetchedPickupPoints := []domain.PickupPoint{
		{
			ID:                 bson.NewObjectID(),
			Name:               "Main Street Pickup",
			Description:        "Pickup point located on Main Street.",
			ProviderSystemName: "FedEx",
			Address:            "123 Main Street",
			City:               "New York",
			County:             "New York County",
			StateAbbreviation:  "NY",
			CountryCode:        "US",
			ZipPostalCode:      "10001",
			Latitude:           new(float64),
			Longitude:          new(float64),
			PickupFee:          5.00,
			OpeningHours:       "9:00 AM - 5:00 PM",
			DisplayOrder:       1,
			TransitDays:        new(int),
		},
		{
			ID:                 bson.NewObjectID(),
			Name:               "Downtown Pickup",
			Description:        "Pickup point located downtown.",
			ProviderSystemName: "UPS",
			Address:            "456 Downtown Ave",
			City:               "Los Angeles",
			County:             "Los Angeles County",
			StateAbbreviation:  "CA",
			CountryCode:        "US",
			ZipPostalCode:      "90001",
			Latitude:           new(float64),
			Longitude:          new(float64),
			PickupFee:          7.50,
			OpeningHours:       "10:00 AM - 6:00 PM",
			DisplayOrder:       2,
			TransitDays:        new(int),
		},
	}
	*fetchedPickupPoints[0].Latitude = 40.7128
	*fetchedPickupPoints[0].Longitude = -74.0060
	*fetchedPickupPoints[0].TransitDays = 2
	*fetchedPickupPoints[1].Latitude = 34.0522
	*fetchedPickupPoints[1].Longitude = -118.2437
	*fetchedPickupPoints[1].TransitDays = 3

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPickupPoints, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPickupPoints, result)
	mockRepo.AssertExpectations(t)
}
