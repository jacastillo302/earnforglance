package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type predefinedproductattributevalueRepository struct {
	database   mongo.Database
	collection string
}

func NewPredefinedProductAttributeValueRepository(db mongo.Database, collection string) domain.PredefinedProductAttributeValueRepository {
	return &predefinedproductattributevalueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *predefinedproductattributevalueRepository) Create(c context.Context, predefinedproductattributevalue *domain.PredefinedProductAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, predefinedproductattributevalue)

	return err
}

func (ur *predefinedproductattributevalueRepository) Update(c context.Context, predefinedproductattributevalue *domain.PredefinedProductAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": predefinedproductattributevalue.ID}
	update := bson.M{
		"$set": predefinedproductattributevalue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *predefinedproductattributevalueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *predefinedproductattributevalueRepository) Fetch(c context.Context) ([]domain.PredefinedProductAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var predefinedproductattributevalues []domain.PredefinedProductAttributeValue

	err = cursor.All(c, &predefinedproductattributevalues)
	if predefinedproductattributevalues == nil {
		return []domain.PredefinedProductAttributeValue{}, err
	}

	return predefinedproductattributevalues, err
}

func (tr *predefinedproductattributevalueRepository) FetchByID(c context.Context, predefinedproductattributevalueID string) (domain.PredefinedProductAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var predefinedproductattributevalue domain.PredefinedProductAttributeValue

	idHex, err := primitive.ObjectIDFromHex(predefinedproductattributevalueID)
	if err != nil {
		return predefinedproductattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&predefinedproductattributevalue)
	return predefinedproductattributevalue, err
}
