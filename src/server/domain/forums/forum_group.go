package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumGroup = "forum_groups"
)

// ForumGroup represents a forum group
type ForumGroup struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc time.Time          `bson:"updated_on_utc"`
}
