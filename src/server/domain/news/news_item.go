package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionNewsItem = "news_items"
)

// NewsItem represents a news item
type NewsItem struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Title           string        `bson:"title"`
	Short           string        `bson:"short"`
	Full            string        `bson:"full"`
	Published       bool          `bson:"published"`
	StartDateUtc    *time.Time    `bson:"start_date_utc"`
	EndDateUtc      *time.Time    `bson:"end_date_utc"`
	AllowComments   bool          `bson:"allow_comments"`
	LimitedToStores bool          `bson:"limited_to_stores"`
	MetaKeywords    string        `bson:"meta_keywords"`
	MetaDescription string        `bson:"meta_description"`
	MetaTitle       string        `bson:"meta_title"`
	CreatedOnUtc    time.Time     `bson:"created_on_utc"`
}

// NewsItemRepository interface
type NewsItemRepository interface {
	CreateMany(c context.Context, items []NewsItem) error
	Create(c context.Context, news_item *NewsItem) error
	Update(c context.Context, news_item *NewsItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsItem, error)
	FetchByID(c context.Context, ID string) (NewsItem, error)
}

// NewsItemUsecase interface
type NewsItemUsecase interface {
	CreateMany(c context.Context, items []NewsItem) error
	FetchByID(c context.Context, ID string) (NewsItem, error)
	Create(c context.Context, news_item *NewsItem) error
	Update(c context.Context, news_item *NewsItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsItem, error)
}
