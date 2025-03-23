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

type MockSingleResultShippingOption struct {
	mock.Mock
}

func (m *MockSingleResultShippingOption) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShippingOption); ok {
		*v.(*domain.ShippingOption) = *result
	}
	return args.Error(1)
}

var mockItemShippingOption = &domain.ShippingOption{
	ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
	ShippingRateComputationMethodSystemName: "ExpressRate",
	Rate:                                    20.00,
	Name:                                    "Express Shipping",
	Description:                             "Delivery within 1-2 business days.",
	TransitDays:                             new(int),
	IsPickupInStore:                         true,
	DisplayOrder:                            new(int),
}

func TestShippingOptionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShippingOption

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingOption{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShippingOption, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingOptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingOption.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingOption{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingOptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingOption.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShippingOptionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingOption

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShippingOption).Return(nil, nil).Once()

	repo := repository.NewShippingOptionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShippingOption)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShippingOptionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingOption

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShippingOption.ID}
	update := bson.M{"$set": mockItemShippingOption}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShippingOptionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShippingOption)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
