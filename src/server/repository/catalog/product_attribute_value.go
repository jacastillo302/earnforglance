package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productattributevalueRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAttributeValueRepository(db mongo.Database, collection string) domain.ProductAttributeValueRepository {
	return &productattributevalueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productattributevalueRepository) CreateMany(c context.Context, items []domain.ProductAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productattributevalueRepository) Create(c context.Context, productattributevalue *domain.ProductAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productattributevalue)

	return err
}

func (ur *productattributevalueRepository) Update(c context.Context, productattributevalue *domain.ProductAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productattributevalue.ID}
	update := bson.M{
		"$set": productattributevalue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *productattributevalueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productattributevalueRepository) Fetch(c context.Context) ([]domain.ProductAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productattributevalues []domain.ProductAttributeValue

	err = cursor.All(c, &productattributevalues)
	if productattributevalues == nil {
		return []domain.ProductAttributeValue{}, err
	}

	return productattributevalues, err
}

func (tr *productattributevalueRepository) FetchByID(c context.Context, productattributevalueID string) (domain.ProductAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var productattributevalue domain.ProductAttributeValue

	idHex, err := primitive.ObjectIDFromHex(productattributevalueID)
	if err != nil {
		return productattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productattributevalue)
	return productattributevalue, err
}
