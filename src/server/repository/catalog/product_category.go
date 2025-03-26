package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productcategoryRepository struct {
	database   mongo.Database
	collection string
}

func NewProductCategoryRepository(db mongo.Database, collection string) domain.ProductCategoryRepository {
	return &productcategoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productcategoryRepository) CreateMany(c context.Context, items []domain.ProductCategory) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productcategoryRepository) Create(c context.Context, productcategory *domain.ProductCategory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productcategory)

	return err
}

func (ur *productcategoryRepository) Update(c context.Context, productcategory *domain.ProductCategory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productcategory.ID}
	update := bson.M{
		"$set": productcategory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *productcategoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productcategoryRepository) Fetch(c context.Context) ([]domain.ProductCategory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productcategories []domain.ProductCategory

	err = cursor.All(c, &productcategories)
	if productcategories == nil {
		return []domain.ProductCategory{}, err
	}

	return productcategories, err
}

func (tr *productcategoryRepository) FetchByID(c context.Context, productcategoryID string) (domain.ProductCategory, error) {
	collection := tr.database.Collection(tr.collection)

	var productcategory domain.ProductCategory

	idHex, err := primitive.ObjectIDFromHex(productcategoryID)
	if err != nil {
		return productcategory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productcategory)
	return productcategory, err
}
