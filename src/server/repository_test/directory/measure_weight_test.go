package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/directory"
	repository "earnforglance/server/repository/directory"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultMeasureWeight struct {
	mock.Mock
}

func (m *MockSingleResultMeasureWeight) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MeasureWeight); ok {
		*v.(*domain.MeasureWeight) = *result
	}
	return args.Error(1)
}

var mockItemMeasureWeight = &domain.MeasureWeight{
	ID:            primitive.NewObjectID(), // Existing ID of the record to update
	Name:          "Pound",
	SystemKeyword: "lb",
	Ratio:         2.20462,
	DisplayOrder:  2,
}

func TestMeasureWeightRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMeasureWeight

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMeasureWeight{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMeasureWeight, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMeasureWeightRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMeasureWeight.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMeasureWeight{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMeasureWeightRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMeasureWeight.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMeasureWeightRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMeasureWeight

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMeasureWeight).Return(nil, nil).Once()

	repo := repository.NewMeasureWeightRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMeasureWeight)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMeasureWeightRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMeasureWeight

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMeasureWeight.ID}
	update := bson.M{"$set": mockItemMeasureWeight}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMeasureWeightRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMeasureWeight)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
