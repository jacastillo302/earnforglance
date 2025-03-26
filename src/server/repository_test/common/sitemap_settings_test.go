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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSitemapSettings struct {
	mock.Mock
}

func (m *MockSingleResultSitemapSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SitemapSettings); ok {
		*v.(*domain.SitemapSettings) = *result
	}
	return args.Error(1)
}

var mockItemSitemapSettings = &domain.SitemapSettings{
	ID:                             primitive.NewObjectID(), // Existing ID of the record to update
	SitemapEnabled:                 false,
	SitemapPageSize:                50,
	SitemapIncludeBlogPosts:        false,
	SitemapIncludeCategories:       false,
	SitemapIncludeManufacturers:    true,
	SitemapIncludeNews:             false,
	SitemapIncludeProducts:         false,
	SitemapIncludeSitemapSettingss: true,
	SitemapIncludeTopics:           false,
}

func TestSitemapSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSitemapSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSitemapSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSitemapSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSitemapSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSitemapSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSitemapSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSitemapSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSitemapSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSitemapSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSitemapSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSitemapSettings).Return(nil, nil).Once()

	repo := repository.NewSitemapSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSitemapSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSitemapSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSitemapSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSitemapSettings.ID}
	update := bson.M{"$set": mockItemSitemapSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSitemapSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSitemapSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
