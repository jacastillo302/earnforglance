package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type displaydefaultmenuitemsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewDisplayDefaultMenuItemSettingsRepository(db mongo.Database, collection string) domain.DisplayDefaultMenuItemSettingsRepository {
	return &displaydefaultmenuitemsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *displaydefaultmenuitemsettingsRepository) CreateMany(c context.Context, items []domain.DisplayDefaultMenuItemSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *displaydefaultmenuitemsettingsRepository) Create(c context.Context, displaydefaultmenuitemsettings *domain.DisplayDefaultMenuItemSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, displaydefaultmenuitemsettings)

	return err
}

func (ur *displaydefaultmenuitemsettingsRepository) Update(c context.Context, displaydefaultmenuitemsettings *domain.DisplayDefaultMenuItemSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": displaydefaultmenuitemsettings.ID}
	update := bson.M{
		"$set": displaydefaultmenuitemsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *displaydefaultmenuitemsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *displaydefaultmenuitemsettingsRepository) Fetch(c context.Context) ([]domain.DisplayDefaultMenuItemSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var displaydefaultmenuitemsettingss []domain.DisplayDefaultMenuItemSettings

	err = cursor.All(c, &displaydefaultmenuitemsettingss)
	if displaydefaultmenuitemsettingss == nil {
		return []domain.DisplayDefaultMenuItemSettings{}, err
	}

	return displaydefaultmenuitemsettingss, err
}

func (tr *displaydefaultmenuitemsettingsRepository) FetchByID(c context.Context, displaydefaultmenuitemsettingsID string) (domain.DisplayDefaultMenuItemSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var displaydefaultmenuitemsettings domain.DisplayDefaultMenuItemSettings

	idHex, err := primitive.ObjectIDFromHex(displaydefaultmenuitemsettingsID)
	if err != nil {
		return displaydefaultmenuitemsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&displaydefaultmenuitemsettings)
	return displaydefaultmenuitemsettings, err
}
