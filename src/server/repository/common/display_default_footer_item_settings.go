package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type displaydefaultfooteritemsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewDisplayDefaultFooterItemSettingsRepository(db mongo.Database, collection string) domain.DisplayDefaultFooterItemSettingsRepository {
	return &displaydefaultfooteritemsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *displaydefaultfooteritemsettingsRepository) CreateMany(c context.Context, items []domain.DisplayDefaultFooterItemSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *displaydefaultfooteritemsettingsRepository) Create(c context.Context, displaydefaultfooteritemsettings *domain.DisplayDefaultFooterItemSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, displaydefaultfooteritemsettings)

	return err
}

func (ur *displaydefaultfooteritemsettingsRepository) Update(c context.Context, displaydefaultfooteritemsettings *domain.DisplayDefaultFooterItemSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": displaydefaultfooteritemsettings.ID}
	update := bson.M{
		"$set": displaydefaultfooteritemsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *displaydefaultfooteritemsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *displaydefaultfooteritemsettingsRepository) Fetch(c context.Context) ([]domain.DisplayDefaultFooterItemSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var displaydefaultfooteritemsettingss []domain.DisplayDefaultFooterItemSettings

	err = cursor.All(c, &displaydefaultfooteritemsettingss)
	if displaydefaultfooteritemsettingss == nil {
		return []domain.DisplayDefaultFooterItemSettings{}, err
	}

	return displaydefaultfooteritemsettingss, err
}

func (tr *displaydefaultfooteritemsettingsRepository) FetchByID(c context.Context, displaydefaultfooteritemsettingsID string) (domain.DisplayDefaultFooterItemSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var displaydefaultfooteritemsettings domain.DisplayDefaultFooterItemSettings

	idHex, err := bson.ObjectIDFromHex(displaydefaultfooteritemsettingsID)
	if err != nil {
		return displaydefaultfooteritemsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&displaydefaultfooteritemsettings)
	return displaydefaultfooteritemsettings, err
}
