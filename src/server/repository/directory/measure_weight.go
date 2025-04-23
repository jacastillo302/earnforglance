package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type measureweightRepository struct {
	database   mongo.Database
	collection string
}

func NewMeasureWeightRepository(db mongo.Database, collection string) domain.MeasureWeightRepository {
	return &measureweightRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *measureweightRepository) CreateMany(c context.Context, items []domain.MeasureWeight) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *measureweightRepository) Create(c context.Context, measureweight *domain.MeasureWeight) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, measureweight)

	return err
}

func (ur *measureweightRepository) Update(c context.Context, measureweight *domain.MeasureWeight) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": measureweight.ID}
	update := bson.M{
		"$set": measureweight,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *measureweightRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *measureweightRepository) Fetch(c context.Context) ([]domain.MeasureWeight, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var measureweights []domain.MeasureWeight

	err = cursor.All(c, &measureweights)
	if measureweights == nil {
		return []domain.MeasureWeight{}, err
	}

	return measureweights, err
}

func (tr *measureweightRepository) FetchByID(c context.Context, measureweightID string) (domain.MeasureWeight, error) {
	collection := tr.database.Collection(tr.collection)

	var measureweight domain.MeasureWeight

	idHex, err := bson.ObjectIDFromHex(measureweightID)
	if err != nil {
		return measureweight, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&measureweight)
	return measureweight, err
}
