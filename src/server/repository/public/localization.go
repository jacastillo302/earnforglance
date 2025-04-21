package repository

import (
	"context"
	directory "earnforglance/server/domain/directory"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	"earnforglance/server/service/data/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type localizationRepository struct {
	database   mongo.Database
	collection string
}

func NewLocalizationRepository(db mongo.Database, collection string) domain.LocalizationRepository {
	return &localizationRepository{
		database:   db,
		collection: collection,
	}
}

func (lr *localizationRepository) GetLocalizations(c context.Context, filter domain.LocalizationRequest) ([]domain.LocalizationsResponse, error) {
	var result []domain.LocalizationsResponse
	var languages []localization.Language
	err := error(nil)

	idHex, err := primitive.ObjectIDFromHex(filter.ID)
	if err == nil {
		var loc localization.Language

		collection := lr.database.Collection(localization.CollectionLanguage)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&loc)
		if err != nil {
			return result, err
		}

		item, err := PrepareLocalization(lr, c, loc, filter)
		if err != nil {
			return result, err
		}

		result = append(result, domain.LocalizationsResponse{Localizations: []domain.LocalizationResponse{item}})
		return result, err
	}
	query := bson.M{"published": true}

	if filter.Rtl {
		query["rtl"] = filter.Rtl
	}

	if filter.Lang != "" {
		query["unique_seo_code"] = filter.Lang
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	for _, value := range filter.Filters {

		bEnable := true
		switch value.Field {
		case "entity_id":
			bEnable = false
		case "locale_key_group":
			bEnable = false
		case "locale_key":
			bEnable = false
		case "locale_value":
			bEnable = false
		case "language_id":
			bEnable = false
		case "resource_name":
			bEnable = false
		case "resource_value":
			bEnable = false
		}

		if bEnable {

			if value.Operator == "contains" {
				query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
			} else if value.Operator == "not_contains" {
				query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
			} else {
				query[value.Field] = value.Value
			}
		}
	}

	fmt.Println("query", query)
	findOptions := options.Find().
		SetSort(bson.D{{Key: "display_order", Value: sortOrder}}).
		SetLimit(int64(filter.Limit))

	collection := lr.database.Collection(localization.CollectionLanguage)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &languages)
	if err != nil {
		return result, err
	}

	var items []domain.LocalizationResponse
	for i := range languages {
		item, err := PrepareLocalization(lr, c, languages[i], filter)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.LocalizationsResponse{Localizations: items})

	return result, err
}

func PrepareLocalization(lr *localizationRepository, c context.Context, loc localization.Language, filter domain.LocalizationRequest) (domain.LocalizationResponse, error) {
	var result domain.LocalizationResponse
	err := error(nil)

	for i := range filter.Content {
		switch filter.Content[i] {
		case "currency":
			result.Currency, err = PrepareCurrency(lr, c, loc)
		case "resources":
			result.Resources, err = PrepareResources(lr, c, loc, filter)
		case "propertie":
			result.Properties, err = PrepareProperties(lr, c, loc, filter)
		}
	}

	result.Language = loc
	return result, err
}

func PrepareCurrency(lr *localizationRepository, c context.Context, lang localization.Language) (directory.Currency, error) {
	var result directory.Currency
	err := error(nil)
	fmt.Println("PrepareCurrency", lang.DefaultCurrencyID)
	collection := lr.database.Collection(directory.CollectionCurrency)
	err = collection.FindOne(c, bson.M{"_id": lang.DefaultCurrencyID}).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, err
}

func PrepareResources(lr *localizationRepository, c context.Context, lang localization.Language, filter domain.LocalizationRequest) ([]localization.LocaleStringResource, error) {
	var result []localization.LocaleStringResource
	err := error(nil)
	query := bson.M{"language_id": lang.ID}

	for _, value := range filter.Filters {

		bEnable := false
		switch value.Field {
		case "resource_name":
			bEnable = true
		case "resource_value":
			bEnable = true
		case "language_id":
			bEnable = true
		}

		if bEnable {

			if value.Operator == "contains" {
				query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
			} else if value.Operator == "not_contains" {
				query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
			} else {
				query[value.Field] = value.Value
			}
		}
	}

	collection := lr.database.Collection(localization.CollectionLocaleStringResource)
	cursor, err := collection.Find(c, query)

	if err != nil {
		return result, err
	}

	err = cursor.All(c, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func PrepareProperties(lr *localizationRepository, c context.Context, lang localization.Language, filter domain.LocalizationRequest) ([]localization.LocalizedProperty, error) {
	var result []localization.LocalizedProperty
	err := error(nil)
	query := bson.M{"language_id": lang.ID}

	for _, value := range filter.Filters {

		bEnable := false
		switch value.Field {
		case "entity_id":
			bEnable = true
		case "locale_key_group":
			bEnable = true
		case "locale_key":
			bEnable = true
		case "locale_value":
			bEnable = true
		case "language_id":
			bEnable = true
		}

		if bEnable {

			if value.Operator == "contains" {
				query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
			} else if value.Operator == "not_contains" {
				query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
			} else {
				query[value.Field] = value.Value
			}
		}
	}

	collection := lr.database.Collection(localization.CollectionLocalizedProperty)
	cursor, err := collection.Find(c, query)

	if err != nil {
		return result, err
	}

	err = cursor.All(c, &result)
	if err != nil {
		return result, err
	}

	return result, err
}
