package install

import (
	"context"
	"time"

	settings "earnforglance/server/domain/configuration"
	directory "earnforglance/server/domain/directory"
	install "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	messages "earnforglance/server/domain/messages"
	security "earnforglance/server/domain/security"
	shipping "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
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

func (tu *InstallUsecase) InstallMessageTemplates(c context.Context) {

}

func (tu *InstallUsecase) InstallTopicTemplates(c context.Context) {

}

func (tu *InstallUsecase) InstallCustomersAndUsers(c context.Context) {

}

func (tu *InstallUsecase) InstallTopics(c context.Context) {

}

func (tu *InstallUsecase) InstallActivityLogTypes(c context.Context) {

}

func (tu *InstallUsecase) InstallProductTemplates(c context.Context) {

}

func (tu *InstallUsecase) InstallCategoryTemplates(c context.Context) {

}

func (tu *InstallUsecase) InstallManufacturerTemplates(c context.Context) {

}

func (tu *InstallUsecase) InstallScheduleTasks(c context.Context) {

}

func (tu *InstallUsecase) InstallReturnRequestReasons(c context.Context) {

}

func (tu *InstallUsecase) InstallReturnRequestActions(c context.Context) {

}

func (tu *InstallUsecase) InstallSampleData(c context.Context) {

}
