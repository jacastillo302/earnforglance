package repository

import (
	"context"
	directory "earnforglance/server/domain/directory"
	domain "earnforglance/server/domain/public"
	"earnforglance/server/service/data/mongo"
	service "earnforglance/server/service/public"

	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type directoryRepository struct {
	database   mongo.Database
	collection string
}

func NewDirectoryRepository(db mongo.Database, collection string) domain.DirectoryRepository {
	return &directoryRepository{
		database:   db,
		collection: collection,
	}
}

func (dr *directoryRepository) GetCountries(c context.Context, filter domain.CountryRequest) ([]domain.CountriesResponse, error) {
	var result []domain.CountriesResponse
	var countries []directory.Country
	err := error(nil)

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var country directory.Country

		collection := dr.database.Collection(directory.CollectionCountry)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&country)
		if err != nil {
			return result, err
		}

		item, err := PrepareCountry(dr, c, country, filter)
		if err != nil {
			return result, err
		}

		result = append(result, domain.CountriesResponse{Countries: []domain.CountryResponse{item}})
		return result, err
	}

	query := bson.M{"published": true}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	for _, value := range filter.Filters {

		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else {
			query[value.Field] = value.Value
		}
	}

	buildFilter := options.Find()
	buildFilter.SetSort(bson.D{{Key: "_id", Value: sortOrder}})

	collection := dr.database.Collection(directory.CollectionCountry)
	cursor, err := collection.Find(c, query, buildFilter)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &countries)
	if err != nil {
		return result, err
	}

	var items []domain.CountryResponse
	for i := range countries {
		item, err := PrepareCountry(dr, c, countries[i], filter)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.CountriesResponse{Countries: items})

	return result, err
}

func PrepareCountry(dr *directoryRepository, c context.Context, country directory.Country, filter domain.CountryRequest) (domain.CountryResponse, error) {
	var result domain.CountryResponse
	err := error(nil)

	result.Countries = country

	for i := range filter.Content {
		switch filter.Content[i] {
		case "states":
			result.States, err = PrepareStates(dr, c, country)
		}
	}

	return result, err

}

func (dr *directoryRepository) GetCurrencies(c context.Context, filter domain.CurrencyRequest) ([]domain.CurrenciesResponse, error) {
	var result []domain.CurrenciesResponse
	var currencies []directory.Currency
	err := error(nil)

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var currency directory.Currency

		collection := dr.database.Collection(directory.CollectionCurrency)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&currency)
		if err != nil {
			return result, err
		}

		item, err := PrepareCurrency(dr, c, currency, filter)
		if err != nil {
			return result, err
		}

		result = append(result, domain.CurrenciesResponse{Currencies: []domain.CurrencyResponse{item}})
		return result, err
	}

	query := bson.M{"published": true}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	for _, value := range filter.Filters {

		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else {
			query[value.Field] = value.Value
		}
	}

	buildFilter := options.Find()
	buildFilter.SetSort(bson.D{{Key: "_id", Value: sortOrder}})

	collection := dr.database.Collection(directory.CollectionCurrency)
	cursor, err := collection.Find(c, query, buildFilter)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &currencies)
	if err != nil {
		return result, err
	}

	var items []domain.CurrencyResponse
	for i := range currencies {
		item, err := PrepareCurrency(dr, c, currencies[i], filter)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.CurrenciesResponse{Currencies: items})

	return result, err
}

func PrepareCurrency(dr *directoryRepository, c context.Context, currency directory.Currency, filter domain.CurrencyRequest) (domain.CurrencyResponse, error) {
	var result domain.CurrencyResponse
	err := error(nil)

	result.Currency = currency

	for i := range filter.Content {
		switch filter.Content[i] {
		case "exchangerate":
			result.ExchangeRate, err = PrepareExchangeRate(dr, c, currency)
		case "roundingtype":
			result.RoundingType, err = PrepareRoundingType(currency)
		}
	}

	return result, err
}

func PrepareStates(dr *directoryRepository, c context.Context, country directory.Country) ([]directory.StateProvince, error) {
	var states []directory.StateProvince
	err := error(nil)

	collection := dr.database.Collection(directory.CollectionStateProvince)

	buildFilter := options.Find()
	buildFilter.SetSort(bson.D{{Key: "display_order", Value: 1}})

	cursor, err := collection.Find(c, bson.M{"country_id": country.ID}, buildFilter)

	if err != nil {
		return states, err
	}

	err = cursor.All(c, &states)
	if err != nil {
		return states, err
	}

	return states, err
}

func PrepareExchangeRate(dr *directoryRepository, c context.Context, currency directory.Currency) ([]directory.ExchangeRate, error) {

	var rate []directory.ExchangeRate
	err := error(nil)

	collection := dr.database.Collection(directory.CollectionExchangeRate)
	limit := int64(3)

	buildFilter := options.Find()
	buildFilter.SetSort(bson.D{{Key: "updated_on", Value: -1}})
	buildFilter.SetLimit(limit)

	cursor, err := collection.Find(c, bson.M{"currency_code": currency.CurrencyCode}, buildFilter)

	if err != nil {
		return rate, err
	}

	err = cursor.All(c, &rate)
	if err != nil {
		return rate, err
	}

	return rate, err
}

func PrepareRoundingType(currency directory.Currency) (domain.RoundingType, error) {
	var Type domain.RoundingType

	items, err := service.ReadJsonMapTypes("directory", "rounding")
	if err != nil {
		return Type, err
	}

	filtered := service.FilterTypesByValue(items, currency.RoundingTypeID)

	Type.Name = filtered[0].Name
	Type.Value = filtered[0].Value
	Type.Description = filtered[0].Description

	return Type, err

}
