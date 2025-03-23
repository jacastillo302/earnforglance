package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/news"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/news"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultNewsItem struct {
	mock.Mock
}

func (m *MockSingleResultNewsItem) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.NewsItem); ok {
		*v.(*domain.NewsItem) = *result
	}
	return args.Error(1)
}

var mockItemNewsItem = &domain.NewsItem{
	ID:              primitive.NewObjectID(), // Existing ID of the record to update
	LanguageID:      primitive.NewObjectID(),
	Title:           "Updated Feature Announcement",
	Short:           "We have updated the feature announcement.",
	Full:            "The new feature has been updated to include additional functionality.",
	Published:       false,
	StartDateUtc:    new(time.Time),
	EndDateUtc:      new(time.Time),
	AllowComments:   false,
	LimitedToStores: true,
	MetaKeywords:    "feature, update",
	MetaDescription: "Updated announcement of a feature.",
	MetaTitle:       "Updated Feature Announcement",
	CreatedOnUtc:    time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestNewsItemRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionNewsItem

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemNewsItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestNewsItemRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemNewsItem).Return(nil, nil).Once()

	repo := repository.NewNewsItemRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemNewsItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestNewsItemRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemNewsItem.ID}
	update := bson.M{"$set": mockItemNewsItem}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewNewsItemRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemNewsItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
