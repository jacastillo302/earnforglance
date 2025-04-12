package domain

import (
	"context"
	"time"

	affiliates "earnforglance/server/domain/affiliate"
	blogs "earnforglance/server/domain/blogs"
	catalog "earnforglance/server/domain/catalog"
	commons "earnforglance/server/domain/common"
	settings "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	directory "earnforglance/server/domain/directory"
	discounts "earnforglance/server/domain/discounts"
	forums "earnforglance/server/domain/forums"
	gdprs "earnforglance/server/domain/gdpr"
	lang "earnforglance/server/domain/localization"
	loggings "earnforglance/server/domain/logging"
	media "earnforglance/server/domain/media"
	messages "earnforglance/server/domain/messages"
	news "earnforglance/server/domain/news"
	orders "earnforglance/server/domain/orders"
	polls "earnforglance/server/domain/polls"
	tasks "earnforglance/server/domain/scheduleTasks"
	security "earnforglance/server/domain/security"
	shippings "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
	topics "earnforglance/server/domain/topics"
	vendors "earnforglance/server/domain/vendors"
)

type Install struct {
	Status       bool
	Details      string
	CreatedOnUtc time.Time
}

type InstallRepository interface {
	PingDatabase(c context.Context) error
	InstallPermissionRecord(c context.Context, items []security.PermissionRecord) error
	InstallCurrencies(c context.Context, items []directory.Currency) error
	InstallMeasureDimension(c context.Context, items []directory.MeasureDimension) error
	InstallMeasureWeight(c context.Context, items []directory.MeasureWeight) error
	InstallTaxCategory(c context.Context, items []taxes.TaxCategory) error
	InstallLanguages(c context.Context, items []lang.Language) error
	InstallLocaleStringResource(c context.Context, items []lang.LocaleStringResource) error
	InstallStores(c context.Context, items []stores.Store) error
	InstallSettings(c context.Context, items []settings.Setting) error
	InstallCountries(c context.Context, items []directory.Country) error
	InstallStateProvince(c context.Context, items []directory.StateProvince) error
	InstallShippingMethod(c context.Context, items []shippings.ShippingMethod) error
	InstallDeliveryDate(c context.Context, items []shippings.DeliveryDate) error
	InstallProductAvailabilityRange(c context.Context, items []shippings.ProductAvailabilityRange) error
	InstallEmailAccount(c context.Context, items []messages.EmailAccount) error
	InstallMessageTemplate(c context.Context, items []messages.MessageTemplate) error
	InstallTopicTemplate(c context.Context, items []topics.TopicTemplate) error
	InstallCustomerRole(c context.Context, items []customers.CustomerRole) error
	InstallCustomer(c context.Context, items []customers.Customer) error
	InstallCustomerPassword(c context.Context, items []customers.CustomerPassword) error
	InstallAddress(c context.Context, items []commons.Address) error
	InstallCustomerAddressMapping(c context.Context, item customers.CustomerAddressMapping) error
	UpdateCustomer(c context.Context, items customers.Customer) error
	InstallCustomerCustomerRoleMapping(c context.Context, items []customers.CustomerCustomerRoleMapping) error
	InstallTopic(c context.Context, items []topics.Topic) error
	InstallActivityLogType(c context.Context, items []loggings.ActivityLogType) error
	InstallActivityLog(c context.Context, items []loggings.ActivityLog) error
	InstallProductTemplate(c context.Context, items []catalog.ProductTemplate) error
	InstallCategoryTemplate(c context.Context, items []catalog.CategoryTemplate) error
	InstallManufacturerTemplate(c context.Context, items []catalog.ManufacturerTemplate) error
	InstallScheduleTask(c context.Context, items []tasks.ScheduleTask) error
	InstallReturnRequestReason(c context.Context, items []orders.ReturnRequestReason) error
	InstallReturnRequestAction(c context.Context, items []orders.ReturnRequestAction) error
	InstallCheckoutAttribute(c context.Context, items []orders.CheckoutAttribute) error
	InstallCheckoutAttributeValue(c context.Context, items []orders.CheckoutAttributeValue) error
	InstallSpecificationAttributeGroup(c context.Context, items []catalog.SpecificationAttributeGroup) error
	InstallSpecificationAttributeOption(c context.Context, items []catalog.SpecificationAttributeOption) error
	InstallSpecificationAttribute(c context.Context, items []catalog.SpecificationAttribute) error
	InstallProductAttribute(c context.Context, items []catalog.ProductAttribute) error
	InstallPicture(c context.Context, items []media.Picture) error
	InstallCategory(c context.Context, items []catalog.Category) error
	InstallManufacturer(c context.Context, items []catalog.Manufacturer) error
	InstallWarehouse(c context.Context, items []shippings.Warehouse) error
	InstallVendor(c context.Context, items []vendors.Vendor) error
	InstallAffiliate(c context.Context, items []affiliates.Affiliate) error
	InstallForumGroup(c context.Context, items []forums.ForumGroup) error
	InstallForum(c context.Context, items []forums.Forum) error
	InstallDiscount(c context.Context, items []discounts.Discount) error
	InstallBlogPost(c context.Context, items []blogs.BlogPost) error
	InstallBlogComment(c context.Context, items []blogs.BlogComment) error
	InstallPoll(c context.Context, items []polls.Poll) error
	InstallPollAnswer(c context.Context, items []polls.PollAnswer) error
	InstallNewsItem(c context.Context, items []news.NewsItem) error
	InstallNewsComment(c context.Context, items []news.NewsComment) error
	InstallSearchTerm(c context.Context, items []commons.SearchTerm) error
	InstallProduct(c context.Context, items []catalog.Product) error
	InstallProductSpecificationAttribute(c context.Context, items []catalog.ProductSpecificationAttribute) error
	InstallProductTag(c context.Context, items []catalog.ProductTag) error
	InstallProductProductTagMapping(c context.Context, items []catalog.ProductProductTagMapping) error
	InstallProductAttributeValue(c context.Context, items []catalog.ProductAttributeValue) error
	InstallProductAttributeMapping(c context.Context, items []catalog.ProductAttributeMapping) error
	InstallProductPicture(c context.Context, items []catalog.ProductPicture) error
	InstallProductCategory(c context.Context, items []catalog.ProductCategory) error
	InstallTierPrice(c context.Context, items []catalog.TierPrice) error
	InstallProductManufacturer(c context.Context, items []catalog.ProductManufacturer) error
	InstallProductAttributeValuePicture(c context.Context, items []catalog.ProductAttributeValuePicture) error
	InstallDownload(c context.Context, items []media.Download) error
	InstallRelatedProduct(c context.Context, items []catalog.RelatedProduct) error
	InstallProductReview(c context.Context, items []catalog.ProductReview) error
	InstallStockQuantityChange(c context.Context, items []catalog.StockQuantityChange) error
	InstallGdprConsent(c context.Context, items []gdprs.GdprConsent) error
	InstallOrder(c context.Context, items []orders.Order) error
	InstallOrderItem(c context.Context, items []orders.OrderItem) error
	InstallShipment(c context.Context, items []shippings.Shipment) error
	InstallShipmentItem(c context.Context, items []shippings.ShipmentItem) error
	InstallOrderNote(c context.Context, items []orders.OrderNote) error
}

// GdprLogUsecase interface
type InstallLogUsecase interface {
	PingDatabase(c context.Context) error
	InstallPermissionRecord(c context.Context, items []security.PermissionRecord) error
	InstallCurrencies(c context.Context, items []directory.Currency) error
	InstallMeasureDimension(c context.Context, items []directory.MeasureDimension) error
	InstallMeasureWeight(c context.Context, items []directory.MeasureWeight) error
	InstallTaxCategory(c context.Context, items []taxes.TaxCategory) error
	InstallLanguages(c context.Context, items []lang.Language) error
	InstallLocaleStringResource(c context.Context, items []lang.LocaleStringResource) error
	InstallStores(c context.Context, items []stores.Store) error
	InstallSettings(c context.Context, items []settings.Setting) error
	InstallCountries(c context.Context, items []directory.Country) error
	InstallStateProvince(c context.Context, items []directory.StateProvince) error
	InstallShippingMethod(c context.Context, items []shippings.ShippingMethod) error
	InstallDeliveryDate(c context.Context, items []shippings.DeliveryDate) error
	InstallProductAvailabilityRange(c context.Context, items []shippings.ProductAvailabilityRange) error
	InstallEmailAccount(c context.Context, items []messages.EmailAccount) error
	InstallMessageTemplate(c context.Context, items []messages.MessageTemplate) error
	InstallTopicTemplate(c context.Context, items []topics.TopicTemplate) error
	InstallCustomerRole(c context.Context, items []customers.CustomerRole) error
	InstallCustomer(c context.Context, items []customers.Customer) error
	InstallCustomerPassword(c context.Context, items []customers.CustomerPassword) error
	InstallAddress(c context.Context, items []commons.Address) error
	InstallCustomerAddressMapping(c context.Context, item customers.CustomerAddressMapping) error
	UpdateCustomer(c context.Context, items customers.Customer) error
	InstallCustomerCustomerRoleMapping(c context.Context, items []customers.CustomerCustomerRoleMapping) error
	InstallTopic(c context.Context, items []topics.Topic) error
	InstallActivityLogType(c context.Context, items []loggings.ActivityLogType) error
	InstallActivityLog(c context.Context, items []loggings.ActivityLog) error
	InstallProductTemplate(c context.Context, items []catalog.ProductTemplate) error
	InstallCategoryTemplate(c context.Context, items []catalog.CategoryTemplate) error
	InstallManufacturerTemplate(c context.Context, items []catalog.ManufacturerTemplate) error
	InstallScheduleTask(c context.Context, items []tasks.ScheduleTask) error
	InstallReturnRequestReason(c context.Context, items []orders.ReturnRequestReason) error
	InstallReturnRequestAction(c context.Context, items []orders.ReturnRequestAction) error
	InstallCheckoutAttribute(c context.Context, items []orders.CheckoutAttribute) error
	InstallCheckoutAttributeValue(c context.Context, items []orders.CheckoutAttributeValue) error
	InstallSpecificationAttributeGroup(c context.Context, items []catalog.SpecificationAttributeGroup) error
	InstallSpecificationAttributeOption(c context.Context, items []catalog.SpecificationAttributeOption) error
	InstallSpecificationAttribute(c context.Context, items []catalog.SpecificationAttribute) error
	InstallProductAttribute(c context.Context, items []catalog.ProductAttribute) error
	InstallPicture(c context.Context, items []media.Picture) error
	InstallCategory(c context.Context, items []catalog.Category) error
	InstallManufacturer(c context.Context, items []catalog.Manufacturer) error
	InstallWarehouse(c context.Context, items []shippings.Warehouse) error
	InstallVendor(c context.Context, items []vendors.Vendor) error
	InstallAffiliate(c context.Context, items []affiliates.Affiliate) error
	InstallForumGroup(c context.Context, items []forums.ForumGroup) error
	InstallForum(c context.Context, items []forums.Forum) error
	InstallDiscount(c context.Context, items []discounts.Discount) error
	InstallBlogPost(c context.Context, items []blogs.BlogPost) error
	InstallBlogComment(c context.Context, items []blogs.BlogComment) error
	InstallPoll(c context.Context, items []polls.Poll) error
	InstallPollAnswer(c context.Context, items []polls.PollAnswer) error
	InstallNewsItem(c context.Context, items []news.NewsItem) error
	InstallNewsComment(c context.Context, items []news.NewsComment) error
	InstallSearchTerm(c context.Context, items []commons.SearchTerm) error
	InstallProduct(c context.Context, items []catalog.Product) error
	InstallProductSpecificationAttribute(c context.Context, items []catalog.ProductSpecificationAttribute) error
	InstallProductTag(c context.Context, items []catalog.ProductTag) error
	InstallProductProductTagMapping(c context.Context, items []catalog.ProductProductTagMapping) error
	InstallProductAttributeValue(c context.Context, items []catalog.ProductAttributeValue) error
	InstallProductAttributeMapping(c context.Context, items []catalog.ProductAttributeMapping) error
	InstallProductPicture(c context.Context, items []catalog.ProductPicture) error
	InstallProductCategory(c context.Context, items []catalog.ProductCategory) error
	InstallTierPrice(c context.Context, items []catalog.TierPrice) error
	InstallProductManufacturer(c context.Context, items []catalog.ProductManufacturer) error
	InstallProductAttributeValuePicture(c context.Context, items []catalog.ProductAttributeValuePicture) error
	InstallDownload(c context.Context, items []media.Download) error
	InstallRelatedProduct(c context.Context, items []catalog.RelatedProduct) error
	InstallProductReview(c context.Context, items []catalog.ProductReview) error
	InstallStockQuantityChange(c context.Context, items []catalog.StockQuantityChange) error
	InstallGdprConsent(c context.Context, items []gdprs.GdprConsent) error
	InstallOrder(c context.Context, items []orders.Order) error
	InstallOrderItem(c context.Context, items []orders.OrderItem) error
	InstallShipment(c context.Context, items []shippings.Shipment) error
	InstallShipmentItem(c context.Context, items []shippings.ShipmentItem) error
	InstallOrderNote(c context.Context, items []orders.OrderNote) error
}
