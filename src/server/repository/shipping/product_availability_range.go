package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productavailabilityrangeRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAvailabilityRangeRepository(db mongo.Database, collection string) domain.ProductAvailabilityRangeRepository {
	return &productavailabilityrangeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productavailabilityrangeRepository) CreateMany(c context.Context, items []domain.ProductAvailabilityRange) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productavailabilityrangeRepository) Create(c context.Context, productavailabilityrange *domain.ProductAvailabilityRange) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productavailabilityrange)

	return err
}

func (ur *productavailabilityrangeRepository) Update(c context.Context, productavailabilityrange *domain.ProductAvailabilityRange) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productavailabilityrange.ID}
	update := bson.M{
		"$set": productavailabilityrange,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productavailabilityrangeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productavailabilityrangeRepository) Fetch(c context.Context) ([]domain.ProductAvailabilityRange, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productavailabilityranges []domain.ProductAvailabilityRange

	err = cursor.All(c, &productavailabilityranges)
	if productavailabilityranges == nil {
		return []domain.ProductAvailabilityRange{}, err
	}

	return productavailabilityranges, err
}

func (tr *productavailabilityrangeRepository) FetchByID(c context.Context, productavailabilityrangeID string) (domain.ProductAvailabilityRange, error) {
	collection := tr.database.Collection(tr.collection)

	var productavailabilityrange domain.ProductAvailabilityRange

	idHex, err := primitive.ObjectIDFromHex(productavailabilityrangeID)
	if err != nil {
		return productavailabilityrange, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productavailabilityrange)
	return productavailabilityrange, err
}
