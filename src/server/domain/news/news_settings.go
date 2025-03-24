package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNewsSettings = "news_settings"
)

// NewsSettings represents news settings.
type NewsSettings struct {
	ID                                     primitive.ObjectID `bson:"_id,omitempty"`
	Enabled                                bool               `bson:"enabled"`
	AllowNotRegisteredUsersToLeaveComments bool               `bson:"allow_not_registered_users_to_leave_comments"`
	NotifyAboutNewNewsComments             bool               `bson:"notify_about_new_news_comments"`
	ShowNewsOnMainPage                     bool               `bson:"show_news_on_main_page"`
	MainPageNewsCount                      int                `bson:"main_page_news_count"`
	NewsArchivePageSize                    int                `bson:"news_archive_page_size"`
	ShowHeaderRssUrl                       bool               `bson:"show_header_rss_url"`
	NewsCommentsMustBeApproved             bool               `bson:"news_comments_must_be_approved"`
	ShowNewsCommentsPerStore               bool               `bson:"show_news_comments_per_store"`
}

// NewsSettingsRepository interface
type NewsSettingsRepository interface {
	CreateMany(c context.Context, items []NewsSettings) error
	Create(c context.Context, news_settings *NewsSettings) error
	Update(c context.Context, news_settings *NewsSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsSettings, error)
	FetchByID(c context.Context, ID string) (NewsSettings, error)
}

// NewsSettingsUsecase interface
type NewsSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (NewsSettings, error)
	Create(c context.Context, news_settings *NewsSettings) error
	Update(c context.Context, news_settings *NewsSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsSettings, error)
}
