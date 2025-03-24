package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type currencyRepository struct {
	database   mongo.Database
	collection string
}

func NewCurrencyRepository(db mongo.Database, collection string) domain.CurrencyRepository {
	return &currencyRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *currencyRepository) CreateMany(c context.Context, items []domain.Currency) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *currencyRepository) Create(c context.Context, currency *domain.Currency) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, currency)

	return err
}

func (ur *currencyRepository) Update(c context.Context, currency *domain.Currency) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": currency.ID}
	update := bson.M{
		"$set": currency,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *currencyRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *currencyRepository) Fetch(c context.Context) ([]domain.Currency, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var currencies []domain.Currency

	err = cursor.All(c, &currencies)
	if currencies == nil {
		return []domain.Currency{}, err
	}

	return currencies, err
}

func (tr *currencyRepository) FetchByID(c context.Context, currencyID string) (domain.Currency, error) {
	collection := tr.database.Collection(tr.collection)

	var currency domain.Currency

	idHex, err := primitive.ObjectIDFromHex(currencyID)
	if err != nil {
		return currency, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&currency)
	return currency, err
}
