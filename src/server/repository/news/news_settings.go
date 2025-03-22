package repository

import (
	"context"

	domain "earnforglance/server/domain/news"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newsSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewNewsSettingsRepository(db mongo.Database, collection string) domain.NewsSettingsRepository {
	return &newsSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *newsSettingsRepository) Create(c context.Context, newsSettings *domain.NewsSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, newsSettings)

	return err
}

func (ur *newsSettingsRepository) Update(c context.Context, newsSettings *domain.NewsSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": newsSettings.ID}
	update := bson.M{
		"$set": newsSettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *newsSettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *newsSettingsRepository) Fetch(c context.Context) ([]domain.NewsSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var newsSettingss []domain.NewsSettings

	err = cursor.All(c, &newsSettingss)
	if newsSettingss == nil {
		return []domain.NewsSettings{}, err
	}

	return newsSettingss, err
}

func (tr *newsSettingsRepository) FetchByID(c context.Context, newsSettingsID string) (domain.NewsSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var newsSettings domain.NewsSettings

	idHex, err := primitive.ObjectIDFromHex(newsSettingsID)
	if err != nil {
		return newsSettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&newsSettings)
	return newsSettings, err
}
