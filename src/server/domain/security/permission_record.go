package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPermissionRecord = "permission_records"
)

// PermissionRecord represents a permission record
type PermissionRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	SystemName string             `bson:"system_name"`
	Category   string             `bson:"category"`
}
