package repository

import (
	"context"

	domain "earnforglance/server/domain/attributes"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type baseattributevalueRepository struct {
	database   mongo.Database
	collection string
}

func NewBaseAttributeValueRepository(db mongo.Database, collection string) domain.BaseAttributeValueRepository {
	return &baseattributevalueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *baseattributevalueRepository) CreateMany(c context.Context, items []domain.BaseAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *baseattributevalueRepository) Create(c context.Context, baseattributevalue *domain.BaseAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, baseattributevalue)

	return err
}

func (ur *baseattributevalueRepository) Update(c context.Context, baseattributevalue *domain.BaseAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": baseattributevalue.ID}
	update := bson.M{
		"$set": baseattributevalue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *baseattributevalueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *baseattributevalueRepository) Fetch(c context.Context) ([]domain.BaseAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var baseattributevalues []domain.BaseAttributeValue

	err = cursor.All(c, &baseattributevalues)
	if baseattributevalues == nil {
		return []domain.BaseAttributeValue{}, err
	}

	return baseattributevalues, err
}

func (tr *baseattributevalueRepository) FetchByID(c context.Context, baseattributevalueID string) (domain.BaseAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var baseattributevalue domain.BaseAttributeValue

	idHex, err := primitive.ObjectIDFromHex(baseattributevalueID)
	if err != nil {
		return baseattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&baseattributevalue)
	return baseattributevalue, err
}
