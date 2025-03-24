package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTopic = "topics"
)

// Topic represents a topic.
type Topic struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	SystemName                string             `bson:"system_name"`
	IncludeInSitemap          bool               `bson:"include_in_sitemap"`
	IncludeInTopMenu          bool               `bson:"include_in_top_menu"`
	IncludeInFooterColumn1    bool               `bson:"include_in_footer_column1"`
	IncludeInFooterColumn2    bool               `bson:"include_in_footer_column2"`
	IncludeInFooterColumn3    bool               `bson:"include_in_footer_column3"`
	DisplayOrder              int                `bson:"display_order"`
	AccessibleWhenStoreClosed bool               `bson:"accessible_when_store_closed"`
	IsPasswordProtected       bool               `bson:"is_password_protected"`
	Password                  string             `bson:"password"`
	Title                     string             `bson:"title"`
	Body                      string             `bson:"body"`
	Published                 bool               `bson:"published"`
	TopicTemplateID           primitive.ObjectID `bson:"topic_template_id"`
	MetaKeywords              string             `bson:"meta_keywords"`
	MetaDescription           string             `bson:"meta_description"`
	MetaTitle                 string             `bson:"meta_title"`
	SubjectToAcl              bool               `bson:"subject_to_acl"`
	LimitedToStores           bool               `bson:"limited_to_stores"`
	AvailableStartDateTimeUtc *time.Time         `bson:"available_start_date_time_utc,omitempty"`
	AvailableEndDateTimeUtc   *time.Time         `bson:"available_end_date_time_utc,omitempty"`
}

// TopicRepository defines the repository interface for Topic
type TopicRepository interface {
	CreateMany(c context.Context, items []Topic) error
	Create(c context.Context, topic *Topic) error
	Update(c context.Context, topic *Topic) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Topic, error)
	FetchByID(c context.Context, ID string) (Topic, error)
}

// TopicUsecase defines the use case interface for Topic
type TopicUsecase interface {
	CreateMany(c context.Context, items []Topic) error
	FetchByID(c context.Context, ID string) (Topic, error)
	Create(c context.Context, topic *Topic) error
	Update(c context.Context, topic *Topic) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Topic, error)
}
