package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/blogs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBlogSettings struct {
	mock.Mock
}

func (m *MockSingleResultBlogSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BlogSettings); ok {
		*v.(*domain.BlogSettings) = *result
	}
	return args.Error(1)
}

func TestBlogSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBlogSettings

	mockItem := domain.BlogSettings{
		ID:                                     primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Enabled:                                false,
		PostsPageSize:                          0,
		AllowNotRegisteredUsersToLeaveComments: false,
		NotifyAboutNewBlogComments:             false,
		NumberOfTags:                           0,
		ShowHeaderRssUrl:                       false,
		BlogCommentsMustBeApproved:             false,
		ShowBlogCommentsPerStore:               false,
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogSettingsRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBlogSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogSettings := &domain.BlogSettings{
		Enabled:                                true,
		PostsPageSize:                          10,
		AllowNotRegisteredUsersToLeaveComments: true,
		NotifyAboutNewBlogComments:             false,
		NumberOfTags:                           5,
		ShowHeaderRssUrl:                       true,
		BlogCommentsMustBeApproved:             false,
		ShowBlogCommentsPerStore:               true,
	}

	collectionHelper.On("InsertOne", mock.Anything, mockBlogSettings).Return(nil, nil).Once()

	repo := repository.NewBlogSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBlogSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBlogSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogSettings := &domain.BlogSettings{
		ID:                                     primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                                false,
		PostsPageSize:                          20,
		AllowNotRegisteredUsersToLeaveComments: false,
		NotifyAboutNewBlogComments:             true,
		NumberOfTags:                           10,
		ShowHeaderRssUrl:                       false,
		BlogCommentsMustBeApproved:             true,
		ShowBlogCommentsPerStore:               false,
	}

	filter := bson.M{"_id": mockBlogSettings.ID}
	update := bson.M{"$set": mockBlogSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBlogSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBlogSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
