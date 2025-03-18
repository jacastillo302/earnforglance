package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlogPostTag = "blog_post_tags"
)

// BlogPostTag represents a blog post tag
type BlogPostTag struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	BlogPostCount int                `bson:"blog_post_count"`
}

type BlogPostTagRepository interface {
	Create(c context.Context, blog_post_tag *BlogPostTag) error
	Update(c context.Context, blog_post_tag *BlogPostTag) error
	Delete(c context.Context, blog_post_tag *BlogPostTag) error
	Fetch(c context.Context) ([]BlogPostTag, error)
	FetchByID(c context.Context, blog_post_tagID string) (BlogPostTag, error)
}

type BlogPostTagUsecase interface {
	FetchByID(c context.Context, blog_post_tagID string) (BlogPostTag, error)
	Create(c context.Context, blog_post_tag *BlogPostTag) error
	Update(c context.Context, blog_post_tag *BlogPostTag) error
	Delete(c context.Context, blog_post_tag *BlogPostTag) error
	Fetch(c context.Context) ([]BlogPostTag, error)
}
