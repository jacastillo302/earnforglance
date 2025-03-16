package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionDiscountRequirement = "discount_requirements"
)

// DiscountRequirement represents a discount requirement
type DiscountRequirement struct {
	ID                                primitive.ObjectID               `bson:"_id,omitempty"`
	DiscountID                        int                              `bson:"discount_id"`
	DiscountRequirementRuleSystemName string                           `bson:"discount_requirement_rule_system_name"`
	ParentID                          *int                             `bson:"parent_id,omitempty"`
	InteractionTypeID                 *int                             `bson:"interaction_type_id,omitempty"`
	IsGroup                           bool                             `bson:"is_group"`
	InteractionType                   *RequirementGroupInteractionType `bson:"interaction_type,omitempty"`
}

// NewDiscountRequirement creates a new DiscountRequirement instance
func NewDiscountRequirement(discountID int, ruleSystemName string) *DiscountRequirement {
	return &DiscountRequirement{
		DiscountID:                        discountID,
		DiscountRequirementRuleSystemName: ruleSystemName,
	}
}
