package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type forumRepository struct {
	database   mongo.Database
	collection string
}

func NewForumRepository(db mongo.Database, collection string) domain.ForumRepository {
	return &forumRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumRepository) Create(c context.Context, forum *domain.Forum) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forum)

	return err
}

func (ur *forumRepository) Update(c context.Context, forum *domain.Forum) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forum.ID}
	update := bson.M{
		"$set": forum,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *forumRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *forumRepository) Fetch(c context.Context) ([]domain.Forum, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forums []domain.Forum

	err = cursor.All(c, &forums)
	if forums == nil {
		return []domain.Forum{}, err
	}

	return forums, err
}

func (tr *forumRepository) FetchByID(c context.Context, forumID string) (domain.Forum, error) {
	collection := tr.database.Collection(tr.collection)

	var forum domain.Forum

	idHex, err := primitive.ObjectIDFromHex(forumID)
	if err != nil {
		return forum, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forum)
	return forum, err
}
