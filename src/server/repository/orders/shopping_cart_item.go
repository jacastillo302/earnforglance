package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type shoppingcartitemRepository struct {
	database   mongo.Database
	collection string
}

func NewShoppingCartItemRepository(db mongo.Database, collection string) domain.ShoppingCartItemRepository {
	return &shoppingcartitemRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shoppingcartitemRepository) CreateMany(c context.Context, items []domain.ShoppingCartItem) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *shoppingcartitemRepository) Create(c context.Context, shoppingcartitem *domain.ShoppingCartItem) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shoppingcartitem)

	return err
}

func (ur *shoppingcartitemRepository) Update(c context.Context, shoppingcartitem *domain.ShoppingCartItem) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shoppingcartitem.ID}
	update := bson.M{
		"$set": shoppingcartitem,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shoppingcartitemRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shoppingcartitemRepository) Fetch(c context.Context) ([]domain.ShoppingCartItem, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shoppingcartitems []domain.ShoppingCartItem

	err = cursor.All(c, &shoppingcartitems)
	if shoppingcartitems == nil {
		return []domain.ShoppingCartItem{}, err
	}

	return shoppingcartitems, err
}

func (tr *shoppingcartitemRepository) FetchByID(c context.Context, shoppingcartitemID string) (domain.ShoppingCartItem, error) {
	collection := tr.database.Collection(tr.collection)

	var shoppingcartitem domain.ShoppingCartItem

	idHex, err := primitive.ObjectIDFromHex(shoppingcartitemID)
	if err != nil {
		return shoppingcartitem, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shoppingcartitem)
	return shoppingcartitem, err
}
