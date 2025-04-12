package install

import (
	"context"
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
	install "earnforglance/server/domain/install"
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

	"time"
)

type InstallUsecase struct {
	InstallRepository install.InstallRepository
	contextTimeout    time.Duration
}

func NewInstallUsecase(IsntallRepository install.InstallRepository, timeout time.Duration) install.InstallLogUsecase {
	return &InstallUsecase{
		InstallRepository: IsntallRepository,
		contextTimeout:    timeout,
	}
}

func (tu *InstallUsecase) PingDatabase(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.PingDatabase(ctx)
}

func (tu *InstallUsecase) InstallPermissionRecord(c context.Context, items []security.PermissionRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallPermissionRecord(ctx, items)
}

func (tu *InstallUsecase) InstallCurrencies(c context.Context, items []directory.Currency) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCurrencies(ctx, items)
}

func (tu *InstallUsecase) InstallStores(c context.Context, items []stores.Store) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallStores(ctx, items)
}

func (tu *InstallUsecase) InstallSettings(c context.Context, items []settings.Setting) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallSettings(ctx, items)
}

func (tu *InstallUsecase) InstallMeasureDimension(c context.Context, items []directory.MeasureDimension) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallMeasureDimension(ctx, items)
}

func (tu *InstallUsecase) InstallMeasureWeight(c context.Context, items []directory.MeasureWeight) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallMeasureWeight(ctx, items)
}

func (tu *InstallUsecase) InstallTaxCategory(c context.Context, items []taxes.TaxCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallTaxCategory(ctx, items)
}

func (tu *InstallUsecase) InstallLanguages(c context.Context, items []lang.Language) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallLanguages(ctx, items)
}

func (tu *InstallUsecase) InstallLocaleStringResource(c context.Context, items []lang.LocaleStringResource) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallLocaleStringResource(ctx, items)
}

func (tu *InstallUsecase) InstallCountries(c context.Context, items []directory.Country) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCountries(ctx, items)
}

func (tu *InstallUsecase) InstallStateProvince(c context.Context, items []directory.StateProvince) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallStateProvince(ctx, items)
}

func (tu *InstallUsecase) InstallShippingMethod(c context.Context, items []shippings.ShippingMethod) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallShippingMethod(ctx, items)
}

func (tu *InstallUsecase) InstallDeliveryDate(c context.Context, items []shippings.DeliveryDate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallDeliveryDate(ctx, items)
}

func (tu *InstallUsecase) InstallProductAvailabilityRange(c context.Context, items []shippings.ProductAvailabilityRange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductAvailabilityRange(ctx, items)
}

func (tu *InstallUsecase) InstallEmailAccount(c context.Context, items []messages.EmailAccount) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallEmailAccount(ctx, items)
}

func (tu *InstallUsecase) InstallMessageTemplate(c context.Context, items []messages.MessageTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallMessageTemplate(ctx, items)
}

func (tu *InstallUsecase) InstallTopicTemplate(c context.Context, items []topics.TopicTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallTopicTemplate(ctx, items)
}

func (tu *InstallUsecase) InstallCustomerRole(c context.Context, items []customers.CustomerRole) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCustomerRole(ctx, items)
}

func (tu *InstallUsecase) InstallCustomer(c context.Context, items []customers.Customer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCustomer(ctx, items)
}

func (tu *InstallUsecase) InstallCustomerPassword(c context.Context, items []customers.CustomerPassword) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCustomerPassword(ctx, items)
}

func (tu *InstallUsecase) InstallAddress(c context.Context, items []commons.Address) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallAddress(ctx, items)
}

func (tu *InstallUsecase) UpdateCustomer(c context.Context, items customers.Customer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.UpdateCustomer(ctx, items)
}

func (tu *InstallUsecase) InstallCustomerAddressMapping(c context.Context, items customers.CustomerAddressMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCustomerAddressMapping(ctx, items)
}

func (tu *InstallUsecase) InstallCustomerCustomerRoleMapping(c context.Context, items []customers.CustomerCustomerRoleMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCustomerCustomerRoleMapping(ctx, items)
}

func (tu *InstallUsecase) InstallTopic(c context.Context, items []topics.Topic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallTopic(ctx, items)
}

func (tu *InstallUsecase) InstallActivityLogType(c context.Context, items []loggings.ActivityLogType) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallActivityLogType(ctx, items)
}

func (tu *InstallUsecase) InstallProductTemplate(c context.Context, items []catalog.ProductTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductTemplate(ctx, items)
}

func (tu *InstallUsecase) InstallCategoryTemplate(c context.Context, items []catalog.CategoryTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCategoryTemplate(ctx, items)

}

func (tu *InstallUsecase) InstallManufacturerTemplate(c context.Context, items []catalog.ManufacturerTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallManufacturerTemplate(ctx, items)
}

func (tu *InstallUsecase) InstallScheduleTask(c context.Context, items []tasks.ScheduleTask) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallScheduleTask(ctx, items)

}

func (tu *InstallUsecase) InstallReturnRequestReason(c context.Context, items []orders.ReturnRequestReason) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallReturnRequestReason(ctx, items)

}

func (tu *InstallUsecase) InstallReturnRequestAction(c context.Context, items []orders.ReturnRequestAction) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallReturnRequestAction(ctx, items)
}

func (tu *InstallUsecase) InstallCheckoutAttribute(c context.Context, items []orders.CheckoutAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCheckoutAttribute(ctx, items)
}

func (tu *InstallUsecase) InstallCheckoutAttributeValue(c context.Context, items []orders.CheckoutAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCheckoutAttributeValue(ctx, items)
}

func (tu *InstallUsecase) InstallSpecificationAttribute(c context.Context, items []catalog.SpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallSpecificationAttribute(ctx, items)
}

func (tu *InstallUsecase) InstallSpecificationAttributeOption(c context.Context, items []catalog.SpecificationAttributeOption) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallSpecificationAttributeOption(ctx, items)
}

func (tu *InstallUsecase) InstallSpecificationAttributeGroup(c context.Context, items []catalog.SpecificationAttributeGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallSpecificationAttributeGroup(ctx, items)
}

func (tu *InstallUsecase) InstallProductAttribute(c context.Context, items []catalog.ProductAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductAttribute(ctx, items)
}

func (tu *InstallUsecase) InstallCategory(c context.Context, items []catalog.Category) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCategory(ctx, items)
}

func (tu *InstallUsecase) InstallPicture(c context.Context, items []media.Picture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallPicture(ctx, items)
}

func (tu *InstallUsecase) InstallManufacturer(c context.Context, items []catalog.Manufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallManufacturer(ctx, items)
}

func (tu *InstallUsecase) InstallProduct(c context.Context, items []catalog.Product) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProduct(ctx, items)
}

func (tu *InstallUsecase) InstallProductSpecificationAttribute(c context.Context, items []catalog.ProductSpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductSpecificationAttribute(ctx, items)
}

func (tu *InstallUsecase) InstallProductTag(c context.Context, items []catalog.ProductTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductTag(ctx, items)
}

func (tu *InstallUsecase) InstallProductProductTagMapping(c context.Context, items []catalog.ProductProductTagMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductProductTagMapping(ctx, items)
}

func (tu *InstallUsecase) InstallProductAttributeValue(c context.Context, items []catalog.ProductAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductAttributeValue(ctx, items)
}

func (tu *InstallUsecase) InstallProductAttributeMapping(c context.Context, items []catalog.ProductAttributeMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductAttributeMapping(ctx, items)
}

func (tu *InstallUsecase) InstallProductPicture(c context.Context, items []catalog.ProductPicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductPicture(ctx, items)
}

func (tu *InstallUsecase) InstallProductCategory(c context.Context, items []catalog.ProductCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductCategory(ctx, items)
}

func (tu *InstallUsecase) InstallTierPrice(c context.Context, items []catalog.TierPrice) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallTierPrice(ctx, items)
}

func (tu *InstallUsecase) InstallProductManufacturer(c context.Context, items []catalog.ProductManufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductManufacturer(ctx, items)
}

func (tu *InstallUsecase) InstallProductAttributeValuePicture(c context.Context, items []catalog.ProductAttributeValuePicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductAttributeValuePicture(ctx, items)
}

func (tu *InstallUsecase) InstallWarehouse(c context.Context, items []shippings.Warehouse) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallWarehouse(ctx, items)
}

func (tu *InstallUsecase) InstallVendor(c context.Context, items []vendors.Vendor) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallVendor(ctx, items)
}

func (tu *InstallUsecase) InstallAffiliate(c context.Context, items []affiliates.Affiliate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallAffiliate(ctx, items)
}

func (tu *InstallUsecase) InstallForum(c context.Context, items []forums.Forum) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallForum(ctx, items)
}

func (tu *InstallUsecase) InstallForumGroup(c context.Context, items []forums.ForumGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallForumGroup(ctx, items)
}

func (tu *InstallUsecase) InstallDiscount(c context.Context, items []discounts.Discount) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallDiscount(ctx, items)
}

func (tu *InstallUsecase) InstallBlogPost(c context.Context, items []blogs.BlogPost) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallBlogPost(ctx, items)
}

func (tu *InstallUsecase) InstallBlogComment(c context.Context, items []blogs.BlogComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallBlogComment(ctx, items)
}

func (tu *InstallUsecase) InstallPoll(c context.Context, items []polls.Poll) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallPoll(ctx, items)
}

func (tu *InstallUsecase) InstallPollAnswer(c context.Context, items []polls.PollAnswer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallPollAnswer(ctx, items)
}

func (tu *InstallUsecase) InstallNewsItem(c context.Context, items []news.NewsItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallNewsItem(ctx, items)
}

func (tu *InstallUsecase) InstallNewsComment(c context.Context, items []news.NewsComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallNewsComment(ctx, items)
}

func (tu *InstallUsecase) InstallActivityLog(c context.Context, items []loggings.ActivityLog) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallActivityLog(ctx, items)
}

func (tu *InstallUsecase) InstallSearchTerm(c context.Context, items []commons.SearchTerm) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallSearchTerm(ctx, items)
}

func (tu *InstallUsecase) InstallDownload(c context.Context, items []media.Download) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallDownload(ctx, items)
}

func (tu *InstallUsecase) InstallRelatedProduct(c context.Context, items []catalog.RelatedProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallRelatedProduct(ctx, items)
}

func (tu *InstallUsecase) InstallProductReview(c context.Context, items []catalog.ProductReview) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallProductReview(ctx, items)
}

func (tu *InstallUsecase) InstallStockQuantityChange(c context.Context, items []catalog.StockQuantityChange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallStockQuantityChange(ctx, items)
}

func (tu *InstallUsecase) InstallGdprConsent(c context.Context, items []gdprs.GdprConsent) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallGdprConsent(ctx, items)
}

func (tu *InstallUsecase) InstallOrder(c context.Context, items []orders.Order) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallOrder(ctx, items)
}

func (tu *InstallUsecase) InstallOrderItem(c context.Context, items []orders.OrderItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallOrderItem(ctx, items)
}

func (tu *InstallUsecase) InstallShipment(c context.Context, items []shippings.Shipment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallShipment(ctx, items)
}

func (tu *InstallUsecase) InstallShipmentItem(c context.Context, items []shippings.ShipmentItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallShipmentItem(ctx, items)
}

func (tu *InstallUsecase) InstallOrderNote(c context.Context, items []orders.OrderNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallOrderNote(ctx, items)
}
