package repository

import (
	"context"

	domain "earnforglance/server/domain/vendors"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type vendornoteRepository struct {
	database   mongo.Database
	collection string
}

func NewVendorNoteRepository(db mongo.Database, collection string) domain.VendorNoteRepository {
	return &vendornoteRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *vendornoteRepository) CreateMany(c context.Context, items []domain.VendorNote) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *vendornoteRepository) Create(c context.Context, vendornote *domain.VendorNote) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, vendornote)

	return err
}

func (ur *vendornoteRepository) Update(c context.Context, vendornote *domain.VendorNote) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": vendornote.ID}
	update := bson.M{
		"$set": vendornote,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *vendornoteRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *vendornoteRepository) Fetch(c context.Context) ([]domain.VendorNote, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var vendornotes []domain.VendorNote

	err = cursor.All(c, &vendornotes)
	if vendornotes == nil {
		return []domain.VendorNote{}, err
	}

	return vendornotes, err
}

func (tr *vendornoteRepository) FetchByID(c context.Context, vendornoteID string) (domain.VendorNote, error) {
	collection := tr.database.Collection(tr.collection)

	var vendornote domain.VendorNote

	idHex, err := primitive.ObjectIDFromHex(vendornoteID)
	if err != nil {
		return vendornote, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&vendornote)
	return vendornote, err
}
