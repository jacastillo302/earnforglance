package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionDiscount = "discounts"
)

// Discount represents a discount
type Discount struct {
	ID                        bson.ObjectID  `bson:"_id,omitempty"`
	Name                      string         `bson:"name"`
	AdminComment              string         `bson:"admin_comment"`
	DiscountTypeID            int            `bson:"discount_type_id"`
	UsePercentage             bool           `bson:"use_percentage"`
	DiscountPercentage        float64        `bson:"discount_percentage"`
	DiscountAmount            float64        `bson:"discount_amount"`
	MaximumDiscountAmount     *float64       `bson:"maximum_discount_amount"`
	StartDateUtc              *time.Time     `bson:"start_date_utc"`
	EndDateUtc                *time.Time     `bson:"end_date_utc"`
	RequiresCouponCode        bool           `bson:"requires_coupon_code"`
	CouponCode                string         `bson:"coupon_code"`
	IsCumulative              bool           `bson:"is_cumulative"`
	DiscountLimitationID      int            `bson:"discount_limitation_id"`
	LimitationTimes           int            `bson:"limitation_times"`
	MaximumDiscountedQuantity *int           `bson:"maximum_discounted_quantity"`
	AppliedToSubCategories    bool           `bson:"applied_to_sub_categories"`
	IsActive                  bool           `bson:"is_active"`
	VendorID                  *bson.ObjectID `bson:"vendor_id"`
}

type DiscountRepository interface {
	CreateMany(c context.Context, items []Discount) error
	Create(c context.Context, discount *Discount) error
	Update(c context.Context, discount *Discount) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Discount, error)
	FetchByID(c context.Context, ID string) (Discount, error)
}

type DiscountUsecase interface {
	CreateMany(c context.Context, items []Discount) error
	FetchByID(c context.Context, ID string) (Discount, error)
	Create(c context.Context, discount *Discount) error
	Update(c context.Context, discount *Discount) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Discount, error)
}
