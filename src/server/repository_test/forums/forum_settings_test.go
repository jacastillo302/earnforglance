package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/forums"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultForumSettings struct {
	mock.Mock
}

func (m *MockSingleResultForumSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ForumSettings); ok {
		*v.(*domain.ForumSettings) = *result
	}
	return args.Error(1)
}

var mockItemForumSettings = &domain.ForumSettings{
	ID:                                  primitive.NewObjectID(), // Existing ID of the record to update
	ForumsEnabled:                       false,
	RelativeDateTimeFormattingEnabled:   false,
	AllowCustomersToEditPosts:           false,
	AllowCustomersToManageSubscriptions: false,
	AllowGuestsToCreatePosts:            true,
	AllowGuestsToCreateTopics:           true,
	AllowCustomersToDeletePosts:         false,
	AllowPostVoting:                     false,
	MaxVotesPerDay:                      5,
	TopicSubjectMaxLength:               50,
	StrippedTopicMaxLength:              25,
	PostMaxLength:                       500,
	TopicsPageSize:                      10,
	PostsPageSize:                       5,
	SearchResultsPageSize:               10,
	ActiveDiscussionsPageSize:           3,
	LatestCustomerPostsPageSize:         5,
	ShowCustomersPostCount:              false,
	EditorTypeID:                        1,
	SignaturesEnabled:                   false,
	AllowPrivateMessages:                false,
	ShowAlertForPM:                      false,
	PrivateMessagesPageSize:             5,
	ForumSubscriptionsPageSize:          5,
	NotifyAboutPrivateMessages:          false,
	PMSubjectMaxLength:                  25,
	PMTextMaxLength:                     250,
	HomepageActiveDiscussionsTopicCount: 3,
	ActiveDiscussionsFeedCount:          5,
	ActiveDiscussionsFeedEnabled:        false,
	ForumFeedsEnabled:                   false,
	ForumFeedCount:                      5,
	ForumSearchTermMinimumLength:        2,
}

func TestForumSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForumSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForumSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForumSettings).Return(nil, nil).Once()

	repo := repository.NewForumSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForumSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForumSettings.ID}
	update := bson.M{"$set": mockItemForumSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForumSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
