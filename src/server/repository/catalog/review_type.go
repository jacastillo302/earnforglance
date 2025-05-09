package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type reviewtypeRepository struct {
	database   mongo.Database
	collection string
}

func NewReviewTypeRepository(db mongo.Database, collection string) domain.ReviewTypeRepository {
	return &reviewtypeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *reviewtypeRepository) CreateMany(c context.Context, items []domain.ReviewType) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *reviewtypeRepository) Create(c context.Context, reviewtype *domain.ReviewType) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, reviewtype)

	return err
}

func (ur *reviewtypeRepository) Update(c context.Context, reviewtype *domain.ReviewType) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": reviewtype.ID}
	update := bson.M{
		"$set": reviewtype,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *reviewtypeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *reviewtypeRepository) Fetch(c context.Context) ([]domain.ReviewType, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var reviewtypes []domain.ReviewType

	err = cursor.All(c, &reviewtypes)
	if reviewtypes == nil {
		return []domain.ReviewType{}, err
	}

	return reviewtypes, err
}

func (tr *reviewtypeRepository) FetchByID(c context.Context, reviewtypeID string) (domain.ReviewType, error) {
	collection := tr.database.Collection(tr.collection)

	var reviewtype domain.ReviewType

	idHex, err := bson.ObjectIDFromHex(reviewtypeID)
	if err != nil {
		return reviewtype, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&reviewtype)
	return reviewtype, err
}
