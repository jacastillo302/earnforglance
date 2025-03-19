package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tierpriceRepository struct {
	database   mongo.Database
	collection string
}

func NewTierPriceRepository(db mongo.Database, collection string) domain.TierPriceRepository {
	return &tierpriceRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *tierpriceRepository) Create(c context.Context, tierprice *domain.TierPrice) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, tierprice)

	return err
}

func (ur *tierpriceRepository) Update(c context.Context, tierprice *domain.TierPrice) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": tierprice.ID}
	update := bson.M{
		"$set": tierprice,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *tierpriceRepository) Delete(c context.Context, tierprice *domain.TierPrice) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": tierprice.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *tierpriceRepository) Fetch(c context.Context) ([]domain.TierPrice, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var tierprices []domain.TierPrice

	err = cursor.All(c, &tierprices)
	if tierprices == nil {
		return []domain.TierPrice{}, err
	}

	return tierprices, err
}

func (tr *tierpriceRepository) FetchByID(c context.Context, tierpriceID string) (domain.TierPrice, error) {
	collection := tr.database.Collection(tr.collection)

	var tierprice domain.TierPrice

	idHex, err := primitive.ObjectIDFromHex(tierpriceID)
	if err != nil {
		return tierprice, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&tierprice)
	return tierprice, err
}
