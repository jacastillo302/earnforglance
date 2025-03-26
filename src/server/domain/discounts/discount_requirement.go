package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscountRequirement = "discount_requirements"
)

// DiscountRequirement represents a discount requirement
type DiscountRequirement struct {
	ID                                primitive.ObjectID  `bson:"_id,omitempty"`
	DiscountID                        primitive.ObjectID  `bson:"discount_id"`
	DiscountRequirementRuleSystemName string              `bson:"discount_requirement_rule_system_name"`
	ParentID                          *primitive.ObjectID `bson:"parent_id"`
	InteractionTypeID                 *int                `bson:"interaction_type_id"`
	IsGroup                           bool                `bson:"is_group"`
}

type DiscountRequirementRepository interface {
	CreateMany(c context.Context, items []DiscountRequirement) error
	Create(c context.Context, discount_requirement *DiscountRequirement) error
	Update(c context.Context, discount_requirement *DiscountRequirement) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountRequirement, error)
	FetchByID(c context.Context, ID string) (DiscountRequirement, error)
}

type DiscountRequirementUsecase interface {
	CreateMany(c context.Context, items []DiscountRequirement) error
	FetchByID(c context.Context, ID string) (DiscountRequirement, error)
	Create(c context.Context, discount_requirement *DiscountRequirement) error
	Update(c context.Context, discount_requirement *DiscountRequirement) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountRequirement, error)
}

// NewDiscountRequirement creates a new DiscountRequirement instance
func NewDiscountRequirement(discountID primitive.ObjectID, ruleSystemName string) *DiscountRequirement {
	return &DiscountRequirement{
		DiscountID:                        discountID,
		DiscountRequirementRuleSystemName: ruleSystemName,
	}
}
