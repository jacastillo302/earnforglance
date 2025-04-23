package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionBlogPostTag = "blog_post_tags"
)

// BlogPostTag represents a blog post tag
type BlogPostTag struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Name          string        `bson:"name"`
	BlogPostCount int           `bson:"blog_post_count"`
}

type BlogPostTagRepository interface {
	CreateMany(c context.Context, items []BlogPostTag) error
	Create(c context.Context, blog_post_tag *BlogPostTag) error
	Update(c context.Context, blog_post_tag *BlogPostTag) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BlogPostTag, error)
	FetchByID(c context.Context, ID string) (BlogPostTag, error)
}

type BlogPostTagUsecase interface {
	CreateMany(c context.Context, items []BlogPostTag) error
	FetchByID(c context.Context, ID string) (BlogPostTag, error)
	Create(c context.Context, blog_post_tag *BlogPostTag) error
	Update(c context.Context, blog_post_tag *BlogPostTag) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BlogPostTag, error)
}
