package repository

import (
	"context"

	domain "earnforglance/server/domain/seo"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type urlrecordRepository struct {
	database   mongo.Database
	collection string
}

func NewUrlRecordRepository(db mongo.Database, collection string) domain.UrlRecordRepository {
	return &urlrecordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *urlrecordRepository) CreateMany(c context.Context, items []domain.UrlRecord) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *urlrecordRepository) Create(c context.Context, urlrecord *domain.UrlRecord) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, urlrecord)

	return err
}

func (ur *urlrecordRepository) Update(c context.Context, urlrecord *domain.UrlRecord) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": urlrecord.ID}
	update := bson.M{
		"$set": urlrecord,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *urlrecordRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *urlrecordRepository) Fetch(c context.Context) ([]domain.UrlRecord, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var urlrecords []domain.UrlRecord

	err = cursor.All(c, &urlrecords)
	if urlrecords == nil {
		return []domain.UrlRecord{}, err
	}

	return urlrecords, err
}

func (tr *urlrecordRepository) FetchByID(c context.Context, urlrecordID string) (domain.UrlRecord, error) {
	collection := tr.database.Collection(tr.collection)

	var urlrecord domain.UrlRecord

	idHex, err := bson.ObjectIDFromHex(urlrecordID)
	if err != nil {
		return urlrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&urlrecord)
	return urlrecord, err
}
