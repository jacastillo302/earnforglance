package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultShoppingCartItem struct {
	mock.Mock
}

func (m *MockSingleResultShoppingCartItem) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShoppingCartItem); ok {
		*v.(*domain.ShoppingCartItem) = *result
	}
	return args.Error(1)
}

var mockItemShoppingCartItem = &domain.ShoppingCartItem{
	ID:                   primitive.NewObjectID(), // Existing ID of the record to update
	StoreID:              primitive.NewObjectID(),
	ShoppingCartTypeID:   1,
	CustomerID:           primitive.NewObjectID(),
	ProductID:            primitive.NewObjectID(),
	AttributesXml:        "<Attributes><Color>Blue</Color><Size>L</Size></Attributes>",
	CustomerEnteredPrice: 59.99,
	Quantity:             3,
	RentalStartDateUtc:   new(time.Time),
	RentalEndDateUtc:     new(time.Time),
	CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
	UpdatedOnUtc:         time.Now(),
}

func TestShoppingCartItemRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShoppingCartItem

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShoppingCartItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShoppingCartItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShoppingCartItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShoppingCartItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShoppingCartItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShoppingCartItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShoppingCartItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShoppingCartItemRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShoppingCartItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShoppingCartItem).Return(nil, nil).Once()

	repo := repository.NewShoppingCartItemRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShoppingCartItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShoppingCartItemRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShoppingCartItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShoppingCartItem.ID}
	update := bson.M{"$set": mockItemShoppingCartItem}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShoppingCartItemRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShoppingCartItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
