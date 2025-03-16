package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionBlogPostTag = "blog_post_tags"
)

// BlogPostTag represents a blog post tag
type BlogPostTag struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	BlogPostCount int                `bson:"blog_post_count"`
}
