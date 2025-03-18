package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCatalogSettings = "catalog_settings"
)

// CatalogSettings represents catalog settings
type CatalogSettings struct {
	ID                                                 primitive.ObjectID                  `bson:"_id,omitempty"`
	AllowViewUnpublishedProductPage                    bool                                `bson:"allow_view_unpublished_product_page"`
	DisplayDiscontinuedMessageForUnpublishedProducts   bool                                `bson:"display_discontinued_message_for_unpublished_products"`
	PublishBackProductWhenCancellingOrders             bool                                `bson:"publish_back_product_when_cancelling_orders"`
	ShowSkuOnProductDetailsPage                        bool                                `bson:"show_sku_on_product_details_page"`
	ShowSkuOnCatalogPages                              bool                                `bson:"show_sku_on_catalog_pages"`
	ShowManufacturerPartNumber                         bool                                `bson:"show_manufacturer_part_number"`
	ShowGtin                                           bool                                `bson:"show_gtin"`
	ShowFreeShippingNotification                       bool                                `bson:"show_free_shipping_notification"`
	ShowShortDescriptionOnCatalogPages                 bool                                `bson:"show_short_description_on_catalog_pages"`
	AllowProductSorting                                bool                                `bson:"allow_product_sorting"`
	AllowProductViewModeChanging                       bool                                `bson:"allow_product_view_mode_changing"`
	DefaultViewMode                                    string                              `bson:"default_view_mode"`
	ShowProductsFromSubcategories                      bool                                `bson:"show_products_from_subcategories"`
	ShowCategoryProductNumber                          bool                                `bson:"show_category_product_number"`
	ShowCategoryProductNumberIncludingSubcategories    bool                                `bson:"show_category_product_number_including_subcategories"`
	CategoryBreadcrumbEnabled                          bool                                `bson:"category_breadcrumb_enabled"`
	ShowShareButton                                    bool                                `bson:"show_share_button"`
	PageShareCode                                      string                              `bson:"page_share_code"`
	ProductReviewsMustBeApproved                       bool                                `bson:"product_reviews_must_be_approved"`
	OneReviewPerProductFromCustomer                    bool                                `bson:"one_review_per_product_from_customer"`
	DefaultProductRatingValue                          int                                 `bson:"default_product_rating_value"`
	AllowAnonymousUsersToReviewProduct                 bool                                `bson:"allow_anonymous_users_to_review_product"`
	ProductReviewPossibleOnlyAfterPurchasing           bool                                `bson:"product_review_possible_only_after_purchasing"`
	NotifyStoreOwnerAboutNewProductReviews             bool                                `bson:"notify_store_owner_about_new_product_reviews"`
	NotifyCustomerAboutProductReviewReply              bool                                `bson:"notify_customer_about_product_review_reply"`
	ShowProductReviewsPerStore                         bool                                `bson:"show_product_reviews_per_store"`
	ShowProductReviewsTabOnAccountPage                 bool                                `bson:"show_product_reviews_tab_on_account_page"`
	ProductReviewsPageSizeOnAccountPage                int                                 `bson:"product_reviews_page_size_on_account_page"`
	ProductReviewsSortByCreatedDateAscending           bool                                `bson:"product_reviews_sort_by_created_date_ascending"`
	EmailAFriendEnabled                                bool                                `bson:"email_a_friend_enabled"`
	AllowAnonymousUsersToEmailAFriend                  bool                                `bson:"allow_anonymous_users_to_email_a_friend"`
	RecentlyViewedProductsNumber                       int                                 `bson:"recently_viewed_products_number"`
	RecentlyViewedProductsEnabled                      bool                                `bson:"recently_viewed_products_enabled"`
	NewProductsEnabled                                 bool                                `bson:"new_products_enabled"`
	NewProductsPageSize                                int                                 `bson:"new_products_page_size"`
	NewProductsAllowCustomersToSelectPageSize          bool                                `bson:"new_products_allow_customers_to_select_page_size"`
	NewProductsPageSizeOptions                         string                              `bson:"new_products_page_size_options"`
	CompareProductsEnabled                             bool                                `bson:"compare_products_enabled"`
	CompareProductsNumber                              int                                 `bson:"compare_products_number"`
	ProductSearchAutoCompleteEnabled                   bool                                `bson:"product_search_auto_complete_enabled"`
	ProductSearchEnabled                               bool                                `bson:"product_search_enabled"`
	ProductSearchAutoCompleteNumberOfProducts          int                                 `bson:"product_search_auto_complete_number_of_products"`
	ShowProductImagesInSearchAutoComplete              bool                                `bson:"show_product_images_in_search_auto_complete"`
	ShowLinkToAllResultInSearchAutoComplete            bool                                `bson:"show_link_to_all_result_in_search_auto_complete"`
	ProductSearchTermMinimumLength                     int                                 `bson:"product_search_term_minimum_length"`
	ShowBestsellersOnHomepage                          bool                                `bson:"show_bestsellers_on_homepage"`
	NumberOfBestsellersOnHomepage                      int                                 `bson:"number_of_bestsellers_on_homepage"`
	ShowSearchBoxCategories                            bool                                `bson:"show_search_box_categories"`
	SearchPageProductsPerPage                          int                                 `bson:"search_page_products_per_page"`
	SearchPageAllowCustomersToSelectPageSize           bool                                `bson:"search_page_allow_customers_to_select_page_size"`
	SearchPagePageSizeOptions                          string                              `bson:"search_page_page_size_options"`
	SearchPagePriceRangeFiltering                      bool                                `bson:"search_page_price_range_filtering"`
	SearchPagePriceFrom                                float64                             `bson:"search_page_price_from"`
	SearchPagePriceTo                                  float64                             `bson:"search_page_price_to"`
	SearchPageManuallyPriceRange                       bool                                `bson:"search_page_manually_price_range"`
	ProductsAlsoPurchasedEnabled                       bool                                `bson:"products_also_purchased_enabled"`
	ProductsAlsoPurchasedNumber                        int                                 `bson:"products_also_purchased_number"`
	AjaxProcessAttributeChange                         bool                                `bson:"ajax_process_attribute_change"`
	NumberOfProductTags                                int                                 `bson:"number_of_product_tags"`
	ProductsByTagPageSize                              int                                 `bson:"products_by_tag_page_size"`
	ProductsByTagAllowCustomersToSelectPageSize        bool                                `bson:"products_by_tag_allow_customers_to_select_page_size"`
	ProductsByTagPageSizeOptions                       string                              `bson:"products_by_tag_page_size_options"`
	ProductsByTagPriceRangeFiltering                   bool                                `bson:"products_by_tag_price_range_filtering"`
	ProductsByTagPriceFrom                             float64                             `bson:"products_by_tag_price_from"`
	ProductsByTagPriceTo                               float64                             `bson:"products_by_tag_price_to"`
	ProductsByTagManuallyPriceRange                    bool                                `bson:"products_by_tag_manually_price_range"`
	IncludeShortDescriptionInCompareProducts           bool                                `bson:"include_short_description_in_compare_products"`
	IncludeFullDescriptionInCompareProducts            bool                                `bson:"include_full_description_in_compare_products"`
	IncludeFeaturedProductsInNormalLists               bool                                `bson:"include_featured_products_in_normal_lists"`
	UseLinksInRequiredProductWarnings                  bool                                `bson:"use_links_in_required_product_warnings"`
	DisplayTierPricesWithDiscounts                     bool                                `bson:"display_tier_prices_with_discounts"`
	IgnoreDiscounts                                    bool                                `bson:"ignore_discounts"`
	IgnoreFeaturedProducts                             bool                                `bson:"ignore_featured_products"`
	IgnoreAcl                                          bool                                `bson:"ignore_acl"`
	IgnoreStoreLimitations                             bool                                `bson:"ignore_store_limitations"`
	CacheProductPrices                                 bool                                `bson:"cache_product_prices"`
	MaximumBackInStockSubscriptions                    int                                 `bson:"maximum_back_in_stock_subscriptions"`
	ManufacturersBlockItemsToDisplay                   int                                 `bson:"manufacturers_block_items_to_display"`
	DisplayTaxShippingInfoFooter                       bool                                `bson:"display_tax_shipping_info_footer"`
	DisplayTaxShippingInfoProductDetailsPage           bool                                `bson:"display_tax_shipping_info_product_details_page"`
	DisplayTaxShippingInfoProductBoxes                 bool                                `bson:"display_tax_shipping_info_product_boxes"`
	DisplayTaxShippingInfoShoppingCart                 bool                                `bson:"display_tax_shipping_info_shopping_cart"`
	DisplayTaxShippingInfoWishlist                     bool                                `bson:"display_tax_shipping_info_wishlist"`
	DisplayTaxShippingInfoOrderDetailsPage             bool                                `bson:"display_tax_shipping_info_order_details_page"`
	DefaultCategoryPageSizeOptions                     string                              `bson:"default_category_page_size_options"`
	DefaultCategoryPageSize                            int                                 `bson:"default_category_page_size"`
	DefaultManufacturerPageSizeOptions                 string                              `bson:"default_manufacturer_page_size_options"`
	DefaultManufacturerPageSize                        int                                 `bson:"default_manufacturer_page_size"`
	ProductSortingEnumDisabled                         []int                               `bson:"product_sorting_enum_disabled"`
	ProductSortingEnumDisplayOrder                     map[int]int                         `bson:"product_sorting_enum_display_order"`
	ExportImportProductAttributes                      bool                                `bson:"export_import_product_attributes"`
	ExportImportProductUseLimitedToStores              bool                                `bson:"export_import_product_use_limited_to_stores"`
	ExportImportCategoryUseLimitedToStores             bool                                `bson:"export_import_category_use_limited_to_stores"`
	ExportImportProductSpecificationAttributes         bool                                `bson:"export_import_product_specification_attributes"`
	ExportImportTierPrices                             bool                                `bson:"export_import_tier_prices"`
	ExportImportUseDropdownlistsForAssociatedEntities  bool                                `bson:"export_import_use_dropdownlists_for_associated_entities"`
	ExportImportProductCategoryBreadcrumb              bool                                `bson:"export_import_product_category_breadcrumb"`
	ExportImportCategoriesUsingCategoryName            bool                                `bson:"export_import_categories_using_category_name"`
	ExportImportAllowDownloadImages                    bool                                `bson:"export_import_allow_download_images"`
	ExportImportSplitProductsFile                      bool                                `bson:"export_import_split_products_file"`
	ExportImportProductsCountInOneFile                 int                                 `bson:"export_import_products_count_in_one_file"`
	RemoveRequiredProducts                             bool                                `bson:"remove_required_products"`
	ExportImportRelatedEntitiesByName                  bool                                `bson:"export_import_related_entities_by_name"`
	CountDisplayedYearsDatePicker                      int                                 `bson:"count_displayed_years_date_picker"`
	DisplayDatePreOrderAvailability                    bool                                `bson:"display_date_pre_order_availability"`
	UseAjaxLoadMenu                                    bool                                `bson:"use_ajax_load_menu"`
	UseAjaxCatalogProductsLoading                      bool                                `bson:"use_ajax_catalog_products_loading"`
	EnableManufacturerFiltering                        bool                                `bson:"enable_manufacturer_filtering"`
	EnablePriceRangeFiltering                          bool                                `bson:"enable_price_range_filtering"`
	EnableSpecificationAttributeFiltering              bool                                `bson:"enable_specification_attribute_filtering"`
	DisplayFromPrices                                  bool                                `bson:"display_from_prices"`
	AttributeValueOutOfStockDisplayType                AttributeValueOutOfStockDisplayType `bson:"attribute_value_out_of_stock_display_type"`
	AllowCustomersToSearchWithManufacturerName         bool                                `bson:"allow_customers_to_search_with_manufacturer_name"`
	AllowCustomersToSearchWithCategoryName             bool                                `bson:"allow_customers_to_search_with_category_name"`
	DisplayAllPicturesOnCatalogPages                   bool                                `bson:"display_all_pictures_on_catalog_pages"`
	ProductUrlStructureTypeId                          int                                 `bson:"product_url_structure_type_id"`
	ActiveSearchProviderSystemName                     string                              `bson:"active_search_provider_system_name"`
	UseStandardSearchWhenSearchProviderThrowsException bool                                `bson:"use_standard_search_when_search_provider_throws_exception"`
	VendorProductReviewsPageSize                       int                                 `bson:"vendor_product_reviews_page_size"`
}

type CatalogSettingsRepository interface {
	Create(c context.Context, catalog_settings *CatalogSettings) error
	Update(c context.Context, catalog_settings *CatalogSettings) error
	Delete(c context.Context, catalog_settings *CatalogSettings) error
	Fetch(c context.Context) ([]CatalogSettings, error)
	FetchByID(c context.Context, catalog_settingsID string) (CatalogSettings, error)
}

type CatalogSettingsUsecase interface {
	FetchByID(c context.Context, catalog_settingsID string) (CatalogSettings, error)
	Create(c context.Context, catalog_settings *CatalogSettings) error
	Update(c context.Context, catalog_settings *CatalogSettings) error
	Delete(c context.Context, catalog_settings *CatalogSettings) error
	Fetch(c context.Context) ([]CatalogSettings, error)
}
