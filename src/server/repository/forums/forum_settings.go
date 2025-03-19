package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type forumsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewForumSettingsRepository(db mongo.Database, collection string) domain.ForumSettingsRepository {
	return &forumsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumsettingsRepository) Create(c context.Context, forumsettings *domain.ForumSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forumsettings)

	return err
}

func (ur *forumsettingsRepository) Update(c context.Context, forumsettings *domain.ForumSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumsettings.ID}
	update := bson.M{
		"$set": forumsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *forumsettingsRepository) Delete(c context.Context, forumsettings *domain.ForumSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumsettings.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *forumsettingsRepository) Fetch(c context.Context) ([]domain.ForumSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forumsettingss []domain.ForumSettings

	err = cursor.All(c, &forumsettingss)
	if forumsettingss == nil {
		return []domain.ForumSettings{}, err
	}

	return forumsettingss, err
}

func (tr *forumsettingsRepository) FetchByID(c context.Context, forumsettingsID string) (domain.ForumSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var forumsettings domain.ForumSettings

	idHex, err := primitive.ObjectIDFromHex(forumsettingsID)
	if err != nil {
		return forumsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forumsettings)
	return forumsettings, err
}
