package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type recurringpaymentRepository struct {
	database   mongo.Database
	collection string
}

func NewRecurringPaymentRepository(db mongo.Database, collection string) domain.RecurringPaymentRepository {
	return &recurringpaymentRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *recurringpaymentRepository) Create(c context.Context, recurringpayment *domain.RecurringPayment) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, recurringpayment)

	return err
}

func (ur *recurringpaymentRepository) Update(c context.Context, recurringpayment *domain.RecurringPayment) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": recurringpayment.ID}
	update := bson.M{
		"$set": recurringpayment,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *recurringpaymentRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *recurringpaymentRepository) Fetch(c context.Context) ([]domain.RecurringPayment, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var recurringpayments []domain.RecurringPayment

	err = cursor.All(c, &recurringpayments)
	if recurringpayments == nil {
		return []domain.RecurringPayment{}, err
	}

	return recurringpayments, err
}

func (tr *recurringpaymentRepository) FetchByID(c context.Context, recurringpaymentID string) (domain.RecurringPayment, error) {
	collection := tr.database.Collection(tr.collection)

	var recurringpayment domain.RecurringPayment

	idHex, err := primitive.ObjectIDFromHex(recurringpaymentID)
	if err != nil {
		return recurringpayment, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&recurringpayment)
	return recurringpayment, err
}
