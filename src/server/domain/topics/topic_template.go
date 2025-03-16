package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectioTopicTemplate = "topic_templates"
)

// TopicTemplate represents a topic template
type TopicTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	ViewPath     string             `bson:"view_path"`
	DisplayOrder int                `bson:"display_order"`
}
