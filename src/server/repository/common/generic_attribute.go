package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type genericattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewGenericAttributeRepository(db mongo.Database, collection string) domain.GenericAttributeRepository {
	return &genericattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *genericattributeRepository) CreateMany(c context.Context, items []domain.GenericAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *genericattributeRepository) Create(c context.Context, genericattribute *domain.GenericAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, genericattribute)

	return err
}

func (ur *genericattributeRepository) Update(c context.Context, genericattribute *domain.GenericAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": genericattribute.ID}
	update := bson.M{
		"$set": genericattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *genericattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *genericattributeRepository) Fetch(c context.Context) ([]domain.GenericAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var genericattributes []domain.GenericAttribute

	err = cursor.All(c, &genericattributes)
	if genericattributes == nil {
		return []domain.GenericAttribute{}, err
	}

	return genericattributes, err
}

func (tr *genericattributeRepository) FetchByID(c context.Context, genericattributeID string) (domain.GenericAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var genericattribute domain.GenericAttribute

	idHex, err := primitive.ObjectIDFromHex(genericattributeID)
	if err != nil {
		return genericattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&genericattribute)
	return genericattribute, err
}
