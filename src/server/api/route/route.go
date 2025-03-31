package route

import (
	"time"

	"earnforglance/server/api/middleware"
	affiliate "earnforglance/server/api/route/affiliate"
	auth "earnforglance/server/api/route/api"
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
	install "earnforglance/server/api/route/install"
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
	"earnforglance/server/service/data/mongo"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

// RouterFunc defines the signature for router registration functions
type RouterFunc func(*bootstrap.Env, time.Duration, mongo.Database, *gin.RouterGroup)

// ModuleRouters defines a map of module routers grouped by domain
type ModuleRouters map[string][]RouterFunc

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})

	gin.Use(corsConfig)

	publicRouter := gin.Group("/api")
	// All Public APIs
	SignupRouter(env, timeout, db, publicRouter)
	LoginRouter(env, timeout, db, publicRouter)
	RefreshTokenRouter(env, timeout, db, publicRouter)
	auth.ApiClientRouter(env, timeout, db, publicRouter)
	install.InstallRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// Core APIs
	ProfileRouter(env, timeout, db, protectedRouter)
	TaskRouter(env, timeout, db, protectedRouter)

	// Register all domain-specific routers
	registerModuleRouters(env, timeout, db, protectedRouter)
}

func registerModuleRouters(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.RouterGroup) {
	// Define routers grouped by module for better organization and maintenance
	moduleRouters := ModuleRouters{

		"affiliate": {
			affiliate.AffiliateRouter,
		},
		"attributes": {
			attributes.BaseAttributeRouter,
			attributes.BaseAttributeValueRouter,
		},
		"blogs": {
			blogs.BlogCommentRouter,
			blogs.BlogPostRouter,
			blogs.BlogPostTagRouter,
			blogs.BlogSettingsRouter,
		},
		"catalog": {
			catalog.BackInStockSubscriptionRouter,
			catalog.CatalogSettingsRouter,
			catalog.CategoryRouter,
			catalog.CategoryTemplateRouter,
			catalog.CrossSellProductRouter,
			catalog.ManufacturerRouter,
			catalog.ManufacturerTemplateRouter,
			catalog.PredefinedProductAttributeValueRouter,
			catalog.ProductRouter,
			catalog.ProductAttributeRouter,
			catalog.ProductAttributeCombinationRouter,
			catalog.ProductAttributeCombinationPictureRouter,
			catalog.ProductAttributeMappingRouter,
			catalog.ProductAttributeValueRouter,
			catalog.ProductAttributeValuePictureRouter,
			catalog.ProductCategoryRouter,
			catalog.ProductEditorSettingsRouter,
			catalog.ProductManufacturerRouter,
			catalog.ProductPictureRouter,
			catalog.ProductProductTagMappingRouter,
			catalog.ProductReviewRouter,
			catalog.ProductReviewHelpfulnessRouter,
			catalog.ProductReviewReviewTypeMappingRouter,
			catalog.ProductSpecificationAttributeRouter,
			catalog.ProductTagRouter,
			catalog.ProductTemplateRouter,
			catalog.ProductVideoRouter,
			catalog.ProductWarehouseInventoryRouter,
			catalog.RelatedProductRouter,
			catalog.ReviewTypeRouter,
			catalog.SpecificationAttributeRouter,
			catalog.SpecificationAttributeGroupRouter,
			catalog.SpecificationAttributeOptionRouter,
			catalog.StockQuantityChangeRouter,
			catalog.TierPriceRouter,
		},
		"cms": {
			cms.WidgetSettingsRouter,
		},
		"common": {
			common.AddressRouter,
			common.AddressAttributeRouter,
			common.AddressAttributeValueRouter,
			common.AddressSettingsRouter,
			common.AdminAreaSettingsRouter,
			common.CommonSettingsRouter,
			common.DisplayDefaultFooterItemSettingsRouter,
			common.DisplayDefaultMenuItemSettingsRouter,
			common.GenericAttributeRouter,
			common.PdfSettingsRouter,
			common.SearchTermRouter,
			common.SearchTermReportLineRouter,
			common.SitemapSettingsRouter,
			common.SitemapXmlSettingsRouter,
		},
		"configuration": {
			configuration.SettingRouter,
		},
		"customers": {
			customers.BestCustomerReportLineRouter,
			customers.CustomerPasswordRouter,
			customers.CustomerRoleRouter,
			customers.CustomerSettingsRouter,
			customers.ExternalAuthenticationRecordRouter,
			customers.ExternalAuthenticationSettingsRouter,
			customers.MultiFactorAuthenticationSettingsRouter,
			customers.CustomerRouter,
			customers.CustomerAddressMappingRouter,
			customers.CustomerAttributeRouter,
			customers.CustomerAttributeValueRouter,
			customers.CustomerCustomerRoleMappingRouter,
			customers.RewardPointsHistoryRouter,
			customers.RewardPointsSettingsRouter,
		},
		"directory": {
			directory.CountryRouter,
			directory.CurrencyRouter,
			directory.CurrencySettingsRouter,
			directory.ExchangeRateRouter,
			directory.MeasureDimensionRouter,
			directory.MeasureSettingsRouter,
			directory.MeasureWeightRouter,
			directory.StateProvinceRouter,
		},
		"discounts": {
			discounts.DiscountRouter,
			discounts.DiscountCategoryMappingRouter,
			discounts.DiscountManufacturerMappingRouter,
			discounts.DiscountMappingRouter,
			discounts.DiscountProductMappingRouter,
			discounts.DiscountRequirementRouter,
			discounts.DiscountUsageHistoryRouter,
		},
		"forums": {
			forums.ForumRouter,
			forums.ForumGroupRouter,
			forums.ForumPostRouter,
			forums.ForumPostVoteRouter,
			forums.ForumSettingsRouter,
			forums.ForumSubscriptionRouter,
			forums.ForumTopicRouter,
			forums.PrivateMessageRouter,
		},
		"gdpr": {
			gdpr.CustomerPermanentlyDeletedRouter,
			gdpr.GdprSettingsRouter,
			gdpr.GdprConsentRouter,
			gdpr.GdprLogRouter,
		},
		"localization": {
			localization.LanguageRouter,
			localization.LocaleStringResourceRouter,
			localization.LocalizationSettingsRouter,
			localization.LocalizedPropertyRouter,
		},
		"logging": {
			logging.ActivityLogRouter,
			logging.ActivityLogTypeRouter,
			logging.LogRouter,
		},
		"media": {
			media.DownloadRouter,
			media.MediaSettingsRouter,
			media.PictureRouter,
			media.PictureBinaryRouter,
			media.PictureHashesRouter,
			media.VideoRouter,
		},
		"messages": {
			messages.CampaignRouter,
			messages.EmailAccountRouter,
			messages.EmailAccountSettingsRouter,
			messages.MessagesSettingsRouter,
			messages.MessageTemplateRouter,
			messages.MessageTemplatesSettingsRouter,
			messages.NewsLetterSubscriptionRouter,
			messages.QueuedEmailRouter,
		},
		"news": {
			news.NewsItemRouter,
			news.NewsCommentRouter,
			news.NewsSettingsRouter,
		},
		"orders": {
			orders.BestSellersReportLineRouter,
			orders.CheckoutAttributeRouter,
			orders.CheckoutAttributeValueRouter,
			orders.GiftCardRouter,
			orders.GiftCardUsageHistoryRouter,
			orders.OrderRouter,
			orders.OrderItemRouter,
			orders.OrderNoteRouter,
			orders.OrderSettingsRouter,
			orders.RecurringPaymentRouter,
			orders.RecurringPaymentHistoryRouter,
			orders.ReturnRequestRouter,
			orders.ReturnRequestActionRouter,
			orders.ReturnRequestReasonRouter,
			orders.SalesSummaryReportLineRouter,
			orders.ShoppingCartItemRouter,
			orders.ShoppingCartSettingsRouter,
		},
		"payments": {
			payments.PaymentSettingsRouter,
		},
		"polls": {
			polls.PollRouter,
			polls.PollAnswerRouter,
			polls.PollVotingRecordRouter,
		},
		"scheduleTasks": {
			scheduleTasks.ScheduleTaskRouter,
		},
		"security": {
			security.AclRecordRouter,
			security.CaptchaSettingsRouter,
			security.PermissionRecordRouter,
			security.PermissionRecordCustomerRoleMappingRouter,
			security.ProxySettingsRouter,
			security.RobotsTxtSettingsRouter,
			security.SecuritySettingsRouter,
		},
		"seo": {
			seo.UrlRecordRouter,
			seo.SeoSettingsRouter,
		},
		"shipping": {
			shipping.DeliveryDateRouter,
			shipping.PickupPointRouter,
			shipping.ProductAvailabilityRangeRouter,
			shipping.ShipmentRouter,
			shipping.ShipmentItemRouter,
			shipping.ShippingMethodRouter,
			shipping.ShippingMethodCountryMappingRouter,
			shipping.ShippingOptionRouter,
			shipping.ShippingSettingsRouter,
			shipping.WarehouseRouter,
		},
		"stores": {
			stores.StoreRouter,
			stores.StoreMappingRouter,
		},
		"tax": {
			tax.TaxCategoryRouter,
			tax.TaxSettingsRouter,
		},
		"topics": {
			topics.TopicRouter,
			topics.TopicTemplateRouter,
		},
		"vendors": {
			vendors.VendorRouter,
			vendors.VendorAttributeRouter,
			vendors.VendorAttributeValueRouter,
			vendors.VendorNoteRouter,
			vendors.VendorSettingsRouter,
		},
	}

	// Register all routers from all modules
	for _, routerFuncs := range moduleRouters {
		for _, routerFunc := range routerFuncs {
			routerFunc(env, timeout, db, router)
		}
	}
}
