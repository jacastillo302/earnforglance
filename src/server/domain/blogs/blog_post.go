package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlogPost = "blog_posts"
)

// BlogPost represents a blog post
type BlogPost struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	LanguageID       primitive.ObjectID `bson:"language_id"`
	IncludeInSitemap bool               `bson:"include_in_sitemap"`
	Title            string             `bson:"title"`
	Body             string             `bson:"body"`
	BodyOverview     string             `bson:"body_overview"`
	AllowComments    bool               `bson:"allow_comments"`
	Tags             string             `bson:"tags"`
	StartDateUtc     *time.Time         `bson:"start_date_utc"`
	EndDateUtc       *time.Time         `bson:"end_date_utc"`
	MetaKeywords     string             `bson:"meta_keywords"`
	MetaDescription  string             `bson:"meta_description"`
	MetaTitle        string             `bson:"meta_title"`
	LimitedToStores  bool               `bson:"limited_to_stores"`
	CreatedOnUtc     time.Time          `bson:"created_on_utc"`
}

type BlogPostRepository interface {
	CreateMany(c context.Context, items []BlogPost) error
	Create(c context.Context, blog_post *BlogPost) error
	Update(c context.Context, blog_post *BlogPost) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BlogPost, error)
	FetchByID(c context.Context, ID string) (BlogPost, error)
}

type BlogPostUsecase interface {
	CreateMany(c context.Context, items []BlogPost) error
	FetchByID(c context.Context, ID string) (BlogPost, error)
	Create(c context.Context, blog_post *BlogPost) error
	Update(c context.Context, blog_post *BlogPost) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BlogPost, error)
}
