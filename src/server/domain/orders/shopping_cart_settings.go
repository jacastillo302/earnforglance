package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShoppingCartSettings = "shopping_cart_settings"
)

// ShoppingCartSettings represents shopping cart settings.
type ShoppingCartSettings struct {
	ID                                          primitive.ObjectID `bson:"_id,omitempty"`
	DisplayCartAfterAddingProduct               bool               `bson:"display_cart_after_adding_product"`
	DisplayWishlistAfterAddingProduct           bool               `bson:"display_wishlist_after_adding_product"`
	MaximumShoppingCartItems                    int                `bson:"maximum_shopping_cart_items"`
	MaximumWishlistItems                        int                `bson:"maximum_wishlist_items"`
	AllowOutOfStockItemsToBeAddedToWishlist     bool               `bson:"allow_out_of_stock_items_to_be_added_to_wishlist"`
	MoveItemsFromWishlistToCart                 bool               `bson:"move_items_from_wishlist_to_cart"`
	CartsSharedBetweenStores                    bool               `bson:"carts_shared_between_stores"`
	ShowProductImagesOnShoppingCart             bool               `bson:"show_product_images_on_shopping_cart"`
	ShowProductImagesOnWishList                 bool               `bson:"show_product_images_on_wish_list"`
	ShowDiscountBox                             bool               `bson:"show_discount_box"`
	ShowGiftCardBox                             bool               `bson:"show_gift_card_box"`
	CrossSellsNumber                            int                `bson:"cross_sells_number"`
	EmailWishlistEnabled                        bool               `bson:"email_wishlist_enabled"`
	AllowAnonymousUsersToEmailWishlist          bool               `bson:"allow_anonymous_users_to_email_wishlist"`
	MiniShoppingCartEnabled                     bool               `bson:"mini_shopping_cart_enabled"`
	ShowProductImagesInMiniShoppingCart         bool               `bson:"show_product_images_in_mini_shopping_cart"`
	MiniShoppingCartProductNumber               int                `bson:"mini_shopping_cart_product_number"`
	RoundPricesDuringCalculation                bool               `bson:"round_prices_during_calculation"`
	GroupTierPricesForDistinctShoppingCartItems bool               `bson:"group_tier_prices_for_distinct_shopping_cart_items"`
	AllowCartItemEditing                        bool               `bson:"allow_cart_item_editing"`
	RenderAssociatedAttributeValueQuantity      bool               `bson:"render_associated_attribute_value_quantity"`
}

// ShoppingCartSettingsRepository defines the repository interface for ShoppingCartSettings
type ShoppingCartSettingsRepository interface {
	CreateMany(c context.Context, items []ShoppingCartSettings) error
	Create(c context.Context, shopping_cart_settings *ShoppingCartSettings) error
	Update(c context.Context, shopping_cart_settings *ShoppingCartSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShoppingCartSettings, error)
	FetchByID(c context.Context, ID string) (ShoppingCartSettings, error)
}

// ShoppingCartSettingsUsecase defines the use case interface for ShoppingCartSettings
type ShoppingCartSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (ShoppingCartSettings, error)
	Create(c context.Context, shopping_cart_settings *ShoppingCartSettings) error
	Update(c context.Context, shopping_cart_settings *ShoppingCartSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShoppingCartSettings, error)
}
