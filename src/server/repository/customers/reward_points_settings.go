package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo"

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

func (ur *rewardpointssettingsRepository) Delete(c context.Context, rewardpointssettings *domain.RewardPointsSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": rewardpointssettings.ID}
	_, err := collection.DeleteOne(c, filter)
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
