package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type shippingoptionRepository struct {
	database   mongo.Database
	collection string
}

func NewShippingOptionRepository(db mongo.Database, collection string) domain.ShippingOptionRepository {
	return &shippingoptionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shippingoptionRepository) CreateMany(c context.Context, items []domain.ShippingOption) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *shippingoptionRepository) Create(c context.Context, shippingoption *domain.ShippingOption) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shippingoption)

	return err
}

func (ur *shippingoptionRepository) Update(c context.Context, shippingoption *domain.ShippingOption) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shippingoption.ID}
	update := bson.M{
		"$set": shippingoption,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shippingoptionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shippingoptionRepository) Fetch(c context.Context) ([]domain.ShippingOption, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shippingoptions []domain.ShippingOption

	err = cursor.All(c, &shippingoptions)
	if shippingoptions == nil {
		return []domain.ShippingOption{}, err
	}

	return shippingoptions, err
}

func (tr *shippingoptionRepository) FetchByID(c context.Context, shippingoptionID string) (domain.ShippingOption, error) {
	collection := tr.database.Collection(tr.collection)

	var shippingoption domain.ShippingOption

	idHex, err := bson.ObjectIDFromHex(shippingoptionID)
	if err != nil {
		return shippingoption, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shippingoption)
	return shippingoption, err
}
