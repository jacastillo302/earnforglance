package domain

import (
	"context"
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

type ForumGroupRepository interface {
	Create(c context.Context, forum_group *ForumGroup) error
	Update(c context.Context, forum_group *ForumGroup) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumGroup, error)
	FetchByID(c context.Context, ID string) (ForumGroup, error)
}

type ForumGroupUsecase interface {
	FetchByID(c context.Context, ID string) (ForumGroup, error)
	Create(c context.Context, forum_group *ForumGroup) error
	Update(c context.Context, forum_group *ForumGroup) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumGroup, error)
}
