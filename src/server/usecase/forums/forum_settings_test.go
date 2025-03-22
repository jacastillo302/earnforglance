package usecase

import (
	"context"
	domain "earnforglance/server/domain/forums"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestForumSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewForumSettingsUsecase(mockRepo, timeout)

	forumSettingsID := primitive.NewObjectID().Hex()

	updatedForumSettings := domain.ForumSettings{
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
		ForumEditor:                         1,
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

	mockRepo.On("FetchByID", mock.Anything, forumSettingsID).Return(updatedForumSettings, nil)

	result, err := usecase.FetchByID(context.Background(), forumSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForumSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestForumSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewForumSettingsUsecase(mockRepo, timeout)

	newForumSettings := &domain.ForumSettings{
		ForumsEnabled:                       true,
		RelativeDateTimeFormattingEnabled:   true,
		AllowCustomersToEditPosts:           true,
		AllowCustomersToManageSubscriptions: true,
		AllowGuestsToCreatePosts:            false,
		AllowGuestsToCreateTopics:           false,
		AllowCustomersToDeletePosts:         true,
		AllowPostVoting:                     true,
		MaxVotesPerDay:                      10,
		TopicSubjectMaxLength:               100,
		StrippedTopicMaxLength:              50,
		PostMaxLength:                       1000,
		TopicsPageSize:                      20,
		PostsPageSize:                       10,
		SearchResultsPageSize:               15,
		ActiveDiscussionsPageSize:           5,
		LatestCustomerPostsPageSize:         10,
		ShowCustomersPostCount:              true,
		ForumEditor:                         2,
		SignaturesEnabled:                   true,
		AllowPrivateMessages:                true,
		ShowAlertForPM:                      true,
		PrivateMessagesPageSize:             10,
		ForumSubscriptionsPageSize:          10,
		NotifyAboutPrivateMessages:          true,
		PMSubjectMaxLength:                  50,
		PMTextMaxLength:                     500,
		HomepageActiveDiscussionsTopicCount: 5,
		ActiveDiscussionsFeedCount:          10,
		ActiveDiscussionsFeedEnabled:        true,
		ForumFeedsEnabled:                   true,
		ForumFeedCount:                      10,
		ForumSearchTermMinimumLength:        3,
	}

	mockRepo.On("Create", mock.Anything, newForumSettings).Return(nil)

	err := usecase.Create(context.Background(), newForumSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewForumSettingsUsecase(mockRepo, timeout)

	updatedForumSettings := &domain.ForumSettings{
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
		ForumEditor:                         1,
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

	mockRepo.On("Update", mock.Anything, updatedForumSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedForumSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewForumSettingsUsecase(mockRepo, timeout)

	forumSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), forumSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewForumSettingsUsecase(mockRepo, timeout)

	fetchedForumSettings := []domain.ForumSettings{
		{
			ID:                                  primitive.NewObjectID(),
			ForumsEnabled:                       true,
			RelativeDateTimeFormattingEnabled:   true,
			AllowCustomersToEditPosts:           true,
			AllowCustomersToManageSubscriptions: true,
			AllowGuestsToCreatePosts:            false,
			AllowGuestsToCreateTopics:           false,
			AllowCustomersToDeletePosts:         true,
			AllowPostVoting:                     true,
			MaxVotesPerDay:                      10,
			TopicSubjectMaxLength:               100,
			StrippedTopicMaxLength:              50,
			PostMaxLength:                       1000,
			TopicsPageSize:                      20,
			PostsPageSize:                       10,
			SearchResultsPageSize:               15,
			ActiveDiscussionsPageSize:           5,
			LatestCustomerPostsPageSize:         10,
			ShowCustomersPostCount:              true,
			ForumEditor:                         2,
			SignaturesEnabled:                   true,
			AllowPrivateMessages:                true,
			ShowAlertForPM:                      true,
			PrivateMessagesPageSize:             10,
			ForumSubscriptionsPageSize:          10,
			NotifyAboutPrivateMessages:          true,
			PMSubjectMaxLength:                  50,
			PMTextMaxLength:                     500,
			HomepageActiveDiscussionsTopicCount: 5,
			ActiveDiscussionsFeedCount:          10,
			ActiveDiscussionsFeedEnabled:        true,
			ForumFeedsEnabled:                   true,
			ForumFeedCount:                      10,
			ForumSearchTermMinimumLength:        3,
		},
		{
			ID:                                  primitive.NewObjectID(),
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
			ForumEditor:                         1,
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
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForumSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForumSettings, result)
	mockRepo.AssertExpectations(t)
}
