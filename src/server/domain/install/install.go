package domain

import (
	"context"
	"time"

	catalog "earnforglance/server/domain/catalog"
	commons "earnforglance/server/domain/common"
	settings "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	directory "earnforglance/server/domain/directory"
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
	InstallShippingMethod(c context.Context, items []shipping.ShippingMethod) error
	InstallDeliveryDate(c context.Context, items []shipping.DeliveryDate) error
	InstallProductAvailabilityRange(c context.Context, items []shipping.ProductAvailabilityRange) error
	InstallEmailAccount(c context.Context, items []messages.EmailAccount) error
	InstallMessageTemplate(c context.Context, items []messages.MessageTemplate) error
	InstallTopicTemplate(c context.Context, items []topics.TopicTemplate) error
	InstallCustomerRole(c context.Context, items []customers.CustomerRole) error
	InstallCustomer(c context.Context, items []customers.Customer) error
	InstallCustomerPassword(c context.Context, items []customers.CustomerPassword) error
	InstallCustomerAddress(c context.Context, items []commons.Address) error
	InstallCustomerAddressMapping(c context.Context, item customers.CustomerAddressMapping) error
	UpdateCustomer(c context.Context, items customers.Customer) error
	InstallCustomerCustomerRoleMapping(c context.Context, items []customers.CustomerCustomerRoleMapping) error
	InstallTopic(c context.Context, items []topics.Topic) error
	InstallActivityLogType(c context.Context, items []loggings.ActivityLogType) error
	InstallProductTemplate(c context.Context, items []catalog.ProductTemplate) error
	InstallCategoryTemplate(c context.Context, items []catalog.CategoryTemplate) error
	InstallManufacturerTemplate(c context.Context, items []catalog.ManufacturerTemplate) error
	InstallScheduleTask(c context.Context, items []tasks.ScheduleTask) error
	InstallReturnRequestReason(c context.Context, items []orders.ReturnRequestReason) error
	InstallReturnRequestAction(c context.Context, items []orders.ReturnRequestAction) error
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
	InstallShippingMethod(c context.Context, items []shipping.ShippingMethod) error
	InstallDeliveryDate(c context.Context, items []shipping.DeliveryDate) error
	InstallProductAvailabilityRange(c context.Context, items []shipping.ProductAvailabilityRange) error
	InstallEmailAccount(c context.Context, items []messages.EmailAccount) error
	InstallMessageTemplate(c context.Context, items []messages.MessageTemplate) error
	InstallTopicTemplate(c context.Context, items []topics.TopicTemplate) error
	InstallCustomerRole(c context.Context, items []customers.CustomerRole) error
	InstallCustomer(c context.Context, items []customers.Customer) error
	InstallCustomerPassword(c context.Context, items []customers.CustomerPassword) error
	InstallCustomerAddress(c context.Context, items []commons.Address) error
	InstallCustomerAddressMapping(c context.Context, item customers.CustomerAddressMapping) error
	UpdateCustomer(c context.Context, items customers.Customer) error
	InstallCustomerCustomerRoleMapping(c context.Context, items []customers.CustomerCustomerRoleMapping) error
	InstallTopic(c context.Context, items []topics.Topic) error
	InstallActivityLogType(c context.Context, items []loggings.ActivityLogType) error
	InstallProductTemplate(c context.Context, items []catalog.ProductTemplate) error
	InstallCategoryTemplate(c context.Context, items []catalog.CategoryTemplate) error
	InstallManufacturerTemplate(c context.Context, items []catalog.ManufacturerTemplate) error
	InstallScheduleTask(c context.Context, items []tasks.ScheduleTask) error
	InstallReturnRequestReason(c context.Context, items []orders.ReturnRequestReason) error
	InstallReturnRequestAction(c context.Context, items []orders.ReturnRequestAction) error
}
