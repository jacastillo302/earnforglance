package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productattributecombinationRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAttributeCombinationRepository(db mongo.Database, collection string) domain.ProductAttributeCombinationRepository {
	return &productattributecombinationRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productattributecombinationRepository) Create(c context.Context, productattributecombination *domain.ProductAttributeCombination) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productattributecombination)

	return err
}

func (ur *productattributecombinationRepository) Update(c context.Context, productattributecombination *domain.ProductAttributeCombination) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productattributecombination.ID}
	update := bson.M{
		"$set": productattributecombination,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productattributecombinationRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productattributecombinationRepository) Fetch(c context.Context) ([]domain.ProductAttributeCombination, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productattributecombinations []domain.ProductAttributeCombination

	err = cursor.All(c, &productattributecombinations)
	if productattributecombinations == nil {
		return []domain.ProductAttributeCombination{}, err
	}

	return productattributecombinations, err
}

func (tr *productattributecombinationRepository) FetchByID(c context.Context, productattributecombinationID string) (domain.ProductAttributeCombination, error) {
	collection := tr.database.Collection(tr.collection)

	var productattributecombination domain.ProductAttributeCombination

	idHex, err := primitive.ObjectIDFromHex(productattributecombinationID)
	if err != nil {
		return productattributecombination, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productattributecombination)
	return productattributecombination, err
}
