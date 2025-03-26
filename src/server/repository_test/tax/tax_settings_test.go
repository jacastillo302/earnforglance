package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/tax"
	repository "earnforglance/server/repository/tax"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultTaxSettings struct {
	mock.Mock
}

func (m *MockSingleResultTaxSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.TaxSettings); ok {
		*v.(*domain.TaxSettings) = *result
	}
	return args.Error(1)
}

var mockItemTaxSettings = &domain.TaxSettings{
	ID:                                    primitive.NewObjectID(), // Existing ID of the record to update
	TaxBasedOnID:                          1,
	TaxBasedOnPickupPointAddress:          false,
	TaxDisplayTypeID:                      2,
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

func TestTaxSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTaxSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTaxSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemTaxSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTaxSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTaxSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTaxSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTaxSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTaxSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestTaxSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTaxSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemTaxSettings).Return(nil, nil).Once()

	repo := repository.NewTaxSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemTaxSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestTaxSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTaxSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemTaxSettings.ID}
	update := bson.M{"$set": mockItemTaxSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewTaxSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemTaxSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
