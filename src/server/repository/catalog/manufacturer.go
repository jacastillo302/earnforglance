package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type manufacturerRepository struct {
	database   mongo.Database
	collection string
}

func NewManufacturerRepository(db mongo.Database, collection string) domain.ManufacturerRepository {
	return &manufacturerRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *manufacturerRepository) CreateMany(c context.Context, items []domain.Manufacturer) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *manufacturerRepository) Create(c context.Context, manufacturer *domain.Manufacturer) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, manufacturer)

	return err
}

func (ur *manufacturerRepository) Update(c context.Context, manufacturer *domain.Manufacturer) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": manufacturer.ID}
	update := bson.M{
		"$set": manufacturer,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *manufacturerRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *manufacturerRepository) Fetch(c context.Context) ([]domain.Manufacturer, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var manufacturers []domain.Manufacturer

	err = cursor.All(c, &manufacturers)
	if manufacturers == nil {
		return []domain.Manufacturer{}, err
	}

	return manufacturers, err
}

func (tr *manufacturerRepository) FetchByID(c context.Context, manufacturerID string) (domain.Manufacturer, error) {
	collection := tr.database.Collection(tr.collection)

	var manufacturer domain.Manufacturer

	idHex, err := primitive.ObjectIDFromHex(manufacturerID)
	if err != nil {
		return manufacturer, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&manufacturer)
	return manufacturer, err
}
