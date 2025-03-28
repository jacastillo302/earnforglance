package repository

import (
	"context"

	domain "earnforglance/server/domain/api"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type api_clientRepository struct {
	database   mongo.Database
	collection string
}

func NewApiClientRepository(db mongo.Database, collection string) domain.ApiClientRepository {
	return &api_clientRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *api_clientRepository) Create(c context.Context, api_client *domain.ApiClient) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, api_client)

	return err
}

func (ur *api_clientRepository) CreateMany(c context.Context, items []domain.ApiClient) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *api_clientRepository) Update(c context.Context, api_client *domain.ApiClient) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": api_client.ID}
	update := bson.M{
		"$set": api_client,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *api_clientRepository) Delete(c context.Context, api_client string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": api_client}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *api_clientRepository) Fetch(c context.Context) ([]domain.ApiClient, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var api_clients []domain.ApiClient

	err = cursor.All(c, &api_clients)
	if api_clients == nil {
		return []domain.ApiClient{}, err
	}

	return api_clients, err
}

func (tr *api_clientRepository) FetchByID(c context.Context, api_clientID string) (domain.ApiClient, error) {
	collection := tr.database.Collection(tr.collection)

	var api_client domain.ApiClient

	idHex, err := primitive.ObjectIDFromHex(api_clientID)
	if err != nil {
		return api_client, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&api_client)
	return api_client, err
}
