package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type productattributemappingRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAttributeMappingRepository(db mongo.Database, collection string) domain.ProductAttributeMappingRepository {
	return &productattributemappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productattributemappingRepository) CreateMany(c context.Context, items []domain.ProductAttributeMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productattributemappingRepository) Create(c context.Context, productattributemapping *domain.ProductAttributeMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productattributemapping)

	return err
}

func (ur *productattributemappingRepository) Update(c context.Context, productattributemapping *domain.ProductAttributeMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productattributemapping.ID}
	update := bson.M{
		"$set": productattributemapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productattributemappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productattributemappingRepository) Fetch(c context.Context) ([]domain.ProductAttributeMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productattributemappings []domain.ProductAttributeMapping

	err = cursor.All(c, &productattributemappings)
	if productattributemappings == nil {
		return []domain.ProductAttributeMapping{}, err
	}

	return productattributemappings, err
}

func (tr *productattributemappingRepository) FetchByID(c context.Context, productattributemappingID string) (domain.ProductAttributeMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var productattributemapping domain.ProductAttributeMapping

	idHex, err := bson.ObjectIDFromHex(productattributemappingID)
	if err != nil {
		return productattributemapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productattributemapping)
	return productattributemapping, err
}
