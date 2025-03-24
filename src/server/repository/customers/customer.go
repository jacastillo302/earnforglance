package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customerRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerRepository(db mongo.Database, collection string) domain.CustomerRepository {
	return &customerRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customerRepository) CreateMany(c context.Context, items []domain.Customer) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *customerRepository) Create(c context.Context, customer *domain.Customer) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customer)

	return err
}

func (ur *customerRepository) Update(c context.Context, customer *domain.Customer) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customer.ID}
	update := bson.M{
		"$set": customer,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customerRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *customerRepository) Fetch(c context.Context) ([]domain.Customer, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customers []domain.Customer

	err = cursor.All(c, &customers)
	if customers == nil {
		return []domain.Customer{}, err
	}

	return customers, err
}

func (tr *customerRepository) FetchByID(c context.Context, customerID string) (domain.Customer, error) {
	collection := tr.database.Collection(tr.collection)

	var customer domain.Customer

	idHex, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return customer, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customer)
	return customer, err
}
