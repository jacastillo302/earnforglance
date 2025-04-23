package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type productmanufacturerRepository struct {
	database   mongo.Database
	collection string
}

func NewProductManufacturerRepository(db mongo.Database, collection string) domain.ProductManufacturerRepository {
	return &productmanufacturerRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productmanufacturerRepository) CreateMany(c context.Context, items []domain.ProductManufacturer) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productmanufacturerRepository) Create(c context.Context, productmanufacturer *domain.ProductManufacturer) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productmanufacturer)

	return err
}

func (ur *productmanufacturerRepository) Update(c context.Context, productmanufacturer *domain.ProductManufacturer) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productmanufacturer.ID}
	update := bson.M{
		"$set": productmanufacturer,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *productmanufacturerRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productmanufacturerRepository) Fetch(c context.Context) ([]domain.ProductManufacturer, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productmanufacturers []domain.ProductManufacturer

	err = cursor.All(c, &productmanufacturers)
	if productmanufacturers == nil {
		return []domain.ProductManufacturer{}, err
	}

	return productmanufacturers, err
}

func (tr *productmanufacturerRepository) FetchByID(c context.Context, productmanufacturerID string) (domain.ProductManufacturer, error) {
	collection := tr.database.Collection(tr.collection)

	var productmanufacturer domain.ProductManufacturer

	idHex, err := bson.ObjectIDFromHex(productmanufacturerID)
	if err != nil {
		return productmanufacturer, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productmanufacturer)
	return productmanufacturer, err
}
