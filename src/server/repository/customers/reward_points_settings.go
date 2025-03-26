package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type rewardpointssettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewRewardPointsSettingsRepository(db mongo.Database, collection string) domain.RewardPointsSettingsRepository {
	return &rewardpointssettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *rewardpointssettingsRepository) CreateMany(c context.Context, items []domain.RewardPointsSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *rewardpointssettingsRepository) Create(c context.Context, rewardpointssettings *domain.RewardPointsSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, rewardpointssettings)

	return err
}

func (ur *rewardpointssettingsRepository) Update(c context.Context, rewardpointssettings *domain.RewardPointsSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": rewardpointssettings.ID}
	update := bson.M{
		"$set": rewardpointssettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *rewardpointssettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *rewardpointssettingsRepository) Fetch(c context.Context) ([]domain.RewardPointsSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var rewardpointssettingss []domain.RewardPointsSettings

	err = cursor.All(c, &rewardpointssettingss)
	if rewardpointssettingss == nil {
		return []domain.RewardPointsSettings{}, err
	}

	return rewardpointssettingss, err
}

func (tr *rewardpointssettingsRepository) FetchByID(c context.Context, rewardpointssettingsID string) (domain.RewardPointsSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var rewardpointssettings domain.RewardPointsSettings

	idHex, err := primitive.ObjectIDFromHex(rewardpointssettingsID)
	if err != nil {
		return rewardpointssettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&rewardpointssettings)
	return rewardpointssettings, err
}
