package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type shoppingcartsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewShoppingCartSettingsRepository(db mongo.Database, collection string) domain.ShoppingCartSettingsRepository {
	return &shoppingcartsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shoppingcartsettingsRepository) CreateMany(c context.Context, items []domain.ShoppingCartSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *shoppingcartsettingsRepository) Create(c context.Context, shoppingcartsettings *domain.ShoppingCartSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shoppingcartsettings)

	return err
}

func (ur *shoppingcartsettingsRepository) Update(c context.Context, shoppingcartsettings *domain.ShoppingCartSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shoppingcartsettings.ID}
	update := bson.M{
		"$set": shoppingcartsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shoppingcartsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shoppingcartsettingsRepository) Fetch(c context.Context) ([]domain.ShoppingCartSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shoppingcartsettingss []domain.ShoppingCartSettings

	err = cursor.All(c, &shoppingcartsettingss)
	if shoppingcartsettingss == nil {
		return []domain.ShoppingCartSettings{}, err
	}

	return shoppingcartsettingss, err
}

func (tr *shoppingcartsettingsRepository) FetchByID(c context.Context, shoppingcartsettingsID string) (domain.ShoppingCartSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var shoppingcartsettings domain.ShoppingCartSettings

	idHex, err := primitive.ObjectIDFromHex(shoppingcartsettingsID)
	if err != nil {
		return shoppingcartsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shoppingcartsettings)
	return shoppingcartsettings, err
}
