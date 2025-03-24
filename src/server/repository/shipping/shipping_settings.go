package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type shippingsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewShippingSettingsRepository(db mongo.Database, collection string) domain.ShippingSettingsRepository {
	return &shippingsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shippingsettingsRepository) CreateMany(c context.Context, items []domain.ShippingSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *shippingsettingsRepository) Create(c context.Context, shippingsettings *domain.ShippingSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shippingsettings)

	return err
}

func (ur *shippingsettingsRepository) Update(c context.Context, shippingsettings *domain.ShippingSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shippingsettings.ID}
	update := bson.M{
		"$set": shippingsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shippingsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shippingsettingsRepository) Fetch(c context.Context) ([]domain.ShippingSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shippingsettingss []domain.ShippingSettings

	err = cursor.All(c, &shippingsettingss)
	if shippingsettingss == nil {
		return []domain.ShippingSettings{}, err
	}

	return shippingsettingss, err
}

func (tr *shippingsettingsRepository) FetchByID(c context.Context, shippingsettingsID string) (domain.ShippingSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var shippingsettings domain.ShippingSettings

	idHex, err := primitive.ObjectIDFromHex(shippingsettingsID)
	if err != nil {
		return shippingsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shippingsettings)
	return shippingsettings, err
}
