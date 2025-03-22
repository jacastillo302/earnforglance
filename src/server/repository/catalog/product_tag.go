package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type producttagRepository struct {
	database   mongo.Database
	collection string
}

func NewProductTagRepository(db mongo.Database, collection string) domain.ProductTagRepository {
	return &producttagRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *producttagRepository) Create(c context.Context, producttag *domain.ProductTag) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, producttag)

	return err
}

func (ur *producttagRepository) Update(c context.Context, producttag *domain.ProductTag) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": producttag.ID}
	update := bson.M{
		"$set": producttag,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *producttagRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *producttagRepository) Fetch(c context.Context) ([]domain.ProductTag, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var producttags []domain.ProductTag

	err = cursor.All(c, &producttags)
	if producttags == nil {
		return []domain.ProductTag{}, err
	}

	return producttags, err
}

func (tr *producttagRepository) FetchByID(c context.Context, producttagID string) (domain.ProductTag, error) {
	collection := tr.database.Collection(tr.collection)

	var producttag domain.ProductTag

	idHex, err := primitive.ObjectIDFromHex(producttagID)
	if err != nil {
		return producttag, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&producttag)
	return producttag, err
}
