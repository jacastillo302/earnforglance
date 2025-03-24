package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type picturehashesRepository struct {
	database   mongo.Database
	collection string
}

func NewPictureHashesRepository(db mongo.Database, collection string) domain.PictureHashesRepository {
	return &picturehashesRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *picturehashesRepository) CreateMany(c context.Context, items []domain.PictureHashes) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *picturehashesRepository) Create(c context.Context, picturehashes *domain.PictureHashes) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, picturehashes)

	return err
}

func (ur *picturehashesRepository) Update(c context.Context, picturehashes *domain.PictureHashes) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": picturehashes.PictureID}
	update := bson.M{
		"$set": picturehashes,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *picturehashesRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *picturehashesRepository) Fetch(c context.Context) ([]domain.PictureHashes, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var picturehashess []domain.PictureHashes

	err = cursor.All(c, &picturehashess)
	if picturehashess == nil {
		return []domain.PictureHashes{}, err
	}

	return picturehashess, err
}

func (tr *picturehashesRepository) FetchByID(c context.Context, picturehashesID string) (domain.PictureHashes, error) {
	collection := tr.database.Collection(tr.collection)

	var picturehashes domain.PictureHashes

	idHex, err := primitive.ObjectIDFromHex(picturehashesID)
	if err != nil {
		return picturehashes, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&picturehashes)
	return picturehashes, err
}
