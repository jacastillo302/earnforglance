package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (ur *productspecificationattributeRepository) CreateMany(c context.Context, items []domain.ProductSpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productspecificationattributeRepository) Create(c context.Context, productspecificationattribute *domain.ProductSpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productspecificationattribute)

	return err
}

func (ur *productspecificationattributeRepository) Update(c context.Context, productspecificationattribute *domain.ProductSpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"product_id": productspecificationattribute.ProductID, "specification_attribute_option_id": productspecificationattribute.SpecificationAttributeOptionID}
	update := bson.M{
		"$set": productspecificationattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productspecificationattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productspecificationattributeRepository) Fetch(c context.Context) ([]domain.ProductSpecificationAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := bson.ObjectIDFromHex(productspecificationattributeID)
	if err != nil {
		return productspecificationattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productspecificationattribute)
	return productspecificationattribute, err
}
