package domain

import (
	"context" // added context library
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrderItem = "order_items"
)

// OrderItem represents an order item
type OrderItem struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty"`
	OrderItemGuid         uuid.UUID           `bson:"order_item_guid"`
	OrderID               primitive.ObjectID  `bson:"order_id"`
	ProductID             primitive.ObjectID  `bson:"product_id"`
	Quantity              int                 `bson:"quantity"`
	UnitPriceInclTax      float64             `bson:"unit_price_incl_tax"`
	UnitPriceExclTax      float64             `bson:"unit_price_excl_tax"`
	PriceInclTax          float64             `bson:"price_incl_tax"`
	PriceExclTax          float64             `bson:"price_excl_tax"`
	DiscountAmountInclTax float64             `bson:"discount_amount_incl_tax"`
	DiscountAmountExclTax float64             `bson:"discount_amount_excl_tax"`
	OriginalProductCost   float64             `bson:"original_product_cost"`
	AttributeDescription  string              `bson:"attribute_description"`
	AttributesXml         string              `bson:"attributes_xml"`
	DownloadCount         int                 `bson:"download_count"`
	IsDownloadActivated   bool                `bson:"is_download_activated"`
	LicenseDownloadID     *primitive.ObjectID `bson:"license_download_id,omitempty"`
	ItemWeight            *float64            `bson:"item_weight,omitempty"`
	RentalStartDateUtc    *time.Time          `bson:"rental_start_date_utc,omitempty"`
	RentalEndDateUtc      *time.Time          `bson:"rental_end_date_utc,omitempty"`
}

// OrderItemRepository represents the repository interface for OrderItem
type OrderItemRepository interface {
	CreateMany(c context.Context, items []OrderItem) error
	Create(c context.Context, order_item *OrderItem) error
	Update(c context.Context, order_item *OrderItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]OrderItem, error)
	FetchByID(c context.Context, ID string) (OrderItem, error)
}

// OrderItemUsecase represents the usecase interface for OrderItem
type OrderItemUsecase interface {
	FetchByID(c context.Context, ID string) (OrderItem, error)
	Create(c context.Context, order_item *OrderItem) error
	Update(c context.Context, order_item *OrderItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]OrderItem, error)
}
