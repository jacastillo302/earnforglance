package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type discountusagehistoryRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountUsageHistoryRepository(db mongo.Database, collection string) domain.DiscountUsageHistoryRepository {
	return &discountusagehistoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountusagehistoryRepository) Create(c context.Context, discountusagehistory *domain.DiscountUsageHistory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discountusagehistory)

	return err
}

func (ur *discountusagehistoryRepository) Update(c context.Context, discountusagehistory *domain.DiscountUsageHistory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discountusagehistory.ID}
	update := bson.M{
		"$set": discountusagehistory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *discountusagehistoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountusagehistoryRepository) Fetch(c context.Context) ([]domain.DiscountUsageHistory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discountusagehistorys []domain.DiscountUsageHistory

	err = cursor.All(c, &discountusagehistorys)
	if discountusagehistorys == nil {
		return []domain.DiscountUsageHistory{}, err
	}

	return discountusagehistorys, err
}

func (tr *discountusagehistoryRepository) FetchByID(c context.Context, discountusagehistoryID string) (domain.DiscountUsageHistory, error) {
	collection := tr.database.Collection(tr.collection)

	var discountusagehistory domain.DiscountUsageHistory

	idHex, err := primitive.ObjectIDFromHex(discountusagehistoryID)
	if err != nil {
		return discountusagehistory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discountusagehistory)
	return discountusagehistory, err
}
