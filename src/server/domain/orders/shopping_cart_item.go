package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionShoppingCartItem = "shopping_cart_items"
)

// ShoppingCartItem represents a shopping cart item
type ShoppingCartItem struct {
	ID                   bson.ObjectID `bson:"_id,omitempty"`
	StoreID              bson.ObjectID `bson:"store_id"`
	ShoppingCartTypeID   int           `bson:"shopping_cart_type_id"`
	CustomerID           bson.ObjectID `bson:"customer_id"`
	ProductID            bson.ObjectID `bson:"product_id"`
	AttributesXml        string        `bson:"attributes_xml"`
	CustomerEnteredPrice float64       `bson:"customer_entered_price"`
	Quantity             int           `bson:"quantity"`
	RentalStartDateUtc   *time.Time    `bson:"rental_start_date_utc"`
	RentalEndDateUtc     *time.Time    `bson:"rental_end_date_utc"`
	CreatedOnUtc         time.Time     `bson:"created_on_utc"`
	UpdatedOnUtc         time.Time     `bson:"updated_on_utc"`
}

// ShoppingCartItemRepository defines the repository interface for ShoppingCartItem
type ShoppingCartItemRepository interface {
	CreateMany(c context.Context, items []ShoppingCartItem) error
	Create(c context.Context, shopping_cart_item *ShoppingCartItem) error
	Update(c context.Context, shopping_cart_item *ShoppingCartItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShoppingCartItem, error)
	FetchByID(c context.Context, ID string) (ShoppingCartItem, error)
}

// ShoppingCartItemUsecase defines the use case interface for ShoppingCartItem
type ShoppingCartItemUsecase interface {
	CreateMany(c context.Context, items []ShoppingCartItem) error
	FetchByID(c context.Context, ID string) (ShoppingCartItem, error)
	Create(c context.Context, shopping_cart_item *ShoppingCartItem) error
	Update(c context.Context, shopping_cart_item *ShoppingCartItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShoppingCartItem, error)
}
