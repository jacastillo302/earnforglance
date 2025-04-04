package repository

import (
	"context"

	catalog "earnforglance/server/domain/catalog"
	commons "earnforglance/server/domain/common"
	settings "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	directory "earnforglance/server/domain/directory"
	install "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	loggings "earnforglance/server/domain/logging"
	messages "earnforglance/server/domain/messages"
	orders "earnforglance/server/domain/orders"
	tasks "earnforglance/server/domain/scheduleTasks"
	security "earnforglance/server/domain/security"
	shipping "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
	topics "earnforglance/server/domain/topics"
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

func (tu *intallRepository) InstallShippingMethod(c context.Context, items []shipping.ShippingMethod) error {

	collection := tu.database.Collection(shipping.CollectionShippingMethod)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallDeliveryDate(c context.Context, items []shipping.DeliveryDate) error {

	collection := tu.database.Collection(shipping.CollectionDeliveryDate)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (tu *intallRepository) InstallProductAvailabilityRange(c context.Context, items []shipping.ProductAvailabilityRange) error {

	collection := tu.database.Collection(shipping.CollectionProductAvailabilityRange)

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

func (tu *intallRepository) InstallCustomerAddress(c context.Context, items []commons.Address) error {

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
