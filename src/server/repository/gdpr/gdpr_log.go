package repository

import (
	"context"

	domain "earnforglance/server/domain/gdpr"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type gdprlogRepository struct {
	database   mongo.Database
	collection string
}

func NewGdprLogRepository(db mongo.Database, collection string) domain.GdprLogRepository {
	return &gdprlogRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *gdprlogRepository) CreateMany(c context.Context, items []domain.GdprLog) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *gdprlogRepository) Create(c context.Context, gdprlog *domain.GdprLog) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, gdprlog)

	return err
}

func (ur *gdprlogRepository) Update(c context.Context, gdprlog *domain.GdprLog) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": gdprlog.ID}
	update := bson.M{
		"$set": gdprlog,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *gdprlogRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *gdprlogRepository) Fetch(c context.Context) ([]domain.GdprLog, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var gdprlogs []domain.GdprLog

	err = cursor.All(c, &gdprlogs)
	if gdprlogs == nil {
		return []domain.GdprLog{}, err
	}

	return gdprlogs, err
}

func (tr *gdprlogRepository) FetchByID(c context.Context, gdprlogID string) (domain.GdprLog, error) {
	collection := tr.database.Collection(tr.collection)

	var gdprlog domain.GdprLog

	idHex, err := bson.ObjectIDFromHex(gdprlogID)
	if err != nil {
		return gdprlog, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&gdprlog)
	return gdprlog, err
}
