package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productspecificationattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewProductSpecificationAttributeRepository(db mongo.Database, collection string) domain.ProductSpecificationAttributeRepository {
	return &productspecificationattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productspecificationattributeRepository) Create(c context.Context, productspecificationattribute *domain.ProductSpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productspecificationattribute)

	return err
}

func (ur *productspecificationattributeRepository) Update(c context.Context, productspecificationattribute *domain.ProductSpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productspecificationattribute.ID}
	update := bson.M{
		"$set": productspecificationattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productspecificationattributeRepository) Delete(c context.Context, productspecificationattribute *domain.ProductSpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productspecificationattribute.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *productspecificationattributeRepository) Fetch(c context.Context) ([]domain.ProductSpecificationAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productspecificationattributes []domain.ProductSpecificationAttribute

	err = cursor.All(c, &productspecificationattributes)
	if productspecificationattributes == nil {
		return []domain.ProductSpecificationAttribute{}, err
	}

	return productspecificationattributes, err
}

func (tr *productspecificationattributeRepository) FetchByID(c context.Context, productspecificationattributeID string) (domain.ProductSpecificationAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var productspecificationattribute domain.ProductSpecificationAttribute

	idHex, err := primitive.ObjectIDFromHex(productspecificationattributeID)
	if err != nil {
		return productspecificationattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productspecificationattribute)
	return productspecificationattribute, err
}
