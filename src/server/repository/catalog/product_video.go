package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type productvideoRepository struct {
	database   mongo.Database
	collection string
}

func NewProductVideoRepository(db mongo.Database, collection string) domain.ProductVideoRepository {
	return &productvideoRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productvideoRepository) CreateMany(c context.Context, items []domain.ProductVideo) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productvideoRepository) Create(c context.Context, productvideo *domain.ProductVideo) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productvideo)

	return err
}

func (ur *productvideoRepository) Update(c context.Context, productvideo *domain.ProductVideo) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productvideo.ID}
	update := bson.M{
		"$set": productvideo,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productvideoRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productvideoRepository) Fetch(c context.Context) ([]domain.ProductVideo, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productvideos []domain.ProductVideo

	err = cursor.All(c, &productvideos)
	if productvideos == nil {
		return []domain.ProductVideo{}, err
	}

	return productvideos, err
}

func (tr *productvideoRepository) FetchByID(c context.Context, productvideoID string) (domain.ProductVideo, error) {
	collection := tr.database.Collection(tr.collection)

	var productvideo domain.ProductVideo

	idHex, err := bson.ObjectIDFromHex(productvideoID)
	if err != nil {
		return productvideo, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productvideo)
	return productvideo, err
}
