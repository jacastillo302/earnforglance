package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProductEditorSettings struct {
	mock.Mock
}

func (m *MockSingleResultProductEditorSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductEditorSettings); ok {
		*v.(*domain.ProductEditorSettings) = *result
	}
	return args.Error(1)
}

var mockItemProductEditorSettings = &domain.ProductEditorSettings{
	ID:                              primitive.NewObjectID(), // Existing ID of the record to update
	ProductType:                     false,
	VisibleIndividually:             false,
	ProductTemplate:                 true,
	AdminComment:                    true,
	Vendor:                          false,
	Stores:                          false,
	ACL:                             true,
	ShowOnHomepage:                  false,
	AllowCustomerReviews:            false,
	ProductTags:                     false,
	ManufacturerPartNumber:          true,
	GTIN:                            false,
	ProductCost:                     false,
	TierPrices:                      true,
	Discounts:                       false,
	DisableBuyButton:                true,
	DisableWishlistButton:           true,
	AvailableForPreOrder:            false,
	CallForPrice:                    true,
	OldPrice:                        false,
	CustomerEntersPrice:             true,
	PAngV:                           true,
	RequireOtherProductsAddedToCart: true,
	IsGiftCard:                      true,
	DownloadableProduct:             false,
	RecurringProduct:                true,
	IsRental:                        true,
	FreeShipping:                    false,
	ShipSeparately:                  true,
	AdditionalShippingCharge:        false,
	DeliveryDate:                    false,
	ProductAvailabilityRange:        true,
	UseMultipleWarehouses:           false,
	Warehouse:                       false,
	DisplayStockAvailability:        false,
	MinimumStockQuantity:            false,
	LowStockActivity:                true,
	NotifyAdminForQuantityBelow:     false,
	Backorders:                      false,
	AllowBackInStockSubscriptions:   false,
	MinimumCartQuantity:             false,
	MaximumCartQuantity:             false,
	AllowedQuantities:               false,
	AllowAddingOnlyExistingAttributeCombinations: true,
	NotReturnable:           true,
	Weight:                  false,
	Dimensions:              false,
	AvailableStartDate:      false,
	AvailableEndDate:        false,
	MarkAsNew:               false,
	Published:               false,
	RelatedProducts:         false,
	CrossSellsProducts:      false,
	Seo:                     false,
	PurchasedWithOrders:     true,
	ProductAttributes:       false,
	SpecificationAttributes: false,
	Manufacturers:           false,
	StockQuantityChange:     false,
	AgeVerification:         true,
}

func TestProductEditorSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductEditorSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductEditorSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductEditorSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductEditorSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductEditorSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductEditorSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductEditorSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductEditorSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductEditorSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductEditorSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductEditorSettings).Return(nil, nil).Once()

	repo := repository.NewProductEditorSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductEditorSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductEditorSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductEditorSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductEditorSettings.ID}
	update := bson.M{"$set": mockItemProductEditorSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductEditorSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductEditorSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
