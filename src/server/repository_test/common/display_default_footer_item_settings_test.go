package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultDisplayDefaultFooterItemSettings struct {
	mock.Mock
}

func (m *MockSingleResultDisplayDefaultFooterItemSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DisplayDefaultFooterItemSettings); ok {
		*v.(*domain.DisplayDefaultFooterItemSettings) = *result
	}
	return args.Error(1)
}

var mockItemDisplayDefaultFooterItemSettings = &domain.DisplayDefaultFooterItemSettings{
	ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
	DisplaySitemapFooterItem:                false,
	DisplayContactUsFooterItem:              false,
	DisplayProductSearchFooterItem:          true,
	DisplayNewsFooterItem:                   true,
	DisplayBlogFooterItem:                   false,
	DisplayForumsFooterItem:                 true,
	DisplayRecentlyViewedProductsFooterItem: false,
	DisplayCompareProductsFooterItem:        true,
	DisplayNewProductsFooterItem:            false,
	DisplayCustomerInfoFooterItem:           false,
	DisplayCustomerOrdersFooterItem:         true,
	DisplayCustomerAddressesFooterItem:      true,
	DisplayShoppingCartFooterItem:           false,
	DisplayWishlistFooterItem:               false,
	DisplayApplyVendorAccountFooterItem:     true,
}

func TestDisplayDefaultFooterItemSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDisplayDefaultFooterItemSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDisplayDefaultFooterItemSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDisplayDefaultFooterItemSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDisplayDefaultFooterItemSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDisplayDefaultFooterItemSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDisplayDefaultFooterItemSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDisplayDefaultFooterItemSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDisplayDefaultFooterItemSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDisplayDefaultFooterItemSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDisplayDefaultFooterItemSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDisplayDefaultFooterItemSettings).Return(nil, nil).Once()

	repo := repository.NewDisplayDefaultFooterItemSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDisplayDefaultFooterItemSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDisplayDefaultFooterItemSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDisplayDefaultFooterItemSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDisplayDefaultFooterItemSettings.ID}
	update := bson.M{"$set": mockItemDisplayDefaultFooterItemSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDisplayDefaultFooterItemSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDisplayDefaultFooterItemSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
