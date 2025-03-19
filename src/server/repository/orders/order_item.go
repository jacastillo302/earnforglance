package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type orderitemRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderItemRepository(db mongo.Database, collection string) domain.OrderItemRepository {
	return &orderitemRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *orderitemRepository) Create(c context.Context, orderitem *domain.OrderItem) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, orderitem)

	return err
}

func (ur *orderitemRepository) Update(c context.Context, orderitem *domain.OrderItem) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": orderitem.ID}
	update := bson.M{
		"$set": orderitem,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *orderitemRepository) Delete(c context.Context, orderitem *domain.OrderItem) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": orderitem.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *orderitemRepository) Fetch(c context.Context) ([]domain.OrderItem, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var orderitems []domain.OrderItem

	err = cursor.All(c, &orderitems)
	if orderitems == nil {
		return []domain.OrderItem{}, err
	}

	return orderitems, err
}

func (tr *orderitemRepository) FetchByID(c context.Context, orderitemID string) (domain.OrderItem, error) {
	collection := tr.database.Collection(tr.collection)

	var orderitem domain.OrderItem

	idHex, err := primitive.ObjectIDFromHex(orderitemID)
	if err != nil {
		return orderitem, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&orderitem)
	return orderitem, err
}
