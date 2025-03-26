package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type giftcardRepository struct {
	database   mongo.Database
	collection string
}

func NewGiftCardRepository(db mongo.Database, collection string) domain.GiftCardRepository {
	return &giftcardRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *giftcardRepository) CreateMany(c context.Context, items []domain.GiftCard) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *giftcardRepository) Create(c context.Context, giftcard *domain.GiftCard) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, giftcard)

	return err
}

func (ur *giftcardRepository) Update(c context.Context, giftcard *domain.GiftCard) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": giftcard.ID}
	update := bson.M{
		"$set": giftcard,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *giftcardRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *giftcardRepository) Fetch(c context.Context) ([]domain.GiftCard, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var giftcards []domain.GiftCard

	err = cursor.All(c, &giftcards)
	if giftcards == nil {
		return []domain.GiftCard{}, err
	}

	return giftcards, err
}

func (tr *giftcardRepository) FetchByID(c context.Context, giftcardID string) (domain.GiftCard, error) {
	collection := tr.database.Collection(tr.collection)

	var giftcard domain.GiftCard

	idHex, err := primitive.ObjectIDFromHex(giftcardID)
	if err != nil {
		return giftcard, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&giftcard)
	return giftcard, err
}
