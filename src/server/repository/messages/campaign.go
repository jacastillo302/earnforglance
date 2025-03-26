package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type campaignRepository struct {
	database   mongo.Database
	collection string
}

func NewCampaignRepository(db mongo.Database, collection string) domain.CampaignRepository {
	return &campaignRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *campaignRepository) CreateMany(c context.Context, items []domain.Campaign) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *campaignRepository) Create(c context.Context, campaign *domain.Campaign) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, campaign)

	return err
}

func (ur *campaignRepository) Update(c context.Context, campaign *domain.Campaign) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": campaign.ID}
	update := bson.M{
		"$set": campaign,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *campaignRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *campaignRepository) Fetch(c context.Context) ([]domain.Campaign, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var campaigns []domain.Campaign

	err = cursor.All(c, &campaigns)
	if campaigns == nil {
		return []domain.Campaign{}, err
	}

	return campaigns, err
}

func (tr *campaignRepository) FetchByID(c context.Context, campaignID string) (domain.Campaign, error) {
	collection := tr.database.Collection(tr.collection)

	var campaign domain.Campaign

	idHex, err := primitive.ObjectIDFromHex(campaignID)
	if err != nil {
		return campaign, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&campaign)
	return campaign, err
}
