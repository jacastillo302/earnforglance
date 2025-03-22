package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productreviewRepository struct {
	database   mongo.Database
	collection string
}

func NewProductReviewRepository(db mongo.Database, collection string) domain.ProductReviewRepository {
	return &productreviewRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productreviewRepository) Create(c context.Context, productreview *domain.ProductReview) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productreview)

	return err
}

func (ur *productreviewRepository) Update(c context.Context, productreview *domain.ProductReview) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productreview.ID}
	update := bson.M{
		"$set": productreview,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productreviewRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productreviewRepository) Fetch(c context.Context) ([]domain.ProductReview, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productreviews []domain.ProductReview

	err = cursor.All(c, &productreviews)
	if productreviews == nil {
		return []domain.ProductReview{}, err
	}

	return productreviews, err
}

func (tr *productreviewRepository) FetchByID(c context.Context, productreviewID string) (domain.ProductReview, error) {
	collection := tr.database.Collection(tr.collection)

	var productreview domain.ProductReview

	idHex, err := primitive.ObjectIDFromHex(productreviewID)
	if err != nil {
		return productreview, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productreview)
	return productreview, err
}
