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

type MockSingleResultProductAvailabilityRange struct {
	mock.Mock
}

func (m *MockSingleResultProductAvailabilityRange) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAvailabilityRange); ok {
		*v.(*domain.ProductAvailabilityRange) = *result
	}
	return args.Error(1)
}

var mockItemProductAvailabilityRange = &domain.ProductAvailabilityRange{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	Name:         "Out of Stock",
	DisplayOrder: 2,
}

func TestProductAvailabilityRangeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAvailabilityRange

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAvailabilityRange{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductAvailabilityRange, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAvailabilityRangeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAvailabilityRange.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAvailabilityRange{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAvailabilityRangeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAvailabilityRange.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAvailabilityRangeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAvailabilityRange

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductAvailabilityRange).Return(nil, nil).Once()

	repo := repository.NewProductAvailabilityRangeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductAvailabilityRange)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAvailabilityRangeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAvailabilityRange

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductAvailabilityRange.ID}
	update := bson.M{"$set": mockItemProductAvailabilityRange}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAvailabilityRangeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductAvailabilityRange)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
