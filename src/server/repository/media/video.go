package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type videoRepository struct {
	database   mongo.Database
	collection string
}

func NewVideoRepository(db mongo.Database, collection string) domain.VideoRepository {
	return &videoRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *videoRepository) Create(c context.Context, video *domain.Video) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, video)

	return err
}

func (ur *videoRepository) Update(c context.Context, video *domain.Video) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": video.ID}
	update := bson.M{
		"$set": video,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *videoRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *videoRepository) Fetch(c context.Context) ([]domain.Video, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var videos []domain.Video

	err = cursor.All(c, &videos)
	if videos == nil {
		return []domain.Video{}, err
	}

	return videos, err
}

func (tr *videoRepository) FetchByID(c context.Context, videoID string) (domain.Video, error) {
	collection := tr.database.Collection(tr.collection)

	var video domain.Video

	idHex, err := primitive.ObjectIDFromHex(videoID)
	if err != nil {
		return video, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&video)
	return video, err
}
