package repository

import (
	"context"

	domain "earnforglance/server/domain/attributes"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PermisionRecordAttributeValueRepository struct {
	database   mongo.Database
	collection string
}

func NewPermisionRecordAttributeValueRepository(db mongo.Database, collection string) domain.PermisionRecordAttributeValueRepository {
	return &PermisionRecordAttributeValueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *PermisionRecordAttributeValueRepository) CreateMany(c context.Context, items []domain.PermisionRecordAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *PermisionRecordAttributeValueRepository) Create(c context.Context, PermisionRecordAttributeValue *domain.PermisionRecordAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, PermisionRecordAttributeValue)

	return err
}

func (ur *PermisionRecordAttributeValueRepository) Update(c context.Context, PermisionRecordAttributeValue *domain.PermisionRecordAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": PermisionRecordAttributeValue.ID}
	update := bson.M{
		"$set": PermisionRecordAttributeValue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *PermisionRecordAttributeValueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *PermisionRecordAttributeValueRepository) Fetch(c context.Context) ([]domain.PermisionRecordAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var PermisionRecordAttributeValues []domain.PermisionRecordAttributeValue

	err = cursor.All(c, &PermisionRecordAttributeValues)
	if PermisionRecordAttributeValues == nil {
		return []domain.PermisionRecordAttributeValue{}, err
	}

	return PermisionRecordAttributeValues, err
}

func (tr *PermisionRecordAttributeValueRepository) FetchByID(c context.Context, PermisionRecordAttributeValueID string) (domain.PermisionRecordAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var PermisionRecordAttributeValue domain.PermisionRecordAttributeValue

	idHex, err := bson.ObjectIDFromHex(PermisionRecordAttributeValueID)
	if err != nil {
		return PermisionRecordAttributeValue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&PermisionRecordAttributeValue)
	return PermisionRecordAttributeValue, err
}
