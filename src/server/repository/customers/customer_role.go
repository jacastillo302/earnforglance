package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customerroleRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerRoleRepository(db mongo.Database, collection string) domain.CustomerRoleRepository {
	return &customerroleRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customerroleRepository) CreateMany(c context.Context, items []domain.CustomerRole) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *customerroleRepository) Create(c context.Context, customerrole *domain.CustomerRole) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customerrole)

	return err
}

func (ur *customerroleRepository) Update(c context.Context, customerrole *domain.CustomerRole) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customerrole.ID}
	update := bson.M{
		"$set": customerrole,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customerroleRepository) Delete(c context.Context, customerrole string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customerrole}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *customerroleRepository) Fetch(c context.Context) ([]domain.CustomerRole, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customerroles []domain.CustomerRole

	err = cursor.All(c, &customerroles)
	if customerroles == nil {
		return []domain.CustomerRole{}, err
	}

	return customerroles, err
}

func (tr *customerroleRepository) FetchByID(c context.Context, customerroleID string) (domain.CustomerRole, error) {
	collection := tr.database.Collection(tr.collection)

	var customerrole domain.CustomerRole

	idHex, err := primitive.ObjectIDFromHex(customerroleID)
	if err != nil {
		return customerrole, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customerrole)
	return customerrole, err
}
