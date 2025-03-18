package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNewsItem = "news_items"
)

// NewsItem represents a news item
type NewsItem struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	LanguageID      primitive.ObjectID `bson:"language_id"`
	Title           string             `bson:"title"`
	Short           string             `bson:"short"`
	Full            string             `bson:"full"`
	Published       bool               `bson:"published"`
	StartDateUtc    *time.Time         `bson:"start_date_utc,omitempty"`
	EndDateUtc      *time.Time         `bson:"end_date_utc,omitempty"`
	AllowComments   bool               `bson:"allow_comments"`
	LimitedToStores bool               `bson:"limited_to_stores"`
	MetaKeywords    string             `bson:"meta_keywords"`
	MetaDescription string             `bson:"meta_description"`
	MetaTitle       string             `bson:"meta_title"`
	CreatedOnUtc    time.Time          `bson:"created_on_utc"`
}

// NewsItemRepository interface
type NewsItemRepository interface {
	Create(c context.Context, news_item *NewsItem) error
	Update(c context.Context, news_item *NewsItem) error
	Delete(c context.Context, news_item *NewsItem) error
	Fetch(c context.Context) ([]NewsItem, error)
	FetchByID(c context.Context, news_itemID string) (NewsItem, error)
}

// NewsItemUsecase interface
type NewsItemUsecase interface {
	FetchByID(c context.Context, news_itemID string) (NewsItem, error)
	Create(c context.Context, news_item *NewsItem) error
	Update(c context.Context, news_item *NewsItem) error
	Delete(c context.Context, news_item *NewsItem) error
	Fetch(c context.Context) ([]NewsItem, error)
}
