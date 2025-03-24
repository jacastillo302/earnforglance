package repository

import (
	"context"

	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogSettingsRepository(db mongo.Database, collection string) domain.BlogSettingsRepository {
	return &blogsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *blogsettingsRepository) CreateMany(c context.Context, items []domain.BlogSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *blogsettingsRepository) Create(c context.Context, blogsettings *domain.BlogSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, blogsettings)

	return err
}

func (ur *blogsettingsRepository) Update(c context.Context, blogsettings *domain.BlogSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": blogsettings.ID}
	update := bson.M{
		"$set": blogsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *blogsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *blogsettingsRepository) Fetch(c context.Context) ([]domain.BlogSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var blogsettingss []domain.BlogSettings

	err = cursor.All(c, &blogsettingss)
	if blogsettingss == nil {
		return []domain.BlogSettings{}, err
	}

	return blogsettingss, err
}

func (tr *blogsettingsRepository) FetchByID(c context.Context, blogsettingsID string) (domain.BlogSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var blogsettings domain.BlogSettings

	idHex, err := primitive.ObjectIDFromHex(blogsettingsID)
	if err != nil {
		return blogsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blogsettings)
	return blogsettings, err
}
