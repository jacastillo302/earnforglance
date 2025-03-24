package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type securitysettingRepository struct {
	database   mongo.Database
	collection string
}

func NewSecuritySettingsRepository(db mongo.Database, collection string) domain.SecuritySettingsRepository {
	return &securitysettingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *securitysettingRepository) CreateMany(c context.Context, items []domain.SecuritySettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *securitysettingRepository) Create(c context.Context, securitysetting *domain.SecuritySettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, securitysetting)

	return err
}

func (ur *securitysettingRepository) Update(c context.Context, securitysetting *domain.SecuritySettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": securitysetting.ID}
	update := bson.M{
		"$set": securitysetting,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *securitysettingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *securitysettingRepository) Fetch(c context.Context) ([]domain.SecuritySettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var securitysettings []domain.SecuritySettings

	err = cursor.All(c, &securitysettings)
	if securitysettings == nil {
		return []domain.SecuritySettings{}, err
	}

	return securitysettings, err
}

func (tr *securitysettingRepository) FetchByID(c context.Context, securitysettingID string) (domain.SecuritySettings, error) {
	collection := tr.database.Collection(tr.collection)

	var securitysetting domain.SecuritySettings

	idHex, err := primitive.ObjectIDFromHex(securitysettingID)
	if err != nil {
		return securitysetting, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&securitysetting)
	return securitysetting, err
}
