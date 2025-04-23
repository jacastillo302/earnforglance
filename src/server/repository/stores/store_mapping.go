package repository

import (
	"context"

	domain "earnforglance/server/domain/stores"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type storemappingRepository struct {
	database   mongo.Database
	collection string
}

func NewStoreMappingRepository(db mongo.Database, collection string) domain.StoreMappingRepository {
	return &storemappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *storemappingRepository) CreateMany(c context.Context, items []domain.StoreMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *storemappingRepository) Create(c context.Context, storemapping *domain.StoreMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, storemapping)

	return err
}

func (ur *storemappingRepository) Update(c context.Context, storemapping *domain.StoreMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": storemapping.ID}
	update := bson.M{
		"$set": storemapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *storemappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *storemappingRepository) Fetch(c context.Context) ([]domain.StoreMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var storemappings []domain.StoreMapping

	err = cursor.All(c, &storemappings)
	if storemappings == nil {
		return []domain.StoreMapping{}, err
	}

	return storemappings, err
}

func (tr *storemappingRepository) FetchByID(c context.Context, storemappingID string) (domain.StoreMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var storemapping domain.StoreMapping

	idHex, err := bson.ObjectIDFromHex(storemappingID)
	if err != nil {
		return storemapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&storemapping)
	return storemapping, err
}
