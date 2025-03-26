package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ordernoteRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderNoteRepository(db mongo.Database, collection string) domain.OrderNoteRepository {
	return &ordernoteRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *ordernoteRepository) CreateMany(c context.Context, items []domain.OrderNote) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *ordernoteRepository) Create(c context.Context, ordernote *domain.OrderNote) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, ordernote)

	return err
}

func (ur *ordernoteRepository) Update(c context.Context, ordernote *domain.OrderNote) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": ordernote.ID}
	update := bson.M{
		"$set": ordernote,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *ordernoteRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *ordernoteRepository) Fetch(c context.Context) ([]domain.OrderNote, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var ordernotes []domain.OrderNote

	err = cursor.All(c, &ordernotes)
	if ordernotes == nil {
		return []domain.OrderNote{}, err
	}

	return ordernotes, err
}

func (tr *ordernoteRepository) FetchByID(c context.Context, ordernoteID string) (domain.OrderNote, error) {
	collection := tr.database.Collection(tr.collection)

	var ordernote domain.OrderNote

	idHex, err := primitive.ObjectIDFromHex(ordernoteID)
	if err != nil {
		return ordernote, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&ordernote)
	return ordernote, err
}
