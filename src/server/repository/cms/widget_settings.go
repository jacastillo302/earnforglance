package repository

import (
	"context"

	domain "earnforglance/server/domain/cms"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type widgetsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewWidgetSettingsRepository(db mongo.Database, collection string) domain.WidgetSettingsRepository {
	return &widgetsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *widgetsettingsRepository) CreateMany(c context.Context, items []domain.WidgetSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *widgetsettingsRepository) Create(c context.Context, widgetsettings *domain.WidgetSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, widgetsettings)

	return err
}

func (ur *widgetsettingsRepository) Update(c context.Context, widgetsettings *domain.WidgetSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": widgetsettings.ID}
	update := bson.M{
		"$set": widgetsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *widgetsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *widgetsettingsRepository) Fetch(c context.Context) ([]domain.WidgetSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var widgetsettingss []domain.WidgetSettings

	err = cursor.All(c, &widgetsettingss)
	if widgetsettingss == nil {
		return []domain.WidgetSettings{}, err
	}

	return widgetsettingss, err
}

func (tr *widgetsettingsRepository) FetchByID(c context.Context, widgetsettingsID string) (domain.WidgetSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var widgetsettings domain.WidgetSettings

	idHex, err := primitive.ObjectIDFromHex(widgetsettingsID)
	if err != nil {
		return widgetsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&widgetsettings)
	return widgetsettings, err
}
