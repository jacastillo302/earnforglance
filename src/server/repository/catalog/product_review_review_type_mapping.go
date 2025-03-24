package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productreviewreviewtypemappingRepository struct {
	database   mongo.Database
	collection string
}

func NewProductReviewReviewTypeMappingRepository(db mongo.Database, collection string) domain.ProductReviewReviewTypeMappingRepository {
	return &productreviewreviewtypemappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productreviewreviewtypemappingRepository) CreateMany(c context.Context, items []domain.ProductReviewReviewTypeMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productreviewreviewtypemappingRepository) Create(c context.Context, productreviewreviewtypemapping *domain.ProductReviewReviewTypeMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productreviewreviewtypemapping)

	return err
}

func (ur *productreviewreviewtypemappingRepository) Update(c context.Context, productreviewreviewtypemapping *domain.ProductReviewReviewTypeMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productreviewreviewtypemapping.ID}
	update := bson.M{
		"$set": productreviewreviewtypemapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productreviewreviewtypemappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productreviewreviewtypemappingRepository) Fetch(c context.Context) ([]domain.ProductReviewReviewTypeMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productreviewreviewtypemappings []domain.ProductReviewReviewTypeMapping

	err = cursor.All(c, &productreviewreviewtypemappings)
	if productreviewreviewtypemappings == nil {
		return []domain.ProductReviewReviewTypeMapping{}, err
	}

	return productreviewreviewtypemappings, err
}

func (tr *productreviewreviewtypemappingRepository) FetchByID(c context.Context, productreviewreviewtypemappingID string) (domain.ProductReviewReviewTypeMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var productreviewreviewtypemapping domain.ProductReviewReviewTypeMapping

	idHex, err := primitive.ObjectIDFromHex(productreviewreviewtypemappingID)
	if err != nil {
		return productreviewreviewtypemapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productreviewreviewtypemapping)
	return productreviewreviewtypemapping, err
}
