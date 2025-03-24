package repository

import (
	"context"

	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogpostRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogPostRepository(db mongo.Database, collection string) domain.BlogPostRepository {
	return &blogpostRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *blogpostRepository) CreateMany(c context.Context, items []domain.BlogPost) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *blogpostRepository) Create(c context.Context, blogpost *domain.BlogPost) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, blogpost)

	return err
}

func (ur *blogpostRepository) Update(c context.Context, blogpost *domain.BlogPost) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": blogpost.ID}
	update := bson.M{
		"$set": blogpost,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *blogpostRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *blogpostRepository) Fetch(c context.Context) ([]domain.BlogPost, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var blogposts []domain.BlogPost

	err = cursor.All(c, &blogposts)
	if blogposts == nil {
		return []domain.BlogPost{}, err
	}

	return blogposts, err
}

func (tr *blogpostRepository) FetchByID(c context.Context, blogpostID string) (domain.BlogPost, error) {
	collection := tr.database.Collection(tr.collection)

	var blogpost domain.BlogPost

	idHex, err := primitive.ObjectIDFromHex(blogpostID)
	if err != nil {
		return blogpost, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blogpost)
	return blogpost, err
}
