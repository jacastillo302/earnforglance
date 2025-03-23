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

type MockSingleResultShippingMethod struct {
	mock.Mock
}

func (m *MockSingleResultShippingMethod) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShippingMethod); ok {
		*v.(*domain.ShippingMethod) = *result
	}
	return args.Error(1)
}

var mockItemShippingMethod = &domain.ShippingMethod{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	Name:         "Express Shipping",
	Description:  "Delivery within 1-2 business days.",
	DisplayOrder: 2,
}

func TestShippingMethodRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShippingMethod

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingMethod{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShippingMethod, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingMethodRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingMethod.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingMethod{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingMethodRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingMethod.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShippingMethodRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingMethod

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShippingMethod).Return(nil, nil).Once()

	repo := repository.NewShippingMethodRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShippingMethod)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShippingMethodRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingMethod

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShippingMethod.ID}
	update := bson.M{"$set": mockItemShippingMethod}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShippingMethodRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShippingMethod)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
