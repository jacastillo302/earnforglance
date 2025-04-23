package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type recurringpaymenthistoryRepository struct {
	database   mongo.Database
	collection string
}

func NewRecurringPaymentHistoryRepository(db mongo.Database, collection string) domain.RecurringPaymentHistoryRepository {
	return &recurringpaymenthistoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *recurringpaymenthistoryRepository) CreateMany(c context.Context, items []domain.RecurringPaymentHistory) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *recurringpaymenthistoryRepository) Create(c context.Context, recurringpaymenthistory *domain.RecurringPaymentHistory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, recurringpaymenthistory)

	return err
}

func (ur *recurringpaymenthistoryRepository) Update(c context.Context, recurringpaymenthistory *domain.RecurringPaymentHistory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": recurringpaymenthistory.ID}
	update := bson.M{
		"$set": recurringpaymenthistory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *recurringpaymenthistoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *recurringpaymenthistoryRepository) Fetch(c context.Context) ([]domain.RecurringPaymentHistory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var recurringpaymenthistorys []domain.RecurringPaymentHistory

	err = cursor.All(c, &recurringpaymenthistorys)
	if recurringpaymenthistorys == nil {
		return []domain.RecurringPaymentHistory{}, err
	}

	return recurringpaymenthistorys, err
}

func (tr *recurringpaymenthistoryRepository) FetchByID(c context.Context, recurringpaymenthistoryID string) (domain.RecurringPaymentHistory, error) {
	collection := tr.database.Collection(tr.collection)

	var recurringpaymenthistory domain.RecurringPaymentHistory

	idHex, err := bson.ObjectIDFromHex(recurringpaymenthistoryID)
	if err != nil {
		return recurringpaymenthistory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&recurringpaymenthistory)
	return recurringpaymenthistory, err
}
