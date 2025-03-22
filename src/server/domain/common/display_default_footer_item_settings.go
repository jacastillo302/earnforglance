package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDisplayDefaultFooterItemSettings = "display_default_footer_item_settings"
)

// DisplayDefaultFooterItemSettings represents display default menu item settings
type DisplayDefaultFooterItemSettings struct {
	ID                                      primitive.ObjectID `bson:"_id,omitempty"`
	DisplaySitemapFooterItem                bool               `bson:"display_sitemap_footer_item"`
	DisplayContactUsFooterItem              bool               `bson:"display_contact_us_footer_item"`
	DisplayProductSearchFooterItem          bool               `bson:"display_product_search_footer_item"`
	DisplayNewsFooterItem                   bool               `bson:"display_news_footer_item"`
	DisplayBlogFooterItem                   bool               `bson:"display_blog_footer_item"`
	DisplayForumsFooterItem                 bool               `bson:"display_forums_footer_item"`
	DisplayRecentlyViewedProductsFooterItem bool               `bson:"display_recently_viewed_products_footer_item"`
	DisplayCompareProductsFooterItem        bool               `bson:"display_compare_products_footer_item"`
	DisplayNewProductsFooterItem            bool               `bson:"display_new_products_footer_item"`
	DisplayCustomerInfoFooterItem           bool               `bson:"display_customer_info_footer_item"`
	DisplayCustomerOrdersFooterItem         bool               `bson:"display_customer_orders_footer_item"`
	DisplayCustomerAddressesFooterItem      bool               `bson:"display_customer_addresses_footer_item"`
	DisplayShoppingCartFooterItem           bool               `bson:"display_shopping_cart_footer_item"`
	DisplayWishlistFooterItem               bool               `bson:"display_wishlist_footer_item"`
	DisplayApplyVendorAccountFooterItem     bool               `bson:"display_apply_vendor_account_footer_item"`
}

type DisplayDefaultFooterItemSettingsRepository interface {
	Create(c context.Context, display_default_footer_item_settings *DisplayDefaultFooterItemSettings) error
	Update(c context.Context, display_default_footer_item_settings *DisplayDefaultFooterItemSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DisplayDefaultFooterItemSettings, error)
	FetchByID(c context.Context, ID string) (DisplayDefaultFooterItemSettings, error)
}

type DisplayDefaultFooterItemSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (DisplayDefaultFooterItemSettings, error)
	Create(c context.Context, display_default_footer_item_settings *DisplayDefaultFooterItemSettings) error
	Update(c context.Context, display_default_footer_item_settings *DisplayDefaultFooterItemSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DisplayDefaultFooterItemSettings, error)
}
