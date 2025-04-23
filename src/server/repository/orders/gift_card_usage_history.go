package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (ur *giftcardusagehistoryRepository) CreateMany(c context.Context, items []domain.GiftCardUsageHistory) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
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

func (ur *giftcardusagehistoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *giftcardusagehistoryRepository) Fetch(c context.Context) ([]domain.GiftCardUsageHistory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := bson.ObjectIDFromHex(giftcardusagehistoryID)
	if err != nil {
		return giftcardusagehistory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&giftcardusagehistory)
	return giftcardusagehistory, err
}
