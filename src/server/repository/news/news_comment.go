package repository

import (
	"context"

	domain "earnforglance/server/domain/news"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newscommentRepository struct {
	database   mongo.Database
	collection string
}

func NewNewsCommentRepository(db mongo.Database, collection string) domain.NewsCommentRepository {
	return &newscommentRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *newscommentRepository) CreateMany(c context.Context, items []domain.NewsComment) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *newscommentRepository) Create(c context.Context, newscomment *domain.NewsComment) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, newscomment)

	return err
}

func (ur *newscommentRepository) Update(c context.Context, newscomment *domain.NewsComment) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": newscomment.ID}
	update := bson.M{
		"$set": newscomment,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *newscommentRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *newscommentRepository) Fetch(c context.Context) ([]domain.NewsComment, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var newscomments []domain.NewsComment

	err = cursor.All(c, &newscomments)
	if newscomments == nil {
		return []domain.NewsComment{}, err
	}

	return newscomments, err
}

func (tr *newscommentRepository) FetchByID(c context.Context, newscommentID string) (domain.NewsComment, error) {
	collection := tr.database.Collection(tr.collection)

	var newscomment domain.NewsComment

	idHex, err := primitive.ObjectIDFromHex(newscommentID)
	if err != nil {
		return newscomment, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&newscomment)
	return newscomment, err
}
