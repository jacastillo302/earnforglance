package repository

import (
	"context"

	domain "earnforglance/server/domain/configuration"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type settingRepository struct {
	database   mongo.Database
	collection string
}

func NewSettingRepository(db mongo.Database, collection string) domain.SettingRepository {
	return &settingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *settingRepository) CreateMany(c context.Context, items []domain.Setting) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *settingRepository) Create(c context.Context, setting *domain.Setting) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, setting)

	return err
}

func (ur *settingRepository) Update(c context.Context, setting *domain.Setting) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": setting.ID}
	update := bson.M{
		"$set": setting,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *settingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *settingRepository) Fetch(c context.Context) ([]domain.Setting, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var settings []domain.Setting

	err = cursor.All(c, &settings)
	if settings == nil {
		return []domain.Setting{}, err
	}

	return settings, err
}

func (tr *settingRepository) FetchByID(c context.Context, settingID string) (domain.Setting, error) {
	collection := tr.database.Collection(tr.collection)

	var setting domain.Setting

	idHex, err := primitive.ObjectIDFromHex(settingID)
	if err != nil {
		return setting, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&setting)
	return setting, err
}
