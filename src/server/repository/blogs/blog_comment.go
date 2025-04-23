package repository

import (
	"context"

	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type blogcommentRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogCommentRepository(db mongo.Database, collection string) domain.BlogCommentRepository {
	return &blogcommentRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *blogcommentRepository) CreateMany(c context.Context, items []domain.BlogComment) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *blogcommentRepository) Create(c context.Context, blogcomment *domain.BlogComment) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, blogcomment)

	return err
}

func (ur *blogcommentRepository) Update(c context.Context, blogcomment *domain.BlogComment) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": blogcomment.ID}
	update := bson.M{
		"$set": blogcomment,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *blogcommentRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *blogcommentRepository) Fetch(c context.Context) ([]domain.BlogComment, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var blogcomments []domain.BlogComment

	err = cursor.All(c, &blogcomments)
	if blogcomments == nil {
		return []domain.BlogComment{}, err
	}

	return blogcomments, err
}

func (tr *blogcommentRepository) FetchByID(c context.Context, blogcommentID string) (domain.BlogComment, error) {
	collection := tr.database.Collection(tr.collection)

	var blogcomment domain.BlogComment

	idHex, err := bson.ObjectIDFromHex(blogcommentID)
	if err != nil {
		return blogcomment, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blogcomment)
	return blogcomment, err
}
