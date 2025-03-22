package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type measuresettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMeasureSettingsRepository(db mongo.Database, collection string) domain.MeasureSettingsRepository {
	return &measuresettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *measuresettingsRepository) Create(c context.Context, measuresettings *domain.MeasureSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, measuresettings)

	return err
}

func (ur *measuresettingsRepository) Update(c context.Context, measuresettings *domain.MeasureSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": measuresettings.ID}
	update := bson.M{
		"$set": measuresettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *measuresettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *measuresettingsRepository) Fetch(c context.Context) ([]domain.MeasureSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var measuresettingss []domain.MeasureSettings

	err = cursor.All(c, &measuresettingss)
	if measuresettingss == nil {
		return []domain.MeasureSettings{}, err
	}

	return measuresettingss, err
}

func (tr *measuresettingsRepository) FetchByID(c context.Context, measuresettingsID string) (domain.MeasureSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var measuresettings domain.MeasureSettings

	idHex, err := primitive.ObjectIDFromHex(measuresettingsID)
	if err != nil {
		return measuresettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&measuresettings)
	return measuresettings, err
}
