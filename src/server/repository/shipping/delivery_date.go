package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type deliverydateRepository struct {
	database   mongo.Database
	collection string
}

func NewDeliveryDateRepository(db mongo.Database, collection string) domain.DeliveryDateRepository {
	return &deliverydateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *deliverydateRepository) CreateMany(c context.Context, items []domain.DeliveryDate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *deliverydateRepository) Create(c context.Context, deliverydate *domain.DeliveryDate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, deliverydate)

	return err
}

func (ur *deliverydateRepository) Update(c context.Context, deliverydate *domain.DeliveryDate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": deliverydate.ID}
	update := bson.M{
		"$set": deliverydate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *deliverydateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *deliverydateRepository) Fetch(c context.Context) ([]domain.DeliveryDate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var deliverydates []domain.DeliveryDate

	err = cursor.All(c, &deliverydates)
	if deliverydates == nil {
		return []domain.DeliveryDate{}, err
	}

	return deliverydates, err
}

func (tr *deliverydateRepository) FetchByID(c context.Context, deliverydateID string) (domain.DeliveryDate, error) {
	collection := tr.database.Collection(tr.collection)

	var deliverydate domain.DeliveryDate

	idHex, err := primitive.ObjectIDFromHex(deliverydateID)
	if err != nil {
		return deliverydate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&deliverydate)
	return deliverydate, err
}
