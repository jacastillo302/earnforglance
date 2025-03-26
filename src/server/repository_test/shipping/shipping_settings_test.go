package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/shipping"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultShippingSettings struct {
	mock.Mock
}

func (m *MockSingleResultShippingSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShippingSettings); ok {
		*v.(*domain.ShippingSettings) = *result
	}
	return args.Error(1)
}

var mockItemShippingSettings = &domain.ShippingSettings{
	ID: primitive.NewObjectID(), // Existing ID of the record to update
	ActiveShippingRateComputationMethodSystemNames:   []string{"ExpressRate"},
	ActivePickupPointProviderSystemNames:             []string{"DHL"},
	ShipToSameAddress:                                false,
	AllowPickupInStore:                               false,
	DisplayPickupPointsOnMap:                         false,
	IgnoreAdditionalShippingChargeForPickupInStore:   true,
	GoogleMapsApiKey:                                 "UpdatedExampleKey",
	UseWarehouseLocation:                             false,
	NotifyCustomerAboutShippingFromMultipleLocations: false,
	FreeShippingOverXEnabled:                         false,
	FreeShippingOverXValue:                           100.00,
	FreeShippingOverXIncludingTax:                    true,
	EstimateShippingCartPageEnabled:                  false,
	EstimateShippingProductPageEnabled:               true,
	EstimateShippingCityNameEnabled:                  false,
	DisplayShipmentEventsToCustomers:                 false,
	DisplayShipmentEventsToStoreOwner:                false,
	HideShippingTotal:                                true,
	ShippingOriginAddressID:                          2,
	ReturnValidOptionsIfThereAreAny:                  false,
	BypassShippingMethodSelectionIfOnlyOne:           false,
	UseCubeRootMethod:                                true,
	ConsiderAssociatedProductsDimensions:             false,
	ShipSeparatelyOneItemEach:                        true,
	RequestDelay:                                     1000,
	ShippingSortingID:                                1,
}

func TestShippingSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShippingSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShippingSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShippingSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShippingSettings).Return(nil, nil).Once()

	repo := repository.NewShippingSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShippingSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShippingSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShippingSettings.ID}
	update := bson.M{"$set": mockItemShippingSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShippingSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShippingSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
