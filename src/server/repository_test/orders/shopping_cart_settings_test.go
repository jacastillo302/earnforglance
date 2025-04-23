package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultShoppingCartSettings struct {
	mock.Mock
}

func (m *MockSingleResultShoppingCartSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShoppingCartSettings); ok {
		*v.(*domain.ShoppingCartSettings) = *result
	}
	return args.Error(1)
}

var mockItemShoppingCartSettings = &domain.ShoppingCartSettings{
	ID:                                          bson.NewObjectID(), // Existing ID of the record to update
	DisplayCartAfterAddingProduct:               false,
	DisplayWishlistAfterAddingProduct:           true,
	MaximumShoppingCartItems:                    100,
	MaximumWishlistItems:                        50,
	AllowOutOfStockItemsToBeAddedToWishlist:     false,
	MoveItemsFromWishlistToCart:                 false,
	CartsSharedBetweenStores:                    true,
	ShowProductImagesOnShoppingCart:             false,
	ShowProductImagesOnWishList:                 false,
	ShowDiscountBox:                             false,
	ShowGiftCardBox:                             true,
	CrossSellsNumber:                            10,
	EmailWishlistEnabled:                        false,
	AllowAnonymousUsersToEmailWishlist:          true,
	MiniShoppingCartEnabled:                     false,
	ShowProductImagesInMiniShoppingCart:         false,
	MiniShoppingCartProductNumber:               5,
	RoundPricesDuringCalculation:                false,
	GroupTierPricesForDistinctShoppingCartItems: true,
	AllowCartItemEditing:                        false,
	RenderAssociatedAttributeValueQuantity:      true,
}

func TestShoppingCartSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShoppingCartSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShoppingCartSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShoppingCartSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShoppingCartSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShoppingCartSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShoppingCartSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShoppingCartSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShoppingCartSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShoppingCartSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShoppingCartSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShoppingCartSettings).Return(nil, nil).Once()

	repo := repository.NewShoppingCartSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShoppingCartSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShoppingCartSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShoppingCartSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShoppingCartSettings.ID}
	update := bson.M{"$set": mockItemShoppingCartSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShoppingCartSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShoppingCartSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
