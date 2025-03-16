package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscount = "discounts"
)

// Discount represents a discount
type Discount struct {
	ID                        primitive.ObjectID     `bson:"_id,omitempty"`
	Name                      string                 `bson:"name"`
	AdminComment              string                 `bson:"admin_comment"`
	DiscountTypeID            int                    `bson:"discount_type_id"`
	UsePercentage             bool                   `bson:"use_percentage"`
	DiscountPercentage        float64                `bson:"discount_percentage"`
	DiscountAmount            float64                `bson:"discount_amount"`
	MaximumDiscountAmount     *float64               `bson:"maximum_discount_amount,omitempty"`
	StartDateUtc              *time.Time             `bson:"start_date_utc,omitempty"`
	EndDateUtc                *time.Time             `bson:"end_date_utc,omitempty"`
	RequiresCouponCode        bool                   `bson:"requires_coupon_code"`
	CouponCode                string                 `bson:"coupon_code"`
	IsCumulative              bool                   `bson:"is_cumulative"`
	DiscountLimitationID      int                    `bson:"discount_limitation_id"`
	LimitationTimes           int                    `bson:"limitation_times"`
	MaximumDiscountedQuantity *int                   `bson:"maximum_discounted_quantity,omitempty"`
	AppliedToSubCategories    bool                   `bson:"applied_to_sub_categories"`
	IsActive                  bool                   `bson:"is_active"`
	VendorID                  *int                   `bson:"vendor_id,omitempty"`
	DiscountType              DiscountType           `bson:"discount_type"`
	DiscountLimitation        DiscountLimitationType `bson:"discount_limitation"`
}
