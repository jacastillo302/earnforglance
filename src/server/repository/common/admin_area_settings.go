package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type adminareasettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewAdminAreaSettingsRepository(db mongo.Database, collection string) domain.AdminAreaSettingsRepository {
	return &adminareasettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *adminareasettingsRepository) CreateMany(c context.Context, items []domain.AdminAreaSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *adminareasettingsRepository) Create(c context.Context, adminareasettings *domain.AdminAreaSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, adminareasettings)

	return err
}

func (ur *adminareasettingsRepository) Update(c context.Context, adminareasettings *domain.AdminAreaSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": adminareasettings.ID}
	update := bson.M{
		"$set": adminareasettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *adminareasettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *adminareasettingsRepository) Fetch(c context.Context) ([]domain.AdminAreaSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var adminareasettingss []domain.AdminAreaSettings

	err = cursor.All(c, &adminareasettingss)
	if adminareasettingss == nil {
		return []domain.AdminAreaSettings{}, err
	}

	return adminareasettingss, err
}

func (tr *adminareasettingsRepository) FetchByID(c context.Context, adminareasettingsID string) (domain.AdminAreaSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var adminareasettings domain.AdminAreaSettings

	idHex, err := bson.ObjectIDFromHex(adminareasettingsID)
	if err != nil {
		return adminareasettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&adminareasettings)
	return adminareasettings, err
}
