package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/shipping"
	repository "earnforglance/server/repository/shipping"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultShippingMethodCountryMapping struct {
	mock.Mock
}

func (m *MockSingleResultShippingMethodCountryMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShippingMethodCountryMapping); ok {
		*v.(*domain.ShippingMethodCountryMapping) = *result
	}
	return args.Error(1)
}

var mockItemShippingMethodCountryMapping = &domain.ShippingMethodCountryMapping{
	ID:               bson.NewObjectID(), // Existing ID of the record to update
	ShippingMethodID: bson.NewObjectID(),
	CountryID:        bson.NewObjectID(),
}

func TestShippingMethodCountryMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShippingMethodCountryMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingMethodCountryMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShippingMethodCountryMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingMethodCountryMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingMethodCountryMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShippingMethodCountryMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShippingMethodCountryMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShippingMethodCountryMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShippingMethodCountryMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingMethodCountryMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShippingMethodCountryMapping).Return(nil, nil).Once()

	repo := repository.NewShippingMethodCountryMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShippingMethodCountryMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShippingMethodCountryMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShippingMethodCountryMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShippingMethodCountryMapping.ID}
	update := bson.M{"$set": mockItemShippingMethodCountryMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShippingMethodCountryMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShippingMethodCountryMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
