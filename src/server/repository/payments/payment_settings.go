package repository

import (
	"context"

	domain "earnforglance/server/domain/payments"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type paymentsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewPaymentSettingsRepository(db mongo.Database, collection string) domain.PaymentSettingsRepository {
	return &paymentsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *paymentsettingsRepository) Create(c context.Context, paymentsettings *domain.PaymentSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, paymentsettings)

	return err
}

func (ur *paymentsettingsRepository) Update(c context.Context, paymentsettings *domain.PaymentSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": paymentsettings.ID}
	update := bson.M{
		"$set": paymentsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *paymentsettingsRepository) Delete(c context.Context, paymentsettings *domain.PaymentSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": paymentsettings.ID}
	_, err := collection.DeleteOne(c, filter)
	return err

}

func (ur *paymentsettingsRepository) Fetch(c context.Context) ([]domain.PaymentSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var paymentsettingss []domain.PaymentSettings

	err = cursor.All(c, &paymentsettingss)
	if paymentsettingss == nil {
		return []domain.PaymentSettings{}, err
	}

	return paymentsettingss, err
}

func (tr *paymentsettingsRepository) FetchByID(c context.Context, paymentsettingsID string) (domain.PaymentSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var paymentsettings domain.PaymentSettings

	idHex, err := primitive.ObjectIDFromHex(paymentsettingsID)
	if err != nil {
		return paymentsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&paymentsettings)
	return paymentsettings, err
}
