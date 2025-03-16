package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionDefaultMenuItemSettings = "display_default_menu_item_settings"
)

// DisplayDefaultMenuItemSettings represents display default menu item settings
type DisplayDefaultMenuItemSettings struct {
	ID                           primitive.ObjectID `bson:"_id,omitempty"`
	DisplayHomepageMenuItem      bool               `bson:"display_homepage_menu_item"`
	DisplayNewProductsMenuItem   bool               `bson:"display_new_products_menu_item"`
	DisplayProductSearchMenuItem bool               `bson:"display_product_search_menu_item"`
	DisplayCustomerInfoMenuItem  bool               `bson:"display_customer_info_menu_item"`
	DisplayBlogMenuItem          bool               `bson:"display_blog_menu_item"`
	DisplayForumsMenuItem        bool               `bson:"display_forums_menu_item"`
	DisplayContactUsMenuItem     bool               `bson:"display_contact_us_menu_item"`
}
