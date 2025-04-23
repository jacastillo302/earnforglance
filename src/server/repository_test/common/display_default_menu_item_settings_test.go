package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultDisplayDefaultMenuItemSettings struct {
	mock.Mock
}

func (m *MockSingleResultDisplayDefaultMenuItemSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DisplayDefaultMenuItemSettings); ok {
		*v.(*domain.DisplayDefaultMenuItemSettings) = *result
	}
	return args.Error(1)
}

var mockItemDisplayDefaultMenuItemSettings = &domain.DisplayDefaultMenuItemSettings{
	ID:                           bson.NewObjectID(), // Existing ID of the record to update
	DisplayHomepageMenuItem:      false,
	DisplayNewProductsMenuItem:   false,
	DisplayProductSearchMenuItem: true,
	DisplayCustomerInfoMenuItem:  true,
	DisplayBlogMenuItem:          false,
	DisplayForumsMenuItem:        true,
	DisplayContactUsMenuItem:     false,
}

func TestDisplayDefaultMenuItemSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDisplayDefaultMenuItemSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDisplayDefaultMenuItemSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDisplayDefaultMenuItemSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDisplayDefaultMenuItemSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDisplayDefaultMenuItemSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDisplayDefaultMenuItemSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDisplayDefaultMenuItemSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDisplayDefaultMenuItemSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDisplayDefaultMenuItemSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDisplayDefaultMenuItemSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDisplayDefaultMenuItemSettings).Return(nil, nil).Once()

	repo := repository.NewDisplayDefaultMenuItemSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDisplayDefaultMenuItemSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDisplayDefaultMenuItemSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDisplayDefaultMenuItemSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDisplayDefaultMenuItemSettings.ID}
	update := bson.M{"$set": mockItemDisplayDefaultMenuItemSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDisplayDefaultMenuItemSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDisplayDefaultMenuItemSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
