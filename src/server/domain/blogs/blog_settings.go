package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionBlogSettings = "blog_settings"
)

// BlogSettings represents blog settings
type BlogSettings struct {
	ID                                     primitive.ObjectID `bson:"_id,omitempty"`
	Enabled                                bool               `bson:"enabled"`
	PostsPageSize                          int                `bson:"posts_page_size"`
	AllowNotRegisteredUsersToLeaveComments bool               `bson:"allow_not_registered_users_to_leave_comments"`
	NotifyAboutNewBlogComments             bool               `bson:"notify_about_new_blog_comments"`
	NumberOfTags                           int                `bson:"number_of_tags"`
	ShowHeaderRssUrl                       bool               `bson:"show_header_rss_url"`
	BlogCommentsMustBeApproved             bool               `bson:"blog_comments_must_be_approved"`
	ShowBlogCommentsPerStore               bool               `bson:"show_blog_comments_per_store"`
}
