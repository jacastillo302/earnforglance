package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionGdprConsent = "gdpr_consents"
)

// GdprConsent represents a GDPR consent
type GdprConsent struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	Message                   string             `bson:"message"`
	IsRequired                bool               `bson:"is_required"`
	RequiredMessage           string             `bson:"required_message"`
	DisplayDuringRegistration bool               `bson:"display_during_registration"`
	DisplayOnCustomerInfoPage bool               `bson:"display_on_customer_info_page"`
	DisplayOrder              int                `bson:"display_order"`
}
