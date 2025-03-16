package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionMessageTemplatesSettings = "message_templates_settings"
)

// MessageTemplatesSettings represents message templates settings
type MessageTemplatesSettings struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	CaseInvariantReplacement bool               `bson:"case_invariant_replacement"`
	Color1                   string             `bson:"color1"`
	Color2                   string             `bson:"color2"`
	Color3                   string             `bson:"color3"`
}
