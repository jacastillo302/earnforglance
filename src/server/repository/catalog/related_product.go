package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type relatedProductRepository struct {
	database   mongo.Database
	collection string
}

func NewRelatedProductRepository(db mongo.Database, collection string) domain.RelatedProductRepository {
	return &relatedProductRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *relatedProductRepository) Create(c context.Context, relatedProduct *domain.RelatedProduct) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, relatedProduct)

	return err
}

func (ur *relatedProductRepository) Update(c context.Context, relatedProduct *domain.RelatedProduct) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": relatedProduct.ID}
	update := bson.M{
		"$set": relatedProduct,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *relatedProductRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *relatedProductRepository) Fetch(c context.Context) ([]domain.RelatedProduct, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var relatedProducts []domain.RelatedProduct

	err = cursor.All(c, &relatedProducts)
	if relatedProducts == nil {
		return []domain.RelatedProduct{}, err
	}

	return relatedProducts, err
}

func (tr *relatedProductRepository) FetchByID(c context.Context, relatedProductID string) (domain.RelatedProduct, error) {
	collection := tr.database.Collection(tr.collection)

	var relatedProduct domain.RelatedProduct

	idHex, err := primitive.ObjectIDFromHex(relatedProductID)
	if err != nil {
		return relatedProduct, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&relatedProduct)
	return relatedProduct, err
}
