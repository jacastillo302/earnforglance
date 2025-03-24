package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type orderRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderRepository(db mongo.Database, collection string) domain.OrderRepository {
	return &orderRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *orderRepository) CreateMany(c context.Context, items []domain.Order) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *orderRepository) Create(c context.Context, order *domain.Order) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, order)

	return err
}

func (ur *orderRepository) Update(c context.Context, order *domain.Order) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": order.ID}
	update := bson.M{
		"$set": order,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *orderRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *orderRepository) Fetch(c context.Context) ([]domain.Order, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var orders []domain.Order

	err = cursor.All(c, &orders)
	if orders == nil {
		return []domain.Order{}, err
	}

	return orders, err
}

func (tr *orderRepository) FetchByID(c context.Context, orderID string) (domain.Order, error) {
	collection := tr.database.Collection(tr.collection)

	var order domain.Order

	idHex, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return order, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&order)
	return order, err
}
