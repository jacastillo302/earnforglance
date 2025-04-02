package repository

import (
	"context"

	settings "earnforglance/server/domain/configuration"
	directory "earnforglance/server/domain/directory"
	install "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
	"earnforglance/server/service/data/mongo"
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

func (tu *intallRepository) InstallCurrencies(c context.Context, items []directory.Currency) error {

	collection := tu.database.Collection(directory.CollectionCurrency)

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

func (tu *intallRepository) InstallSettings(c context.Context, items []settings.Setting) error {

	collection := tu.database.Collection(stores.CollectionStore)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}
