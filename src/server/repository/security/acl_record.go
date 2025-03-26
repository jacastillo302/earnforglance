package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type aclrecordRepository struct {
	database   mongo.Database
	collection string
}

func NewAclRecordRepository(db mongo.Database, collection string) domain.AclRecordRepository {
	return &aclrecordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *aclrecordRepository) CreateMany(c context.Context, items []domain.AclRecord) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *aclrecordRepository) Create(c context.Context, aclrecord *domain.AclRecord) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, aclrecord)

	return err
}

func (ur *aclrecordRepository) Update(c context.Context, aclrecord *domain.AclRecord) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": aclrecord.ID}
	update := bson.M{
		"$set": aclrecord,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *aclrecordRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *aclrecordRepository) Fetch(c context.Context) ([]domain.AclRecord, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var aclrecords []domain.AclRecord

	err = cursor.All(c, &aclrecords)
	if aclrecords == nil {
		return []domain.AclRecord{}, err
	}

	return aclrecords, err
}

func (tr *aclrecordRepository) FetchByID(c context.Context, aclrecordID string) (domain.AclRecord, error) {
	collection := tr.database.Collection(tr.collection)

	var aclrecord domain.AclRecord

	idHex, err := primitive.ObjectIDFromHex(aclrecordID)
	if err != nil {
		return aclrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&aclrecord)
	return aclrecord, err
}
