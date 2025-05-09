package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/cms"
	repository "earnforglance/server/repository/cms"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultWidgetSettings struct {
	mock.Mock
}

func (m *MockSingleResultWidgetSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.WidgetSettings); ok {
		*v.(*domain.WidgetSettings) = *result
	}
	return args.Error(1)
}

var mockItemWidgetSettings = &domain.WidgetSettings{
	ID:                      bson.NewObjectID(), // Existing ID of the record to update
	ActiveWidgetSystemNames: []string{"WidgetX", "WidgetY"},
}

func TestWidgetSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionWidgetSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultWidgetSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemWidgetSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewWidgetSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemWidgetSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultWidgetSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewWidgetSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemWidgetSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestWidgetSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionWidgetSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemWidgetSettings).Return(nil, nil).Once()

	repo := repository.NewWidgetSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemWidgetSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestWidgetSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionWidgetSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemWidgetSettings.ID}
	update := bson.M{"$set": mockItemWidgetSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewWidgetSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemWidgetSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
