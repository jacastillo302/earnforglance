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

func TestShippingSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShippingSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingSettingsUsecase(mockRepo, timeout)

	shippingID := primitive.NewObjectID().Hex()

	updatedShippingSettings := domain.ShippingSettings{
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
		ShippingSorting:                                  1,
	}

	mockRepo.On("FetchByID", mock.Anything, shippingID).Return(updatedShippingSettings, nil)

	result, err := usecase.FetchByID(context.Background(), shippingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShippingSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestShippingSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShippingSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingSettingsUsecase(mockRepo, timeout)

	newShippingSettings := &domain.ShippingSettings{
		ActiveShippingRateComputationMethodSystemNames:   []string{"FlatRate", "WeightBased"},
		ActivePickupPointProviderSystemNames:             []string{"FedEx", "UPS"},
		ShipToSameAddress:                                true,
		AllowPickupInStore:                               true,
		DisplayPickupPointsOnMap:                         true,
		IgnoreAdditionalShippingChargeForPickupInStore:   false,
		GoogleMapsApiKey:                                 "AIzaSyExampleKey",
		UseWarehouseLocation:                             true,
		NotifyCustomerAboutShippingFromMultipleLocations: true,
		FreeShippingOverXEnabled:                         true,
		FreeShippingOverXValue:                           50.00,
		FreeShippingOverXIncludingTax:                    false,
		EstimateShippingCartPageEnabled:                  true,
		EstimateShippingProductPageEnabled:               false,
		EstimateShippingCityNameEnabled:                  true,
		DisplayShipmentEventsToCustomers:                 true,
		DisplayShipmentEventsToStoreOwner:                true,
		HideShippingTotal:                                false,
		ShippingOriginAddressID:                          1,
		ReturnValidOptionsIfThereAreAny:                  true,
		BypassShippingMethodSelectionIfOnlyOne:           true,
		UseCubeRootMethod:                                false,
		ConsiderAssociatedProductsDimensions:             true,
		ShipSeparatelyOneItemEach:                        false,
		RequestDelay:                                     500,
		ShippingSorting:                                  2,
	}

	mockRepo.On("Create", mock.Anything, newShippingSettings).Return(nil)

	err := usecase.Create(context.Background(), newShippingSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShippingSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingSettingsUsecase(mockRepo, timeout)

	updatedShippingSettings := &domain.ShippingSettings{
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
		ShippingSorting:                                  1,
	}

	mockRepo.On("Update", mock.Anything, updatedShippingSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedShippingSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShippingSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingSettingsUsecase(mockRepo, timeout)

	shippingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shippingID).Return(nil)

	err := usecase.Delete(context.Background(), shippingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShippingSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingSettingsUsecase(mockRepo, timeout)

	fetchedShippingSettings := []domain.ShippingSettings{
		{
			ID: primitive.NewObjectID(),
			ActiveShippingRateComputationMethodSystemNames:   []string{"FlatRate", "WeightBased"},
			ActivePickupPointProviderSystemNames:             []string{"FedEx", "UPS"},
			ShipToSameAddress:                                true,
			AllowPickupInStore:                               true,
			DisplayPickupPointsOnMap:                         true,
			IgnoreAdditionalShippingChargeForPickupInStore:   false,
			GoogleMapsApiKey:                                 "AIzaSyExampleKey",
			UseWarehouseLocation:                             true,
			NotifyCustomerAboutShippingFromMultipleLocations: true,
			FreeShippingOverXEnabled:                         true,
			FreeShippingOverXValue:                           50.00,
			FreeShippingOverXIncludingTax:                    false,
			EstimateShippingCartPageEnabled:                  true,
			EstimateShippingProductPageEnabled:               false,
			EstimateShippingCityNameEnabled:                  true,
			DisplayShipmentEventsToCustomers:                 true,
			DisplayShipmentEventsToStoreOwner:                true,
			HideShippingTotal:                                false,
			ShippingOriginAddressID:                          1,
			ReturnValidOptionsIfThereAreAny:                  true,
			BypassShippingMethodSelectionIfOnlyOne:           true,
			UseCubeRootMethod:                                false,
			ConsiderAssociatedProductsDimensions:             true,
			ShipSeparatelyOneItemEach:                        false,
			RequestDelay:                                     500,
			ShippingSorting:                                  2,
		},
		{
			ID: primitive.NewObjectID(),
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
			ShippingSorting:                                  1,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShippingSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShippingSettings, result)
	mockRepo.AssertExpectations(t)
}
