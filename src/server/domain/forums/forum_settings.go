package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionForumSettings = "forum_settings"
)

// ForumSettings represents forum settings
type ForumSettings struct {
	ID                                  bson.ObjectID `bson:"_id,omitempty"`
	ForumsEnabled                       bool          `bson:"forums_enabled"`
	RelativeDateTimeFormattingEnabled   bool          `bson:"relative_date_time_formatting_enabled"`
	AllowCustomersToEditPosts           bool          `bson:"allow_customers_to_edit_posts"`
	AllowCustomersToManageSubscriptions bool          `bson:"allow_customers_to_manage_subscriptions"`
	AllowGuestsToCreatePosts            bool          `bson:"allow_guests_to_create_posts"`
	AllowGuestsToCreateTopics           bool          `bson:"allow_guests_to_create_topics"`
	AllowCustomersToDeletePosts         bool          `bson:"allow_customers_to_delete_posts"`
	AllowPostVoting                     bool          `bson:"allow_post_voting"`
	MaxVotesPerDay                      int           `bson:"max_votes_per_day"`
	TopicSubjectMaxLength               int           `bson:"topic_subject_max_length"`
	StrippedTopicMaxLength              int           `bson:"stripped_topic_max_length"`
	PostMaxLength                       int           `bson:"post_max_length"`
	TopicsPageSize                      int           `bson:"topics_page_size"`
	PostsPageSize                       int           `bson:"posts_page_size"`
	SearchResultsPageSize               int           `bson:"search_results_page_size"`
	ActiveDiscussionsPageSize           int           `bson:"active_discussions_page_size"`
	LatestCustomerPostsPageSize         int           `bson:"latest_customer_posts_page_size"`
	ShowCustomersPostCount              bool          `bson:"show_customers_post_count"`
	EditorTypeID                        int           `bson:"forum_editor"`
	SignaturesEnabled                   bool          `bson:"signatures_enabled"`
	AllowPrivateMessages                bool          `bson:"allow_private_messages"`
	ShowAlertForPM                      bool          `bson:"show_alert_for_pm"`
	PrivateMessagesPageSize             int           `bson:"private_messages_page_size"`
	ForumSubscriptionsPageSize          int           `bson:"forum_subscriptions_page_size"`
	NotifyAboutPrivateMessages          bool          `bson:"notify_about_private_messages"`
	PMSubjectMaxLength                  int           `bson:"pm_subject_max_length"`
	PMTextMaxLength                     int           `bson:"pm_text_max_length"`
	HomepageActiveDiscussionsTopicCount int           `bson:"homepage_active_discussions_topic_count"`
	ActiveDiscussionsFeedCount          int           `bson:"active_discussions_feed_count"`
	ActiveDiscussionsFeedEnabled        bool          `bson:"active_discussions_feed_enabled"`
	ForumFeedsEnabled                   bool          `bson:"forum_feeds_enabled"`
	ForumFeedCount                      int           `bson:"forum_feed_count"`
	ForumSearchTermMinimumLength        int           `bson:"forum_search_term_minimum_length"`
}

// ForumSettingsRepository represents the forum settings repository interface
type ForumSettingsRepository interface {
	CreateMany(c context.Context, items []ForumSettings) error
	Create(c context.Context, forum_settings *ForumSettings) error
	Update(c context.Context, forum_settings *ForumSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumSettings, error)
	FetchByID(c context.Context, ID string) (ForumSettings, error)
}

// ForumSettingsUsecase represents the forum settings usecase interface
type ForumSettingsUsecase interface {
	CreateMany(c context.Context, items []ForumSettings) error
	FetchByID(c context.Context, ID string) (ForumSettings, error)
	Create(c context.Context, forum_settings *ForumSettings) error
	Update(c context.Context, forum_settings *ForumSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumSettings, error)
}
