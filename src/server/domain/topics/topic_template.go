package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTopicTemplate = "topic_templates"
)

// TopicTemplate represents a topic template.
type TopicTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	ViewPath     string             `bson:"view_path"`
	DisplayOrder int                `bson:"display_order"`
}

// TopicTemplateRepository defines the repository interface for TopicTemplate
type TopicTemplateRepository interface {
	Create(c context.Context, topic_template *TopicTemplate) error
	Update(c context.Context, topic_template *TopicTemplate) error
	Delete(c context.Context, topic_template *TopicTemplate) error
	Fetch(c context.Context) ([]TopicTemplate, error)
	FetchByID(c context.Context, topic_templateID string) (TopicTemplate, error)
}

// TopicTemplateUsecase defines the use case interface for TopicTemplate
type TopicTemplateUsecase interface {
	FetchByID(c context.Context, topic_templateID string) (TopicTemplate, error)
	Create(c context.Context, topic_template *TopicTemplate) error
	Update(c context.Context, topic_template *TopicTemplate) error
	Delete(c context.Context, topic_template *TopicTemplate) error
	Fetch(c context.Context) ([]TopicTemplate, error)
}
