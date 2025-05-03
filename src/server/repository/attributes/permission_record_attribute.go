package repository

import (
	"context"

	domain "earnforglance/server/domain/attributes"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type customerattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewPermisionRecordAttributeRepository(db mongo.Database, collection string) domain.PermisionRecordAttributeRepository {
	return &customerattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customerattributeRepository) CreateMany(c context.Context, items []domain.PermisionRecordAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *customerattributeRepository) Create(c context.Context, customerattribute *domain.PermisionRecordAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customerattribute)

	return err
}

func (ur *customerattributeRepository) Update(c context.Context, customerattribute *domain.PermisionRecordAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customerattribute.ID}
	update := bson.M{
		"$set": customerattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customerattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *customerattributeRepository) Fetch(c context.Context) ([]domain.PermisionRecordAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customerattributes []domain.PermisionRecordAttribute

	err = cursor.All(c, &customerattributes)
	if customerattributes == nil {
		return []domain.PermisionRecordAttribute{}, err
	}

	return customerattributes, err
}

func (tr *customerattributeRepository) FetchByID(c context.Context, customerattributeID string) (domain.PermisionRecordAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var customerattribute domain.PermisionRecordAttribute

	idHex, err := bson.ObjectIDFromHex(customerattributeID)
	if err != nil {
		return customerattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customerattribute)
	return customerattribute, err
}
