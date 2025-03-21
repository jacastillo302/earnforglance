package route

import (
	"time"

	"earnforglance/server/api/middleware"
	affiliate "earnforglance/server/api/route/affiliate"
	attributes "earnforglance/server/api/route/attributes"
	blogs "earnforglance/server/api/route/blogs"
	catalog "earnforglance/server/api/route/catalog"
	cms "earnforglance/server/api/route/cms"
	common "earnforglance/server/api/route/common"
	configuration "earnforglance/server/api/route/configuration"
	customers "earnforglance/server/api/route/customers"
	directory "earnforglance/server/api/route/directory"
	discounts "earnforglance/server/api/route/discounts"
	forums "earnforglance/server/api/route/forums"
	gdpr "earnforglance/server/api/route/gdpr"
	localization "earnforglance/server/api/route/localization"
	logging "earnforglance/server/api/route/logging"
	media "earnforglance/server/api/route/media"
	messages "earnforglance/server/api/route/messages"
	news "earnforglance/server/api/route/news"
	orders "earnforglance/server/api/route/orders"
	payments "earnforglance/server/api/route/payments"
	polls "earnforglance/server/api/route/polls"
	scheduleTasks "earnforglance/server/api/route/scheduleTasks"
	security "earnforglance/server/api/route/security"
	seo "earnforglance/server/api/route/seo"
	shipping "earnforglance/server/api/route/shipping"
	stores "earnforglance/server/api/route/stores"
	tax "earnforglance/server/api/route/tax"
	topics "earnforglance/server/api/route/topics"
	vendors "earnforglance/server/api/route/vendors"

	"earnforglance/server/bootstrap"
	"earnforglance/server/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {

	publicRouter := gin.Group("/api")
	// All Public APIs
	SignupRouter(env, timeout, db, publicRouter)
	LoginRouter(env, timeout, db, publicRouter)
	RefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("/api")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	ProfileRouter(env, timeout, db, protectedRouter)
	TaskRouter(env, timeout, db, protectedRouter)

	affiliate.AffiliateRouter(env, timeout, db, protectedRouter)
	attributes.BaseAttributeRouter(env, timeout, db, protectedRouter)
	attributes.BaseAttributeValueRouter(env, timeout, db, protectedRouter)
	blogs.BlogCommentRouter(env, timeout, db, protectedRouter)
	blogs.BlogPostRouter(env, timeout, db, protectedRouter)
	blogs.BlogPostTagRouter(env, timeout, db, protectedRouter)
	blogs.BlogSettingsRouter(env, timeout, db, protectedRouter)
	catalog.BackInStockSubscriptionRouter(env, timeout, db, protectedRouter)
	catalog.CatalogSettingsRouter(env, timeout, db, protectedRouter)
	catalog.CategoryRouter(env, timeout, db, protectedRouter)
	catalog.CategoryTemplateRouter(env, timeout, db, protectedRouter)
	catalog.CrossSellProductRouter(env, timeout, db, protectedRouter)
	catalog.ManufacturerRouter(env, timeout, db, protectedRouter)
	catalog.ManufacturerTemplateRouter(env, timeout, db, protectedRouter)
	catalog.PredefinedProductAttributeValueRouter(env, timeout, db, protectedRouter)
	catalog.ProductRouter(env, timeout, db, protectedRouter)
	catalog.ProductAttributeRouter(env, timeout, db, protectedRouter)
	catalog.ProductAttributeCombinationRouter(env, timeout, db, protectedRouter)
	catalog.ProductAttributeCombinationPictureRouter(env, timeout, db, protectedRouter)
	catalog.ProductAttributeMappingRouter(env, timeout, db, protectedRouter)
	catalog.ProductAttributeValueRouter(env, timeout, db, protectedRouter)
	catalog.ProductAttributeValuePictureRouter(env, timeout, db, protectedRouter)
	catalog.ProductCategoryRouter(env, timeout, db, protectedRouter)
	catalog.ProductEditorSettingsRouter(env, timeout, db, protectedRouter)
	catalog.ProductManufacturerRouter(env, timeout, db, protectedRouter)
	catalog.ProductPictureRouter(env, timeout, db, protectedRouter)
	catalog.ProductProductTagMappingRouter(env, timeout, db, protectedRouter)
	catalog.ProductReviewRouter(env, timeout, db, protectedRouter)
	catalog.ProductReviewHelpfulnessRouter(env, timeout, db, protectedRouter)
	catalog.ProductReviewReviewTypeMappingRouter(env, timeout, db, protectedRouter)
	catalog.ProductSpecificationAttributeRouter(env, timeout, db, protectedRouter)
	catalog.ProductTagRouter(env, timeout, db, protectedRouter)
	catalog.ProductTemplateRouter(env, timeout, db, protectedRouter)
	catalog.ProductVideoRouter(env, timeout, db, protectedRouter)
	catalog.ProductWarehouseInventoryRouter(env, timeout, db, protectedRouter)
	catalog.RelatedProductRouter(env, timeout, db, protectedRouter)
	catalog.ReviewTypeRouter(env, timeout, db, protectedRouter)
	catalog.SpecificationAttributeRouter(env, timeout, db, protectedRouter)
	catalog.SpecificationAttributeGroupRouter(env, timeout, db, protectedRouter)
	catalog.SpecificationAttributeOptionRouter(env, timeout, db, protectedRouter)
	catalog.StockQuantityChangeRouter(env, timeout, db, protectedRouter)
	catalog.TierPriceRouter(env, timeout, db, protectedRouter)
	cms.WidgetSettingsRouter(env, timeout, db, protectedRouter)
	common.AddressRouter(env, timeout, db, protectedRouter)
	common.AddressAttributeRouter(env, timeout, db, protectedRouter)
	common.AddressAttributeValueRouter(env, timeout, db, protectedRouter)
	common.AddressSettingsRouter(env, timeout, db, protectedRouter)
	common.AdminAreaSettingsRouter(env, timeout, db, protectedRouter)
	common.CommonSettingsRouter(env, timeout, db, protectedRouter)
	common.DisplayDefaultFooterItemSettingsRouter(env, timeout, db, protectedRouter)
	common.DisplayDefaultMenuItemSettingsRouter(env, timeout, db, protectedRouter)
	common.GenericAttributeRouter(env, timeout, db, protectedRouter)
	common.PdfSettingsRouter(env, timeout, db, protectedRouter)
	common.SearchTermRouter(env, timeout, db, protectedRouter)
	common.SearchTermReportLineRouter(env, timeout, db, protectedRouter)
	common.SitemapSettingsRouter(env, timeout, db, protectedRouter)
	common.SitemapXmlSettingsRouter(env, timeout, db, protectedRouter)
	configuration.SettingRouter(env, timeout, db, protectedRouter)
	customers.BestCustomerReportLineRouter(env, timeout, db, protectedRouter)
	customers.CustomerRouter(env, timeout, db, protectedRouter)
	customers.CustomerAddressMappingRouter(env, timeout, db, protectedRouter)
	customers.CustomerAttributeRouter(env, timeout, db, protectedRouter)
	customers.CustomerAttributeValueRouter(env, timeout, db, protectedRouter)
	customers.CustomerCustomerRoleMappingRouter(env, timeout, db, protectedRouter)
	customers.RewardPointsHistoryRouter(env, timeout, db, protectedRouter)
	customers.RewardPointsSettingsRouter(env, timeout, db, protectedRouter)
	directory.CountryRouter(env, timeout, db, protectedRouter)
	directory.CurrencyRouter(env, timeout, db, protectedRouter)
	directory.CurrencySettingsRouter(env, timeout, db, protectedRouter)
	directory.ExchangeRateRouter(env, timeout, db, protectedRouter)
	directory.MeasureDimensionRouter(env, timeout, db, protectedRouter)
	directory.MeasureSettingsRouter(env, timeout, db, protectedRouter)
	directory.MeasureWeightRouter(env, timeout, db, protectedRouter)
	directory.StateProvinceRouter(env, timeout, db, protectedRouter)
	discounts.DiscountRouter(env, timeout, db, protectedRouter)
	discounts.DiscountCategoryMappingRouter(env, timeout, db, protectedRouter)
	discounts.DiscountManufacturerMappingRouter(env, timeout, db, protectedRouter)
	discounts.DiscountMappingRouter(env, timeout, db, protectedRouter)
	discounts.DiscountProductMappingRouter(env, timeout, db, protectedRouter)
	discounts.DiscountRequirementRouter(env, timeout, db, protectedRouter)
	discounts.DiscountUsageHistoryRouter(env, timeout, db, protectedRouter)
	forums.ForumRouter(env, timeout, db, protectedRouter)
	forums.ForumGroupRouter(env, timeout, db, protectedRouter)
	forums.ForumPostRouter(env, timeout, db, protectedRouter)
	forums.ForumPostVoteRouter(env, timeout, db, protectedRouter)
	forums.ForumSettingsRouter(env, timeout, db, protectedRouter)
	forums.ForumSubscriptionRouter(env, timeout, db, protectedRouter)
	forums.ForumTopicRouter(env, timeout, db, protectedRouter)
	forums.PrivateMessageRouter(env, timeout, db, protectedRouter)
	gdpr.CustomerPermanentlyDeletedRouter(env, timeout, db, protectedRouter)
	gdpr.GdprConsentRouter(env, timeout, db, protectedRouter)
	gdpr.GdprLogRouter(env, timeout, db, protectedRouter)
	gdpr.GdprSettingsRouter(env, timeout, db, protectedRouter)
	localization.LanguageRouter(env, timeout, db, protectedRouter)
	localization.LocaleStringResourceRouter(env, timeout, db, protectedRouter)
	localization.LocalizationSettingsRouter(env, timeout, db, protectedRouter)
	localization.LocalizedPropertyRouter(env, timeout, db, protectedRouter)
	logging.ActivityLogRouter(env, timeout, db, protectedRouter)
	logging.ActivityLogTypeRouter(env, timeout, db, protectedRouter)
	logging.LogRouter(env, timeout, db, protectedRouter)
	media.DownloadRouter(env, timeout, db, protectedRouter)
	media.MediaSettingsRouter(env, timeout, db, protectedRouter)
	media.PictureRouter(env, timeout, db, protectedRouter)
	media.PictureBinaryRouter(env, timeout, db, protectedRouter)
	media.PictureHashesRouter(env, timeout, db, protectedRouter)
	media.VideoRouter(env, timeout, db, protectedRouter)
	messages.CampaignRouter(env, timeout, db, protectedRouter)
	messages.EmailAccountRouter(env, timeout, db, protectedRouter)
	messages.EmailAccountSettingsRouter(env, timeout, db, protectedRouter)
	messages.MessagesSettingsRouter(env, timeout, db, protectedRouter)
	messages.MessageTemplateRouter(env, timeout, db, protectedRouter)
	messages.MessageTemplatesSettingsRouter(env, timeout, db, protectedRouter)
	messages.NewsLetterSubscriptionRouter(env, timeout, db, protectedRouter)
	messages.QueuedEmailRouter(env, timeout, db, protectedRouter)
	news.NewsCommentRouter(env, timeout, db, protectedRouter)
	news.NewsItemRouter(env, timeout, db, protectedRouter)
	news.NewsSettingsRouter(env, timeout, db, protectedRouter)
	orders.BestSellersReportLineRouter(env, timeout, db, protectedRouter)
	orders.CheckoutAttributeRouter(env, timeout, db, protectedRouter)
	orders.CheckoutAttributeValueRouter(env, timeout, db, protectedRouter)
	orders.GiftCardRouter(env, timeout, db, protectedRouter)
	orders.GiftCardUsageHistoryRouter(env, timeout, db, protectedRouter)
	orders.OrderRouter(env, timeout, db, protectedRouter)
	orders.OrderItemRouter(env, timeout, db, protectedRouter)
	orders.OrderNoteRouter(env, timeout, db, protectedRouter)
	orders.OrderSettingsRouter(env, timeout, db, protectedRouter)
	orders.RecurringPaymentRouter(env, timeout, db, protectedRouter)
	orders.RecurringPaymentHistoryRouter(env, timeout, db, protectedRouter)
	orders.ReturnRequestRouter(env, timeout, db, protectedRouter)
	orders.ReturnRequestActionRouter(env, timeout, db, protectedRouter)
	orders.ReturnRequestReasonRouter(env, timeout, db, protectedRouter)
	orders.SalesSummaryReportLineRouter(env, timeout, db, protectedRouter)
	orders.ShoppingCartItemRouter(env, timeout, db, protectedRouter)
	orders.ShoppingCartSettingsRouter(env, timeout, db, protectedRouter)
	payments.PaymentSettingsRouter(env, timeout, db, protectedRouter)
	polls.PollRouter(env, timeout, db, protectedRouter)
	polls.PollAnswerRouter(env, timeout, db, protectedRouter)
	polls.PollVotingRecordRouter(env, timeout, db, protectedRouter)
	scheduleTasks.ScheduleTaskRouter(env, timeout, db, protectedRouter)
	security.AclRecordRouter(env, timeout, db, protectedRouter)
	security.CaptchaSettingsRouter(env, timeout, db, protectedRouter)
	security.PermissionRecordRouter(env, timeout, db, protectedRouter)
	security.PermissionRecordCustomerRoleMappingRouter(env, timeout, db, protectedRouter)
	security.ProxySettingsRouter(env, timeout, db, protectedRouter)
	security.RobotsTxtSettingsRouter(env, timeout, db, protectedRouter)
	security.SecuritySettingsRouter(env, timeout, db, protectedRouter)
	seo.SeoSettingsRouter(env, timeout, db, protectedRouter)
	seo.UrlRecordRouter(env, timeout, db, protectedRouter)
	shipping.DeliveryDateRouter(env, timeout, db, protectedRouter)
	shipping.PickupPointRouter(env, timeout, db, protectedRouter)
	shipping.ProductAvailabilityRangeRouter(env, timeout, db, protectedRouter)
	shipping.ShipmentRouter(env, timeout, db, protectedRouter)
	shipping.ShipmentItemRouter(env, timeout, db, protectedRouter)
	shipping.ShippingMethodRouter(env, timeout, db, protectedRouter)
	shipping.ShippingMethodCountryMappingRouter(env, timeout, db, protectedRouter)
	shipping.ShippingOptionRouter(env, timeout, db, protectedRouter)
	shipping.ShippingSettingsRouter(env, timeout, db, protectedRouter)
	shipping.WarehouseRouter(env, timeout, db, protectedRouter)
	stores.StoreRouter(env, timeout, db, protectedRouter)
	stores.StoreMappingRouter(env, timeout, db, protectedRouter)
	tax.TaxCategoryRouter(env, timeout, db, protectedRouter)
	tax.TaxSettingsRouter(env, timeout, db, protectedRouter)
	topics.TopicRouter(env, timeout, db, protectedRouter)
	topics.TopicTemplateRouter(env, timeout, db, protectedRouter)
	vendors.VendorRouter(env, timeout, db, protectedRouter)
	vendors.VendorAttributeRouter(env, timeout, db, protectedRouter)
	vendors.VendorAttributeValueRouter(env, timeout, db, protectedRouter)
	vendors.VendorNoteRouter(env, timeout, db, protectedRouter)
	vendors.VendorSettingsRouter(env, timeout, db, protectedRouter)

}
