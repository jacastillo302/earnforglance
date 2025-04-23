package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type measuredimensionRepository struct {
	database   mongo.Database
	collection string
}

func NewMeasureDimensionRepository(db mongo.Database, collection string) domain.MeasureDimensionRepository {
	return &measuredimensionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *measuredimensionRepository) Create(c context.Context, measuredimension *domain.MeasureDimension) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, measuredimension)

	return err
}

func (ur *measuredimensionRepository) CreateMany(c context.Context, items []domain.MeasureDimension) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *measuredimensionRepository) Update(c context.Context, measuredimension *domain.MeasureDimension) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": measuredimension.ID}
	update := bson.M{
		"$set": measuredimension,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *measuredimensionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *measuredimensionRepository) Fetch(c context.Context) ([]domain.MeasureDimension, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var measuredimensions []domain.MeasureDimension

	err = cursor.All(c, &measuredimensions)
	if measuredimensions == nil {
		return []domain.MeasureDimension{}, err
	}

	return measuredimensions, err
}

func (tr *measuredimensionRepository) FetchByID(c context.Context, measuredimensionID string) (domain.MeasureDimension, error) {
	collection := tr.database.Collection(tr.collection)

	var measuredimension domain.MeasureDimension

	idHex, err := bson.ObjectIDFromHex(measuredimensionID)
	if err != nil {
		return measuredimension, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&measuredimension)
	return measuredimension, err
}
