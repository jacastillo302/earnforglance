package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type stockquantitychangeRepository struct {
	database   mongo.Database
	collection string
}

func NewStockQuantityChangeRepository(db mongo.Database, collection string) domain.StockQuantityChangeRepository {
	return &stockquantitychangeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *stockquantitychangeRepository) CreateMany(c context.Context, items []domain.StockQuantityChange) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *stockquantitychangeRepository) Create(c context.Context, stockquantitychange *domain.StockQuantityChange) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, stockquantitychange)

	return err
}

func (ur *stockquantitychangeRepository) Update(c context.Context, stockquantitychange *domain.StockQuantityChange) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": stockquantitychange.ID}
	update := bson.M{
		"$set": stockquantitychange,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *stockquantitychangeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *stockquantitychangeRepository) Fetch(c context.Context) ([]domain.StockQuantityChange, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var stockquantitychanges []domain.StockQuantityChange

	err = cursor.All(c, &stockquantitychanges)
	if stockquantitychanges == nil {
		return []domain.StockQuantityChange{}, err
	}

	return stockquantitychanges, err
}

func (tr *stockquantitychangeRepository) FetchByID(c context.Context, stockquantitychangeID string) (domain.StockQuantityChange, error) {
	collection := tr.database.Collection(tr.collection)

	var stockquantitychange domain.StockQuantityChange

	idHex, err := primitive.ObjectIDFromHex(stockquantitychangeID)
	if err != nil {
		return stockquantitychange, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&stockquantitychange)
	return stockquantitychange, err
}
