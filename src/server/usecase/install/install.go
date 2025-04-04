package install

import (
	"context"
	"time"

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

func (tu *InstallUsecase) InstallShippingMethod(c context.Context, items []shipping.ShippingMethod) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallShippingMethod(ctx, items)
}

func (tu *InstallUsecase) InstallDeliveryDate(c context.Context, items []shipping.DeliveryDate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallDeliveryDate(ctx, items)
}

func (tu *InstallUsecase) InstallProductAvailabilityRange(c context.Context, items []shipping.ProductAvailabilityRange) error {
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

func (tu *InstallUsecase) InstallCustomerAddress(c context.Context, items []commons.Address) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.InstallRepository.InstallCustomerAddress(ctx, items)
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

func (tu *InstallUsecase) InstallSampleData(c context.Context) {

}
