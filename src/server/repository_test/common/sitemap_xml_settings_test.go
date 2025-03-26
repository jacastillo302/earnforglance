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

type MockSingleResultSitemapXmlSettings struct {
	mock.Mock
}

func (m *MockSingleResultSitemapXmlSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SitemapXmlSettings); ok {
		*v.(*domain.SitemapXmlSettings) = *result
	}
	return args.Error(1)
}

var mockItemSitemapXmlSettings = &domain.SitemapXmlSettings{
	ID:                                   primitive.NewObjectID(), // Existing ID of the record to update
	SitemapXmlEnabled:                    false,
	SitemapXmlIncludeBlogPosts:           false,
	SitemapXmlIncludeCategories:          false,
	SitemapXmlIncludeCustomUrls:          false,
	SitemapXmlIncludeManufacturers:       true,
	SitemapXmlIncludeNews:                false,
	SitemapXmlIncludeProducts:            false,
	SitemapXmlIncludeSitemapXmlSettingss: true,
	SitemapXmlIncludeTopics:              false,
	SitemapCustomUrls:                    []string{"https://example.com/updated1", "https://example.com/updated2"},
	RebuildSitemapXmlAfterHours:          48,
	SitemapBuildOperationDelay:           10,
}

func TestSitemapXmlSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSitemapXmlSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSitemapXmlSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSitemapXmlSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSitemapXmlSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSitemapXmlSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSitemapXmlSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSitemapXmlSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSitemapXmlSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSitemapXmlSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSitemapXmlSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSitemapXmlSettings).Return(nil, nil).Once()

	repo := repository.NewSitemapXmlSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSitemapXmlSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSitemapXmlSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSitemapXmlSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSitemapXmlSettings.ID}
	update := bson.M{"$set": mockItemSitemapXmlSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSitemapXmlSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSitemapXmlSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
