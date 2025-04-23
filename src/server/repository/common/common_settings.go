package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type commonsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewCommonSettingsRepository(db mongo.Database, collection string) domain.CommonSettingsRepository {
	return &commonsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *commonsettingsRepository) CreateMany(c context.Context, items []domain.CommonSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *commonsettingsRepository) Create(c context.Context, commonsettings *domain.CommonSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, commonsettings)

	return err
}

func (ur *commonsettingsRepository) Update(c context.Context, commonsettings *domain.CommonSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": commonsettings.ID}
	update := bson.M{
		"$set": commonsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *commonsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *commonsettingsRepository) Fetch(c context.Context) ([]domain.CommonSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var commonsettingss []domain.CommonSettings

	err = cursor.All(c, &commonsettingss)
	if commonsettingss == nil {
		return []domain.CommonSettings{}, err
	}

	return commonsettingss, err
}

func (tr *commonsettingsRepository) FetchByID(c context.Context, commonsettingsID string) (domain.CommonSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var commonsettings domain.CommonSettings

	idHex, err := bson.ObjectIDFromHex(commonsettingsID)
	if err != nil {
		return commonsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&commonsettings)
	return commonsettings, err
}
