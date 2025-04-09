package repository

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
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type intallRepository struct {
	database mongo.Database
}

// NewInstallRepository creates a new instance of intallRepository
func NewInstallRepository(db mongo.Database) install.InstallRepository {
	return &intallRepository{
		database: db,
	}
}

func (ur *intallRepository) PingDatabase(c context.Context) error {
	client := ur.database.Client()
	err := client.Ping(context.Background())
	return err
}

func (tu *intallRepository) InstallPermissionRecord(c context.Context, items []security.PermissionRecord) error {

	collection := tu.database.Collection(security.CollectionPermissionRecord)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCurrencies(c context.Context, items []directory.Currency) error {

	collection := tu.database.Collection(directory.CollectionCurrency)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallMeasureDimension(c context.Context, items []directory.MeasureDimension) error {

	collection := tu.database.Collection(directory.CollectionMeasureDimension)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallMeasureWeight(c context.Context, items []directory.MeasureWeight) error {

	collection := tu.database.Collection(directory.CollectionMeasureWeight)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallTaxCategory(c context.Context, items []taxes.TaxCategory) error {

	collection := tu.database.Collection(taxes.CollectionTaxCategory)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallLanguages(c context.Context, items []lang.Language) error {

	collection := tu.database.Collection(lang.CollectionLanguage)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallLocaleStringResource(c context.Context, items []lang.LocaleStringResource) error {

	collection := tu.database.Collection(lang.CollectionLocaleStringResource)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallStores(c context.Context, items []stores.Store) error {

	collection := tu.database.Collection(stores.CollectionStore)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallSettings(c context.Context, items []settings.Setting) error {

	collection := tu.database.Collection(settings.CollectionSetting)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCountries(c context.Context, items []directory.Country) error {

	collection := tu.database.Collection(directory.CollectionCountry)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallStateProvince(c context.Context, items []directory.StateProvince) error {

	collection := tu.database.Collection(directory.CollectionStateProvince)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallShippingMethod(c context.Context, items []shippings.ShippingMethod) error {

	collection := tu.database.Collection(shippings.CollectionShippingMethod)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallDeliveryDate(c context.Context, items []shippings.DeliveryDate) error {

	collection := tu.database.Collection(shippings.CollectionDeliveryDate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductAvailabilityRange(c context.Context, items []shippings.ProductAvailabilityRange) error {

	collection := tu.database.Collection(shippings.CollectionProductAvailabilityRange)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallEmailAccount(c context.Context, items []messages.EmailAccount) error {

	collection := tu.database.Collection(messages.CollectionEmailAccount)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallMessageTemplate(c context.Context, items []messages.MessageTemplate) error {

	collection := tu.database.Collection(messages.CollectionMessageTemplate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallTopicTemplate(c context.Context, items []topics.TopicTemplate) error {

	collection := tu.database.Collection(topics.CollectionTopicTemplate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCustomerRole(c context.Context, items []customers.CustomerRole) error {

	collection := tu.database.Collection(customers.CollectionCustomerRole)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCustomer(c context.Context, items []customers.Customer) error {

	collection := tu.database.Collection(customers.CollectionCustomer)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) UpdateCustomer(c context.Context, customer customers.Customer) error {

	collection := tu.database.Collection(customers.CollectionCustomer)

	filter := bson.M{"_id": customer.ID}
	update := bson.M{
		"$set": customer,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (tu *intallRepository) InstallCustomerPassword(c context.Context, items []customers.CustomerPassword) error {

	collection := tu.database.Collection(customers.CollectionCustomerPassword)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallAddress(c context.Context, items []commons.Address) error {

	collection := tu.database.Collection(commons.CollectionAddress)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCustomerAddressMapping(c context.Context, item customers.CustomerAddressMapping) error {

	collection := tu.database.Collection(customers.CollectionCustomerAddressMapping)

	_, err := collection.InsertOne(c, item)

	return err
}

func (tu *intallRepository) InstallCustomerCustomerRoleMapping(c context.Context, items []customers.CustomerCustomerRoleMapping) error {

	collection := tu.database.Collection(customers.CollectionCustomerCustomerRoleMapping)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallTopic(c context.Context, items []topics.Topic) error {

	collection := tu.database.Collection(topics.CollectionTopic)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallActivityLogType(c context.Context, items []loggings.ActivityLogType) error {

	collection := tu.database.Collection(loggings.CollectionActivityLogType)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductTemplate(c context.Context, items []catalog.ProductTemplate) error {

	collection := tu.database.Collection(catalog.CollectionProductTemplate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCategoryTemplate(c context.Context, items []catalog.CategoryTemplate) error {

	collection := tu.database.Collection(catalog.CollectionCategoryTemplate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallManufacturerTemplate(c context.Context, items []catalog.ManufacturerTemplate) error {

	collection := tu.database.Collection(catalog.CollectionManufacturerTemplate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallScheduleTask(c context.Context, items []tasks.ScheduleTask) error {

	collection := tu.database.Collection(tasks.CollectionScheduleTask)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallReturnRequestReason(c context.Context, items []orders.ReturnRequestReason) error {

	collection := tu.database.Collection(orders.CollectionReturnRequestReason)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallReturnRequestAction(c context.Context, items []orders.ReturnRequestAction) error {

	collection := tu.database.Collection(orders.CollectionReturnRequestAction)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCheckoutAttribute(c context.Context, items []orders.CheckoutAttribute) error {

	collection := tu.database.Collection(orders.CollectionCheckoutAttribute)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCheckoutAttributeValue(c context.Context, items []orders.CheckoutAttributeValue) error {

	collection := tu.database.Collection(orders.CollectionCheckoutAttributeValue)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallSpecificationAttribute(c context.Context, items []catalog.SpecificationAttribute) error {

	collection := tu.database.Collection(catalog.CollectionSpecificationAttribute)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallSpecificationAttributeOption(c context.Context, items []catalog.SpecificationAttributeOption) error {

	collection := tu.database.Collection(catalog.CollectionSpecificationAttributeOption)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallSpecificationAttributeGroup(c context.Context, items []catalog.SpecificationAttributeGroup) error {

	collection := tu.database.Collection(catalog.CollectionSpecificationAttributeGroup)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductAttribute(c context.Context, items []catalog.ProductAttribute) error {

	collection := tu.database.Collection(catalog.CollectionProductAttribute)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallCategory(c context.Context, items []catalog.Category) error {

	collection := tu.database.Collection(catalog.CollectionCategory)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallPicture(c context.Context, items []media.Picture) error {

	collection := tu.database.Collection(media.CollectionPicture)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallManufacturer(c context.Context, items []catalog.Manufacturer) error {

	collection := tu.database.Collection(catalog.CollectionManufacturer)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProduct(c context.Context, items []catalog.Product) error {

	collection := tu.database.Collection(catalog.CollectionProduct)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductSpecificationAttribute(c context.Context, items []catalog.ProductSpecificationAttribute) error {

	collection := tu.database.Collection(catalog.CollectionProductSpecificationAttribute)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductTag(c context.Context, items []catalog.ProductTag) error {

	collection := tu.database.Collection(catalog.CollectionProductTag)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductProductTagMapping(c context.Context, items []catalog.ProductProductTagMapping) error {

	collection := tu.database.Collection(catalog.CollectionProductProductTagMapping)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductAttributeValue(c context.Context, items []catalog.ProductAttributeValue) error {

	collection := tu.database.Collection(catalog.CollectionProductAttributeValue)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductAttributeMapping(c context.Context, items []catalog.ProductAttributeMapping) error {

	collection := tu.database.Collection(catalog.CollectionProductAttributeMapping)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductPicture(c context.Context, items []catalog.ProductPicture) error {

	collection := tu.database.Collection(catalog.CollectionProductPicture)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductCategory(c context.Context, items []catalog.ProductCategory) error {

	collection := tu.database.Collection(catalog.CollectionProductCategory)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallTierPrice(c context.Context, items []catalog.TierPrice) error {

	collection := tu.database.Collection(commons.CollectionSearchTerm)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductManufacturer(c context.Context, items []catalog.ProductManufacturer) error {

	collection := tu.database.Collection(catalog.CollectionProductManufacturer)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductAttributeValuePicture(c context.Context, items []catalog.ProductAttributeValuePicture) error {

	collection := tu.database.Collection(catalog.CollectionProductAttributeValuePicture)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallWarehouse(c context.Context, items []shippings.Warehouse) error {

	collection := tu.database.Collection(shippings.CollectionWarehouse)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallVendor(c context.Context, items []vendors.Vendor) error {

	collection := tu.database.Collection(vendors.CollectionVendor)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallAffiliate(c context.Context, items []affiliates.Affiliate) error {

	collection := tu.database.Collection(affiliates.CollectionAffiliate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallForum(c context.Context, items []forums.Forum) error {

	collection := tu.database.Collection(forums.CollectionForum)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallForumGroup(c context.Context, items []forums.ForumGroup) error {

	collection := tu.database.Collection(forums.CollectionForumGroup)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallDiscount(c context.Context, items []discounts.Discount) error {

	collection := tu.database.Collection(discounts.CollectionDiscount)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallBlogPost(c context.Context, items []blogs.BlogPost) error {

	collection := tu.database.Collection(blogs.CollectionBlogPost)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallBlogComment(c context.Context, items []blogs.BlogComment) error {

	collection := tu.database.Collection(blogs.CollectionBlogComment)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallPoll(c context.Context, items []polls.Poll) error {

	collection := tu.database.Collection(polls.CollectionPoll)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallPollAnswer(c context.Context, items []polls.PollAnswer) error {

	collection := tu.database.Collection(polls.CollectionPollAnswer)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallNewsItem(c context.Context, items []news.NewsItem) error {

	collection := tu.database.Collection(news.CollectionNewsItem)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallNewsComment(c context.Context, items []news.NewsComment) error {

	collection := tu.database.Collection(news.CollectionNewsComment)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallActivityLog(c context.Context, items []loggings.ActivityLog) error {

	collection := tu.database.Collection(loggings.CollectionActivityLog)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallSearchTerm(c context.Context, items []commons.SearchTerm) error {

	collection := tu.database.Collection(commons.CollectionSearchTerm)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallDownload(c context.Context, items []media.Download) error {

	collection := tu.database.Collection(media.CollectionDownload)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}
