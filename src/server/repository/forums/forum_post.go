package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type forumpostRepository struct {
	database   mongo.Database
	collection string
}

func NewForumPostRepository(db mongo.Database, collection string) domain.ForumPostRepository {
	return &forumpostRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumpostRepository) CreateMany(c context.Context, items []domain.ForumPost) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *forumpostRepository) Create(c context.Context, forumpost *domain.ForumPost) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forumpost)

	return err
}

func (ur *forumpostRepository) Update(c context.Context, forumpost *domain.ForumPost) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumpost.ID}
	update := bson.M{
		"$set": forumpost,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *forumpostRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *forumpostRepository) Fetch(c context.Context) ([]domain.ForumPost, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forumposts []domain.ForumPost

	err = cursor.All(c, &forumposts)
	if forumposts == nil {
		return []domain.ForumPost{}, err
	}

	return forumposts, err
}

func (tr *forumpostRepository) FetchByID(c context.Context, forumpostID string) (domain.ForumPost, error) {
	collection := tr.database.Collection(tr.collection)

	var forumpost domain.ForumPost

	idHex, err := primitive.ObjectIDFromHex(forumpostID)
	if err != nil {
		return forumpost, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forumpost)
	return forumpost, err
}
