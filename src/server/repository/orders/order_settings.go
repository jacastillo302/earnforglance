package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ordersettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderSettingsRepository(db mongo.Database, collection string) domain.OrderSettingsRepository {
	return &ordersettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *ordersettingsRepository) CreateMany(c context.Context, items []domain.OrderSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *ordersettingsRepository) Create(c context.Context, ordersettings *domain.OrderSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, ordersettings)

	return err
}

func (ur *ordersettingsRepository) Update(c context.Context, ordersettings *domain.OrderSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": ordersettings.ID}
	update := bson.M{
		"$set": ordersettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *ordersettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *ordersettingsRepository) Fetch(c context.Context) ([]domain.OrderSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var ordersettingss []domain.OrderSettings

	err = cursor.All(c, &ordersettingss)
	if ordersettingss == nil {
		return []domain.OrderSettings{}, err
	}

	return ordersettingss, err
}

func (tr *ordersettingsRepository) FetchByID(c context.Context, ordersettingsID string) (domain.OrderSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var ordersettings domain.OrderSettings

	idHex, err := primitive.ObjectIDFromHex(ordersettingsID)
	if err != nil {
		return ordersettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&ordersettings)
	return ordersettings, err
}
