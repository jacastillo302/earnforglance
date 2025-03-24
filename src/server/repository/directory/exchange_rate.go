package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type exchangeRateRepository struct {
	database   mongo.Database
	collection string
}

func NewExchangeRateRepository(db mongo.Database, collection string) domain.ExchangeRateRepository {
	return &exchangeRateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *exchangeRateRepository) CreateMany(c context.Context, items []domain.ExchangeRate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *exchangeRateRepository) Create(c context.Context, exchangeRate *domain.ExchangeRate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, exchangeRate)

	return err
}

func (ur *exchangeRateRepository) Update(c context.Context, exchangeRate *domain.ExchangeRate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": exchangeRate.ID}
	update := bson.M{
		"$set": exchangeRate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *exchangeRateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *exchangeRateRepository) Fetch(c context.Context) ([]domain.ExchangeRate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var exchangeRates []domain.ExchangeRate

	err = cursor.All(c, &exchangeRates)
	if exchangeRates == nil {
		return []domain.ExchangeRate{}, err
	}

	return exchangeRates, err
}

func (tr *exchangeRateRepository) FetchByID(c context.Context, exchangeRateID string) (domain.ExchangeRate, error) {
	collection := tr.database.Collection(tr.collection)

	var exchangeRate domain.ExchangeRate

	idHex, err := primitive.ObjectIDFromHex(exchangeRateID)
	if err != nil {
		return exchangeRate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&exchangeRate)
	return exchangeRate, err
}
