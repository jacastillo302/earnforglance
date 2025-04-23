package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type productreviewhelpfulnessRepository struct {
	database   mongo.Database
	collection string
}

func NewProductReviewHelpfulnessRepository(db mongo.Database, collection string) domain.ProductReviewHelpfulnessRepository {
	return &productreviewhelpfulnessRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productreviewhelpfulnessRepository) CreateMany(c context.Context, items []domain.ProductReviewHelpfulness) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productreviewhelpfulnessRepository) Create(c context.Context, productreviewhelpfulness *domain.ProductReviewHelpfulness) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productreviewhelpfulness)

	return err
}

func (ur *productreviewhelpfulnessRepository) Update(c context.Context, productreviewhelpfulness *domain.ProductReviewHelpfulness) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productreviewhelpfulness.ID}
	update := bson.M{
		"$set": productreviewhelpfulness,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productreviewhelpfulnessRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productreviewhelpfulnessRepository) Fetch(c context.Context) ([]domain.ProductReviewHelpfulness, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productreviewhelpfulnesses []domain.ProductReviewHelpfulness

	err = cursor.All(c, &productreviewhelpfulnesses)
	if productreviewhelpfulnesses == nil {
		return []domain.ProductReviewHelpfulness{}, err
	}

	return productreviewhelpfulnesses, err
}

func (tr *productreviewhelpfulnessRepository) FetchByID(c context.Context, productreviewhelpfulnessID string) (domain.ProductReviewHelpfulness, error) {
	collection := tr.database.Collection(tr.collection)

	var productreviewhelpfulness domain.ProductReviewHelpfulness

	idHex, err := bson.ObjectIDFromHex(productreviewhelpfulnessID)
	if err != nil {
		return productreviewhelpfulness, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productreviewhelpfulness)
	return productreviewhelpfulness, err
}
