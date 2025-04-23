package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type productattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAttributeRepository(db mongo.Database, collection string) domain.ProductAttributeRepository {
	return &productattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productattributeRepository) CreateMany(c context.Context, items []domain.ProductAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productattributeRepository) Create(c context.Context, productattribute *domain.ProductAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productattribute)

	return err
}

func (ur *productattributeRepository) Update(c context.Context, productattribute *domain.ProductAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productattribute.ID}
	update := bson.M{
		"$set": productattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *productattributeRepository) Fetch(c context.Context) ([]domain.ProductAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productattributes []domain.ProductAttribute

	err = cursor.All(c, &productattributes)
	if productattributes == nil {
		return []domain.ProductAttribute{}, err
	}

	return productattributes, err
}

func (tr *productattributeRepository) FetchByID(c context.Context, productattributeID string) (domain.ProductAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var productattribute domain.ProductAttribute

	idHex, err := bson.ObjectIDFromHex(productattributeID)
	if err != nil {
		return productattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productattribute)
	return productattribute, err
}
