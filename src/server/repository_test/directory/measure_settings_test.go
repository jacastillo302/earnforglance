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

type MockSingleResultMeasureSettings struct {
	mock.Mock
}

func (m *MockSingleResultMeasureSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MeasureSettings); ok {
		*v.(*domain.MeasureSettings) = *result
	}
	return args.Error(1)
}

var mockItemMeasureSettings = &domain.MeasureSettings{
	ID:              primitive.NewObjectID(), // Existing ID of the record to update
	BaseDimensionID: 3,
	BaseWeightID:    4,
}

func TestMeasureSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMeasureSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMeasureSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMeasureSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMeasureSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMeasureSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMeasureSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMeasureSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMeasureSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMeasureSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMeasureSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMeasureSettings).Return(nil, nil).Once()

	repo := repository.NewMeasureSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMeasureSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMeasureSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMeasureSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMeasureSettings.ID}
	update := bson.M{"$set": mockItemMeasureSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMeasureSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMeasureSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
