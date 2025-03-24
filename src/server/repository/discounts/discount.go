package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type discountRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountRepository(db mongo.Database, collection string) domain.DiscountRepository {
	return &discountRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountRepository) CreateMany(c context.Context, items []domain.Discount) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *discountRepository) Create(c context.Context, discount *domain.Discount) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discount)

	return err
}

func (ur *discountRepository) Update(c context.Context, discount *domain.Discount) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discount.ID}
	update := bson.M{
		"$set": discount,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *discountRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountRepository) Fetch(c context.Context) ([]domain.Discount, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discounts []domain.Discount

	err = cursor.All(c, &discounts)
	if discounts == nil {
		return []domain.Discount{}, err
	}

	return discounts, err
}

func (tr *discountRepository) FetchByID(c context.Context, discountID string) (domain.Discount, error) {
	collection := tr.database.Collection(tr.collection)

	var discount domain.Discount

	idHex, err := primitive.ObjectIDFromHex(discountID)
	if err != nil {
		return discount, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discount)
	return discount, err
}
