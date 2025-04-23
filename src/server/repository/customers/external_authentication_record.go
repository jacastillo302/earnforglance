package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type externalauthenticationrecordRepository struct {
	database   mongo.Database
	collection string
}

func NewExternalAuthenticationRecordRepository(db mongo.Database, collection string) domain.ExternalAuthenticationRecordRepository {
	return &externalauthenticationrecordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *externalauthenticationrecordRepository) CreateMany(c context.Context, items []domain.ExternalAuthenticationRecord) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *externalauthenticationrecordRepository) Create(c context.Context, externalauthenticationrecord *domain.ExternalAuthenticationRecord) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, externalauthenticationrecord)

	return err
}

func (ur *externalauthenticationrecordRepository) Update(c context.Context, externalauthenticationrecord *domain.ExternalAuthenticationRecord) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": externalauthenticationrecord.ID}
	update := bson.M{
		"$set": externalauthenticationrecord,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *externalauthenticationrecordRepository) Delete(c context.Context, externalauthenticationrecord string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": externalauthenticationrecord}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *externalauthenticationrecordRepository) Fetch(c context.Context) ([]domain.ExternalAuthenticationRecord, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var externalauthenticationrecords []domain.ExternalAuthenticationRecord

	err = cursor.All(c, &externalauthenticationrecords)
	if externalauthenticationrecords == nil {
		return []domain.ExternalAuthenticationRecord{}, err
	}

	return externalauthenticationrecords, err
}

func (tr *externalauthenticationrecordRepository) FetchByID(c context.Context, externalauthenticationrecordID string) (domain.ExternalAuthenticationRecord, error) {
	collection := tr.database.Collection(tr.collection)

	var externalauthenticationrecord domain.ExternalAuthenticationRecord

	idHex, err := bson.ObjectIDFromHex(externalauthenticationrecordID)
	if err != nil {
		return externalauthenticationrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&externalauthenticationrecord)
	return externalauthenticationrecord, err
}
