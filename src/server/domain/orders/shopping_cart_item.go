package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShoppingCartItem = "shopping_cart_items"
)

// ShoppingCartItem represents a shopping cart item
type ShoppingCartItem struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	StoreID              int                `bson:"store_id"`
	ShoppingCartTypeID   int                `bson:"shopping_cart_type_id"`
	CustomerID           int                `bson:"customer_id"`
	ProductID            int                `bson:"product_id"`
	AttributesXml        string             `bson:"attributes_xml"`
	CustomerEnteredPrice float64            `bson:"customer_entered_price"`
	Quantity             int                `bson:"quantity"`
	RentalStartDateUtc   *time.Time         `bson:"rental_start_date_utc,omitempty"`
	RentalEndDateUtc     *time.Time         `bson:"rental_end_date_utc,omitempty"`
	CreatedOnUtc         time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc         time.Time          `bson:"updated_on_utc"`
	ShoppingCartType     ShoppingCartType   `bson:"shopping_cart_type"`
}
