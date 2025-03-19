package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type rewardpointshistoryRepository struct {
	database   mongo.Database
	collection string
}

func NewRewardPointsHistoryRepository(db mongo.Database, collection string) domain.RewardPointsHistoryRepository {
	return &rewardpointshistoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *rewardpointshistoryRepository) Create(c context.Context, rewardpointshistory *domain.RewardPointsHistory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, rewardpointshistory)

	return err
}

func (ur *rewardpointshistoryRepository) Update(c context.Context, rewardpointshistory *domain.RewardPointsHistory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": rewardpointshistory.ID}
	update := bson.M{
		"$set": rewardpointshistory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *rewardpointshistoryRepository) Delete(c context.Context, rewardpointshistory *domain.RewardPointsHistory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": rewardpointshistory.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *rewardpointshistoryRepository) Fetch(c context.Context) ([]domain.RewardPointsHistory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var rewardpointshistorys []domain.RewardPointsHistory

	err = cursor.All(c, &rewardpointshistorys)
	if rewardpointshistorys == nil {
		return []domain.RewardPointsHistory{}, err
	}

	return rewardpointshistorys, err
}

func (tr *rewardpointshistoryRepository) FetchByID(c context.Context, rewardpointshistoryID string) (domain.RewardPointsHistory, error) {
	collection := tr.database.Collection(tr.collection)

	var rewardpointshistory domain.RewardPointsHistory

	idHex, err := primitive.ObjectIDFromHex(rewardpointshistoryID)
	if err != nil {
		return rewardpointshistory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&rewardpointshistory)
	return rewardpointshistory, err
}
