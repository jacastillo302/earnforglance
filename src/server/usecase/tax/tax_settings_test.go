package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/tax"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaxSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.TaxSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewTaxSettingsUsecase(mockRepo, timeout)

	taxSettingsID := primitive.NewObjectID().Hex()

	updatedTaxSettings := domain.TaxSettings{
		ID:                                    primitive.NewObjectID(), // Existing ID of the record to update
		TaxBasedOn:                            1,
		TaxBasedOnPickupPointAddress:          false,
		TaxDisplayType:                        2,
		ActiveTaxProviderSystemName:           "UpdatedTaxProvider",
		DefaultTaxAddressID:                   primitive.NewObjectID(),
		DisplayTaxSuffix:                      false,
		DisplayTaxRates:                       false,
		PricesIncludeTax:                      true,
		AutomaticallyDetectCountry:            false,
		AllowCustomersToSelectTaxDisplayType:  true,
		HideZeroTax:                           false,
		HideTaxInOrderSummary:                 true,
		ForceTaxExclusionFromOrderSubtotal:    false,
		DefaultTaxCategoryID:                  primitive.NewObjectID(),
		ShippingIsTaxable:                     false,
		ShippingPriceIncludesTax:              true,
		ShippingTaxClassID:                    primitive.NewObjectID(),
		PaymentMethodAdditionalFeeIsTaxable:   false,
		PaymentMethodAdditionalFeeIncludesTax: true,
		PaymentMethodAdditionalFeeTaxClassID:  primitive.NewObjectID(),
		EuVatEnabled:                          false,
		EuVatRequired:                         false,
		EuVatEnabledForGuests:                 true,
		EuVatShopCountryID:                    primitive.NewObjectID(),
		EuVatAllowVatExemption:                false,
		EuVatUseWebService:                    true,
		EuVatAssumeValid:                      true,
		EuVatEmailAdminWhenNewVatSubmitted:    false,
		LogErrors:                             false,
	}

	mockRepo.On("FetchByID", mock.Anything, taxSettingsID).Return(updatedTaxSettings, nil)

	result, err := usecase.FetchByID(context.Background(), taxSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedTaxSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestTaxSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.TaxSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewTaxSettingsUsecase(mockRepo, timeout)

	newTaxSettings := &domain.TaxSettings{
		TaxBasedOn:                            2,
		TaxBasedOnPickupPointAddress:          true,
		TaxDisplayType:                        1,
		ActiveTaxProviderSystemName:           "DefaultTaxProvider",
		DefaultTaxAddressID:                   primitive.NewObjectID(),
		DisplayTaxSuffix:                      true,
		DisplayTaxRates:                       true,
		PricesIncludeTax:                      false,
		AutomaticallyDetectCountry:            true,
		AllowCustomersToSelectTaxDisplayType:  false,
		HideZeroTax:                           true,
		HideTaxInOrderSummary:                 false,
		ForceTaxExclusionFromOrderSubtotal:    true,
		DefaultTaxCategoryID:                  primitive.NewObjectID(),
		ShippingIsTaxable:                     true,
		ShippingPriceIncludesTax:              false,
		ShippingTaxClassID:                    primitive.NewObjectID(),
		PaymentMethodAdditionalFeeIsTaxable:   true,
		PaymentMethodAdditionalFeeIncludesTax: false,
		PaymentMethodAdditionalFeeTaxClassID:  primitive.NewObjectID(),
		EuVatEnabled:                          true,
		EuVatRequired:                         true,
		EuVatEnabledForGuests:                 false,
		EuVatShopCountryID:                    primitive.NewObjectID(),
		EuVatAllowVatExemption:                true,
		EuVatUseWebService:                    false,
		EuVatAssumeValid:                      false,
		EuVatEmailAdminWhenNewVatSubmitted:    true,
		LogErrors:                             true,
	}

	mockRepo.On("Create", mock.Anything, newTaxSettings).Return(nil)

	err := usecase.Create(context.Background(), newTaxSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaxSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.TaxSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewTaxSettingsUsecase(mockRepo, timeout)

	updatedTaxSettings := &domain.TaxSettings{
		ID:                                    primitive.NewObjectID(), // Existing ID of the record to update
		TaxBasedOn:                            1,
		TaxBasedOnPickupPointAddress:          false,
		TaxDisplayType:                        2,
		ActiveTaxProviderSystemName:           "UpdatedTaxProvider",
		DefaultTaxAddressID:                   primitive.NewObjectID(),
		DisplayTaxSuffix:                      false,
		DisplayTaxRates:                       false,
		PricesIncludeTax:                      true,
		AutomaticallyDetectCountry:            false,
		AllowCustomersToSelectTaxDisplayType:  true,
		HideZeroTax:                           false,
		HideTaxInOrderSummary:                 true,
		ForceTaxExclusionFromOrderSubtotal:    false,
		DefaultTaxCategoryID:                  primitive.NewObjectID(),
		ShippingIsTaxable:                     false,
		ShippingPriceIncludesTax:              true,
		ShippingTaxClassID:                    primitive.NewObjectID(),
		PaymentMethodAdditionalFeeIsTaxable:   false,
		PaymentMethodAdditionalFeeIncludesTax: true,
		PaymentMethodAdditionalFeeTaxClassID:  primitive.NewObjectID(),
		EuVatEnabled:                          false,
		EuVatRequired:                         false,
		EuVatEnabledForGuests:                 true,
		EuVatShopCountryID:                    primitive.NewObjectID(),
		EuVatAllowVatExemption:                false,
		EuVatUseWebService:                    true,
		EuVatAssumeValid:                      true,
		EuVatEmailAdminWhenNewVatSubmitted:    false,
		LogErrors:                             false,
	}

	mockRepo.On("Update", mock.Anything, updatedTaxSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedTaxSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaxSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.TaxSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewTaxSettingsUsecase(mockRepo, timeout)

	taxSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, taxSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), taxSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaxSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.TaxSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewTaxSettingsUsecase(mockRepo, timeout)

	fetchedTaxSettings := []domain.TaxSettings{
		{
			ID:                                    primitive.NewObjectID(),
			TaxBasedOn:                            2,
			TaxBasedOnPickupPointAddress:          true,
			TaxDisplayType:                        1,
			ActiveTaxProviderSystemName:           "DefaultTaxProvider",
			DefaultTaxAddressID:                   primitive.NewObjectID(),
			DisplayTaxSuffix:                      true,
			DisplayTaxRates:                       true,
			PricesIncludeTax:                      false,
			AutomaticallyDetectCountry:            true,
			AllowCustomersToSelectTaxDisplayType:  false,
			HideZeroTax:                           true,
			HideTaxInOrderSummary:                 false,
			ForceTaxExclusionFromOrderSubtotal:    true,
			DefaultTaxCategoryID:                  primitive.NewObjectID(),
			ShippingIsTaxable:                     true,
			ShippingPriceIncludesTax:              false,
			ShippingTaxClassID:                    primitive.NewObjectID(),
			PaymentMethodAdditionalFeeIsTaxable:   true,
			PaymentMethodAdditionalFeeIncludesTax: false,
			PaymentMethodAdditionalFeeTaxClassID:  primitive.NewObjectID(),
			EuVatEnabled:                          true,
			EuVatRequired:                         true,
			EuVatEnabledForGuests:                 false,
			EuVatShopCountryID:                    primitive.NewObjectID(),
			EuVatAllowVatExemption:                true,
			EuVatUseWebService:                    false,
			EuVatAssumeValid:                      false,
			EuVatEmailAdminWhenNewVatSubmitted:    true,
			LogErrors:                             true,
		},
		{
			ID:                                    primitive.NewObjectID(),
			TaxBasedOn:                            1,
			TaxBasedOnPickupPointAddress:          false,
			TaxDisplayType:                        2,
			ActiveTaxProviderSystemName:           "UpdatedTaxProvider",
			DefaultTaxAddressID:                   primitive.NewObjectID(),
			DisplayTaxSuffix:                      false,
			DisplayTaxRates:                       false,
			PricesIncludeTax:                      true,
			AutomaticallyDetectCountry:            false,
			AllowCustomersToSelectTaxDisplayType:  true,
			HideZeroTax:                           false,
			HideTaxInOrderSummary:                 true,
			ForceTaxExclusionFromOrderSubtotal:    false,
			DefaultTaxCategoryID:                  primitive.NewObjectID(),
			ShippingIsTaxable:                     false,
			ShippingPriceIncludesTax:              true,
			ShippingTaxClassID:                    primitive.NewObjectID(),
			PaymentMethodAdditionalFeeIsTaxable:   false,
			PaymentMethodAdditionalFeeIncludesTax: true,
			PaymentMethodAdditionalFeeTaxClassID:  primitive.NewObjectID(),
			EuVatEnabled:                          false,
			EuVatRequired:                         false,
			EuVatEnabledForGuests:                 true,
			EuVatShopCountryID:                    primitive.NewObjectID(),
			EuVatAllowVatExemption:                false,
			EuVatUseWebService:                    true,
			EuVatAssumeValid:                      true,
			EuVatEmailAdminWhenNewVatSubmitted:    false,
			LogErrors:                             false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedTaxSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedTaxSettings, result)
	mockRepo.AssertExpectations(t)
}
