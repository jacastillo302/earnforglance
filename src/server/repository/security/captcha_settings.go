package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type captchasettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewCaptchaSettingsRepository(db mongo.Database, collection string) domain.CaptchaSettingsRepository {
	return &captchasettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *captchasettingsRepository) Create(c context.Context, captchasettings *domain.CaptchaSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, captchasettings)

	return err
}

func (ur *captchasettingsRepository) Update(c context.Context, captchasettings *domain.CaptchaSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": captchasettings.ID}
	update := bson.M{
		"$set": captchasettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *captchasettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *captchasettingsRepository) Fetch(c context.Context) ([]domain.CaptchaSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var captchasettingss []domain.CaptchaSettings

	err = cursor.All(c, &captchasettingss)
	if captchasettingss == nil {
		return []domain.CaptchaSettings{}, err
	}

	return captchasettingss, err
}

func (tr *captchasettingsRepository) FetchByID(c context.Context, captchasettingsID string) (domain.CaptchaSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var captchasettings domain.CaptchaSettings

	idHex, err := primitive.ObjectIDFromHex(captchasettingsID)
	if err != nil {
		return captchasettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&captchasettings)
	return captchasettings, err
}
