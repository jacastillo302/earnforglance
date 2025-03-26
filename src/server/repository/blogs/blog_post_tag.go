package repository

import (
	"context"

	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogposttagRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogPostTagRepository(db mongo.Database, collection string) domain.BlogPostTagRepository {
	return &blogposttagRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *blogposttagRepository) CreateMany(c context.Context, items []domain.BlogPostTag) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *blogposttagRepository) Create(c context.Context, blogposttag *domain.BlogPostTag) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, blogposttag)

	return err
}

func (ur *blogposttagRepository) Update(c context.Context, blogposttag *domain.BlogPostTag) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": blogposttag.ID}
	update := bson.M{
		"$set": blogposttag,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *blogposttagRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *blogposttagRepository) Fetch(c context.Context) ([]domain.BlogPostTag, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var blogposttags []domain.BlogPostTag

	err = cursor.All(c, &blogposttags)
	if blogposttags == nil {
		return []domain.BlogPostTag{}, err
	}

	return blogposttags, err
}

func (tr *blogposttagRepository) FetchByID(c context.Context, blogposttagID string) (domain.BlogPostTag, error) {
	collection := tr.database.Collection(tr.collection)

	var blogposttag domain.BlogPostTag

	idHex, err := primitive.ObjectIDFromHex(blogposttagID)
	if err != nil {
		return blogposttag, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blogposttag)
	return blogposttag, err
}
