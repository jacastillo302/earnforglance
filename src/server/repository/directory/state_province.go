package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type stateprovinceRepository struct {
	database   mongo.Database
	collection string
}

func NewStateProvinceRepository(db mongo.Database, collection string) domain.StateProvinceRepository {
	return &stateprovinceRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *stateprovinceRepository) CreateMany(c context.Context, items []domain.StateProvince) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *stateprovinceRepository) Create(c context.Context, stateprovince *domain.StateProvince) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, stateprovince)

	return err
}

func (ur *stateprovinceRepository) Update(c context.Context, stateprovince *domain.StateProvince) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": stateprovince.ID}
	update := bson.M{
		"$set": stateprovince,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *stateprovinceRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *stateprovinceRepository) Fetch(c context.Context) ([]domain.StateProvince, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var stateprovinces []domain.StateProvince

	err = cursor.All(c, &stateprovinces)
	if stateprovinces == nil {
		return []domain.StateProvince{}, err
	}

	return stateprovinces, err
}

func (tr *stateprovinceRepository) FetchByID(c context.Context, stateprovinceID string) (domain.StateProvince, error) {
	collection := tr.database.Collection(tr.collection)

	var stateprovince domain.StateProvince

	idHex, err := bson.ObjectIDFromHex(stateprovinceID)
	if err != nil {
		return stateprovince, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&stateprovince)
	return stateprovince, err
}
