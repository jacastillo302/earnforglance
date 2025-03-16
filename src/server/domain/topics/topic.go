package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectioTopic = "topics"
)

// Topic represents a topic
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
	TopicTemplateID           int                `bson:"topic_template_id"`
	MetaKeywords              string             `bson:"meta_keywords"`
	MetaDescription           string             `bson:"meta_description"`
	MetaTitle                 string             `bson:"meta_title"`
	SubjectToAcl              bool               `bson:"subject_to_acl"`
	LimitedToStores           bool               `bson:"limited_to_stores"`
	AvailableStartDateTimeUtc *time.Time         `bson:"available_start_date_time_utc,omitempty"`
	AvailableEndDateTimeUtc   *time.Time         `bson:"available_end_date_time_utc,omitempty"`
}
