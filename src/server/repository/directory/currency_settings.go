package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type currencysettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewCurrencySettingsRepository(db mongo.Database, collection string) domain.CurrencySettingsRepository {
	return &currencysettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *currencysettingsRepository) CreateMany(c context.Context, items []domain.CurrencySettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *currencysettingsRepository) Create(c context.Context, currencysettings *domain.CurrencySettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, currencysettings)

	return err
}

func (ur *currencysettingsRepository) Update(c context.Context, currencysettings *domain.CurrencySettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": currencysettings.ID}
	update := bson.M{
		"$set": currencysettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *currencysettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *currencysettingsRepository) Fetch(c context.Context) ([]domain.CurrencySettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var currencysettingss []domain.CurrencySettings

	err = cursor.All(c, &currencysettingss)
	if currencysettingss == nil {
		return []domain.CurrencySettings{}, err
	}

	return currencysettingss, err
}

func (tr *currencysettingsRepository) FetchByID(c context.Context, currencysettingsID string) (domain.CurrencySettings, error) {
	collection := tr.database.Collection(tr.collection)

	var currencysettings domain.CurrencySettings

	idHex, err := bson.ObjectIDFromHex(currencysettingsID)
	if err != nil {
		return currencysettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&currencysettings)
	return currencysettings, err
}
