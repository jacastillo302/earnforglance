package repository_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProduct struct {
	mock.Mock
}

func (m *MockSingleResultProduct) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Product); ok {
		*v.(*domain.Product) = *result
	}
	return args.Error(1)
}

var mockItemProduct = &domain.Product{
	ID:                               primitive.NewObjectID(), // Existing ID of the record to update
	ProductTypeID:                    2,
	ParentGroupedID:                  primitive.NewObjectID(),
	VisibleIndividually:              false,
	Name:                             "Updated Product",
	ShortDescription:                 "Updated short description.",
	FullDescription:                  "Updated full description of the product.",
	AdminComment:                     "Updated admin comment.",
	ProductTemplateID:                primitive.NewObjectID(),
	VendorID:                         primitive.NewObjectID(),
	ShowOnHomepage:                   false,
	MetaKeywords:                     "updated, product, ecommerce",
	MetaDescription:                  "Updated meta description for the product.",
	MetaTitle:                        "Updated Product Title",
	AllowCustomerReviews:             false,
	ApprovedRatingSum:                30,
	NotApprovedRatingSum:             10,
	ApprovedTotalReviews:             5,
	NotApprovedTotalReviews:          3,
	SubjectToAcl:                     true,
	LimitedToStores:                  true,
	Sku:                              "SKU54321",
	ManufacturerPartNumber:           "MPN09876",
	Gtin:                             "0987654321098",
	IsGiftCard:                       true,
	GiftCardTypeID:                   1,
	OverriddenGiftCardAmount:         new(float64),
	RequireOtherProducts:             true,
	RequiredIDs:                      "123,456",
	AutomaticallyAddRequiredProducts: true,
	IsDownload:                       true,
	DownloadID:                       primitive.NewObjectID(),
	UnlimitedDownloads:               true,
	MaxNumberOfDownloads:             5,
	DownloadExpirationDays:           new(int),
	DownloadActivationTypeID:         1,
	HasSampleDownload:                true,
	SampleDownloadID:                 primitive.NewObjectID(),
	HasUserAgreement:                 true,
	UserAgreementText:                "Updated user agreement text.",
	IsRecurring:                      true,
	RecurringCycleLength:             30,
	RecurringProductCyclePeriodID:    2,
	RecurringTotalCycles:             12,
	IsRental:                         true,
	RentalPriceLength:                7,
	RentalPricePeriodID:              1,
	IsShipEnabled:                    false,
	IsFreeShipping:                   true,
	ShipSeparately:                   true,
	AdditionalShippingCharge:         0.0,
	DeliveryDateID:                   primitive.NewObjectID(),
	IsTaxExempt:                      true,
	TaxCategoryID:                    primitive.NewObjectID(),
	ManageInventoryMethodID:          2,
	ProductAvailabilityRangeID:       primitive.NewObjectID(),
	UseMultipleWarehouses:            true,
	WarehouseID:                      primitive.NewObjectID(),
	StockQuantity:                    50,
	DisplayStockAvailability:         false,
	DisplayStockQuantity:             false,
	MinStockQuantity:                 2,
	LowStockActivityID:               2,
	NotifyAdminForQuantityBelow:      5,
	BackorderModeID:                  1,
	AllowBackInStockSubscriptions:    false,
	OrderMinimumQuantity:             2,
	OrderMaximumQuantity:             5,
	AllowedQuantities:                "2,3,4,5",
	AllowAddingOnlyExistingAttributeCombinations: true,
	DisplayAttributeCombinationImagesOnly:        true,
	NotReturnable:                                true,
	DisableBuyButton:                             true,
	DisableWishlistButton:                        true,
	AvailableForPreOrder:                         true,
	PreOrderAvailabilityStartDateTimeUtc:         new(time.Time),
	CallForPrice:                                 true,
	Price:                                        79.99,
	OldPrice:                                     100.00,
	ProductCost:                                  40.00,
	CustomerEntersPrice:                          true,
	MinimumCustomerEnteredPrice:                  10.0,
	MaximumCustomerEnteredPrice:                  50.0,
	BasepriceEnabled:                             true,
	BasepriceAmount:                              1.0,
	BasepriceUnitID:                              primitive.NewObjectID(),
	BasepriceBaseAmount:                          10.0,
	BasepriceBaseUnitID:                          primitive.NewObjectID(),
	MarkAsNew:                                    false,
	MarkAsNewStartDateTimeUtc:                    new(time.Time),
	MarkAsNewEndDateTimeUtc:                      new(time.Time),
	Weight:                                       2.0,
	Length:                                       15.0,
	Width:                                        7.0,
	Height:                                       4.0,
	AvailableStartDateTimeUtc:                    new(time.Time),
	AvailableEndDateTimeUtc:                      new(time.Time),
	DisplayOrder:                                 2,
	Published:                                    false,
	Deleted:                                      true,
	CreatedOnUtc:                                 time.Now().AddDate(0, 0, -30), // Created 30 days ago
	UpdatedOnUtc:                                 time.Now(),
	AgeVerification:                              true,
	MinimumAgeToPurchase:                         18,
}

func TestProductRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProduct

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProduct{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProduct, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProduct.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProduct{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProduct.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProduct

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProduct).Return(nil, nil).Once()

	repo := repository.NewProductRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProduct)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProduct

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProduct.ID}
	update := bson.M{"$set": mockItemProduct}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProduct)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
