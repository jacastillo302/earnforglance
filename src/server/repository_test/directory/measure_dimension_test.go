package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/directory"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultMeasureDimension struct {
	mock.Mock
}

func (m *MockSingleResultMeasureDimension) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MeasureDimension); ok {
		*v.(*domain.MeasureDimension) = *result
	}
	return args.Error(1)
}

var mockItemMeasureDimension = &domain.MeasureDimension{
	ID:            primitive.NewObjectID(), // Existing ID of the record to update
	Name:          "Inch",
	SystemKeyword: "in",
	Ratio:         2.54,
	DisplayOrder:  2,
}

func TestMeasureDimensionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMeasureDimension

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMeasureDimension{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMeasureDimension, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMeasureDimensionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMeasureDimension.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMeasureDimension{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMeasureDimensionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMeasureDimension.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMeasureDimensionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMeasureDimension

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMeasureDimension).Return(nil, nil).Once()

	repo := repository.NewMeasureDimensionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMeasureDimension)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMeasureDimensionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMeasureDimension

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMeasureDimension.ID}
	update := bson.M{"$set": mockItemMeasureDimension}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMeasureDimensionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMeasureDimension)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
