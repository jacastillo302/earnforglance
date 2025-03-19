package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type giftcardusagehistoryRepository struct {
	database   mongo.Database
	collection string
}

func NewGiftCardUsageHistoryRepository(db mongo.Database, collection string) domain.GiftCardUsageHistoryRepository {
	return &giftcardusagehistoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *giftcardusagehistoryRepository) Create(c context.Context, giftcardusagehistory *domain.GiftCardUsageHistory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, giftcardusagehistory)

	return err
}

func (ur *giftcardusagehistoryRepository) Update(c context.Context, giftcardusagehistory *domain.GiftCardUsageHistory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": giftcardusagehistory.ID}
	update := bson.M{
		"$set": giftcardusagehistory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *giftcardusagehistoryRepository) Delete(c context.Context, giftcardusagehistory *domain.GiftCardUsageHistory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": giftcardusagehistory.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *giftcardusagehistoryRepository) Fetch(c context.Context) ([]domain.GiftCardUsageHistory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var giftcardusagehistorys []domain.GiftCardUsageHistory

	err = cursor.All(c, &giftcardusagehistorys)
	if giftcardusagehistorys == nil {
		return []domain.GiftCardUsageHistory{}, err
	}

	return giftcardusagehistorys, err
}

func (tr *giftcardusagehistoryRepository) FetchByID(c context.Context, giftcardusagehistoryID string) (domain.GiftCardUsageHistory, error) {
	collection := tr.database.Collection(tr.collection)

	var giftcardusagehistory domain.GiftCardUsageHistory

	idHex, err := primitive.ObjectIDFromHex(giftcardusagehistoryID)
	if err != nil {
		return giftcardusagehistory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&giftcardusagehistory)
	return giftcardusagehistory, err
}
