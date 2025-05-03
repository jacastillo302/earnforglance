package repository

import (
	"context"
	configuration "earnforglance/server/domain/configuration"
	domain "earnforglance/server/domain/public"
	store "earnforglance/server/domain/stores"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type configurationRepository struct {
	database   mongo.Database
	collection string
}

func NewConfigurationRepository(db mongo.Database, collection string) domain.ConfigurationRepository {
	return &configurationRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *configurationRepository) GetConfigurations(c context.Context, filter domain.ConfigurationRequest) ([]domain.ConfigurationsResponse, error) {
	var result []domain.ConfigurationsResponse
	var configurations []configuration.Setting

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var configurationRecord configuration.Setting

		collection := cr.database.Collection(configuration.CollectionSetting)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&configurationRecord)
		if err != nil {
			return result, err
		}

		item, err := PrepareConfiguration(cr, c, configurationRecord, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.ConfigurationsResponse{Configurations: []domain.ConfigurationResponse{item}})
		return result, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	query := bson.M{}

	limit := int64(filter.Limit)
	//skip := int64(filter.Page * filter.Limit)

	for _, value := range filter.Filters {
		// "contains", "eq", etc.
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else {
			query[value.Field] = value.Value
		}

		//skip = 0
	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: sortOrder}}).
		SetLimit(limit)

	collection := cr.database.Collection(configuration.CollectionSetting)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &configurations)
	if err != nil {
		return result, err
	}

	var items []domain.ConfigurationResponse
	for i := range configurations {
		item, err := PrepareConfiguration(cr, c, configurations[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.ConfigurationsResponse{Configurations: items})

	return result, err
}

func PrepareConfiguration(vr *configurationRepository, c context.Context, configuration configuration.Setting, content []string, lang string) (domain.ConfigurationResponse, error) {
	var result domain.ConfigurationResponse
	err := error(nil)

	result.Configuration = configuration

	for i := range content {
		switch content[i] {
		case "store":
			storeID := configuration.StoreID
			result.Store, err = PrepareStore(vr, c, storeID)
		}
	}

	return result, err
}

func PrepareStore(vr *configurationRepository, c context.Context, ID bson.ObjectID) (*store.Store, error) {
	store, err := GetStoreByID(c, ID, vr.database.Collection(store.CollectionStore))
	if err != nil {
		return nil, err
	}
	return store, err
}

func GetSettingByName(c context.Context, name string, collection mongo.Collection) (configuration.Setting, error) {
	var item configuration.Setting
	err := collection.FindOne(c, bson.M{"name": name}).Decode(&item)
	return item, err
}

func GetSettingByNames(c context.Context, names []string, collection mongo.Collection) ([]configuration.Setting, error) {
	var settings []configuration.Setting

	findOptions := options.Find().
		SetSort(bson.D{{Key: "name", Value: 1}})

	query := bson.M{
		"name": bson.M{"$in": names},
	}

	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return settings, err
	}

	err = cursor.All(c, &settings)
	if err != nil {
		return settings, err
	}

	return settings, err
}
