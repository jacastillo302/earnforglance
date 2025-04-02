package domain

import (
	"context"
	"time"

	settings "earnforglance/server/domain/configuration"
	directory "earnforglance/server/domain/directory"
	lang "earnforglance/server/domain/localization"
	messages "earnforglance/server/domain/messages"
	security "earnforglance/server/domain/security"
	shipping "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
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
}
