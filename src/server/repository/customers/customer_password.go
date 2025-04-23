package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type customerpasswordsRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerPasswordRepository(db mongo.Database, collection string) domain.CustomerPasswordRepository {
	return &customerpasswordsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customerpasswordsRepository) CreateMany(c context.Context, items []domain.CustomerPassword) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *customerpasswordsRepository) Create(c context.Context, customerpasswords *domain.CustomerPassword) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customerpasswords)

	return err
}

func (ur *customerpasswordsRepository) Update(c context.Context, customerpasswords *domain.CustomerPassword) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customerpasswords.ID}
	update := bson.M{
		"$set": customerpasswords,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customerpasswordsRepository) Delete(c context.Context, customerpasswords string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customerpasswords}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *customerpasswordsRepository) Fetch(c context.Context) ([]domain.CustomerPassword, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customerpasswords []domain.CustomerPassword

	err = cursor.All(c, &customerpasswords)
	if customerpasswords == nil {
		return []domain.CustomerPassword{}, err
	}

	return customerpasswords, err
}

func (tr *customerpasswordsRepository) FetchByID(c context.Context, customerpasswordsID string) (domain.CustomerPassword, error) {
	collection := tr.database.Collection(tr.collection)

	var customerpasswords domain.CustomerPassword

	idHex, err := bson.ObjectIDFromHex(customerpasswordsID)
	if err != nil {
		return customerpasswords, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customerpasswords)
	return customerpasswords, err
}
