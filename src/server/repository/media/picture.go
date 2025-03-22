package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pictureRepository struct {
	database   mongo.Database
	collection string
}

func NewPictureRepository(db mongo.Database, collection string) domain.PictureRepository {
	return &pictureRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *pictureRepository) Create(c context.Context, picture *domain.Picture) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, picture)

	return err
}

func (ur *pictureRepository) Update(c context.Context, picture *domain.Picture) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": picture.ID}
	update := bson.M{
		"$set": picture,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *pictureRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *pictureRepository) Fetch(c context.Context) ([]domain.Picture, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var pictures []domain.Picture

	err = cursor.All(c, &pictures)
	if pictures == nil {
		return []domain.Picture{}, err
	}

	return pictures, err
}

func (tr *pictureRepository) FetchByID(c context.Context, pictureID string) (domain.Picture, error) {
	collection := tr.database.Collection(tr.collection)

	var picture domain.Picture

	idHex, err := primitive.ObjectIDFromHex(pictureID)
	if err != nil {
		return picture, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&picture)
	return picture, err
}
