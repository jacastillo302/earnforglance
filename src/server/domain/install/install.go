package domain

import (
	"context"
	"time"

	settings "earnforglance/server/domain/configuration"
	directory "earnforglance/server/domain/directory"
	lang "earnforglance/server/domain/localization"
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
	InstallStores(c context.Context, stores []stores.Store) error
	InstallSettings(c context.Context, settings []settings.Setting) error
	InstallCurrencies(c context.Context, items []directory.Currency) error
	InstallMeasureDimension(c context.Context, items []directory.MeasureDimension) error
	InstallMeasureWeight(c context.Context, items []directory.MeasureWeight) error
	InstallTaxCategory(c context.Context, items []taxes.TaxCategory) error
	InstallLanguages(c context.Context, items []lang.Language) error
}

// GdprLogUsecase interface
type InstallLogUsecase interface {
	PingDatabase(c context.Context) error
	InstallStores(c context.Context, stores []stores.Store) error
	InstallSettings(c context.Context, settings []settings.Setting) error
	InstallCurrencies(c context.Context, items []directory.Currency) error
	InstallMeasureDimension(c context.Context, items []directory.MeasureDimension) error
	InstallMeasureWeight(c context.Context, items []directory.MeasureWeight) error
	InstallTaxCategory(c context.Context, items []taxes.TaxCategory) error
	InstallLanguages(c context.Context, items []lang.Language) error
}
