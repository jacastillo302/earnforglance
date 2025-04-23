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

type MockSingleResultCheckoutAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultCheckoutAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CheckoutAttributeValue); ok {
		*v.(*domain.CheckoutAttributeValue) = *result
	}
	return args.Error(1)
}

var mockItemCheckoutAttributeValue = &domain.CheckoutAttributeValue{
	ID:               bson.NewObjectID(), // Existing ID of the record to update
	ColorSquaresRgb:  "#33FF57",
	PriceAdjustment:  15.75,
	WeightAdjustment: 0.50,
}

func TestCheckoutAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCheckoutAttributeValue

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCheckoutAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCheckoutAttributeValue, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCheckoutAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCheckoutAttributeValue.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCheckoutAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCheckoutAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCheckoutAttributeValue.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCheckoutAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCheckoutAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCheckoutAttributeValue).Return(nil, nil).Once()

	repo := repository.NewCheckoutAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCheckoutAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCheckoutAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCheckoutAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCheckoutAttributeValue.ID}
	update := bson.M{"$set": mockItemCheckoutAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCheckoutAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCheckoutAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
