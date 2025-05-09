package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type productRepository struct {
	database   mongo.Database
	collection string
}

func NewProductRepository(db mongo.Database, collection string) domain.ProductRepository {
	return &productRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productRepository) CreateMany(c context.Context, items []domain.Product) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productRepository) Create(c context.Context, product *domain.Product) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, product)

	return err
}

func (ur *productRepository) Update(c context.Context, product *domain.Product) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": product.ID}
	update := bson.M{
		"$set": product,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productRepository) Fetch(c context.Context) ([]domain.Product, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var products []domain.Product

	err = cursor.All(c, &products)
	if products == nil {
		return []domain.Product{}, err
	}

	return products, err
}

func (tr *productRepository) FetchByID(c context.Context, productID string) (domain.Product, error) {
	collection := tr.database.Collection(tr.collection)

	var product domain.Product

	idHex, err := bson.ObjectIDFromHex(productID)
	if err != nil {
		return product, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&product)
	return product, err
}
