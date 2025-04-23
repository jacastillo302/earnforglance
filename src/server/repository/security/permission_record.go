package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type permissionrecordRepository struct {
	database   mongo.Database
	collection string
}

func NewPermissionRecordRepository(db mongo.Database, collection string) domain.PermissionRecordRepository {
	return &permissionrecordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *permissionrecordRepository) CreateMany(c context.Context, items []domain.PermissionRecord) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *permissionrecordRepository) Create(c context.Context, permissionrecord *domain.PermissionRecord) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, permissionrecord)

	return err
}

func (ur *permissionrecordRepository) Update(c context.Context, permissionrecord *domain.PermissionRecord) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": permissionrecord.ID}
	update := bson.M{
		"$set": permissionrecord,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *permissionrecordRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *permissionrecordRepository) Fetch(c context.Context) ([]domain.PermissionRecord, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var permissionrecords []domain.PermissionRecord

	err = cursor.All(c, &permissionrecords)
	if permissionrecords == nil {
		return []domain.PermissionRecord{}, err
	}

	return permissionrecords, err
}

func (tr *permissionrecordRepository) FetchByID(c context.Context, permissionrecordID string) (domain.PermissionRecord, error) {
	collection := tr.database.Collection(tr.collection)

	var permissionrecord domain.PermissionRecord

	idHex, err := bson.ObjectIDFromHex(permissionrecordID)
	if err != nil {
		return permissionrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&permissionrecord)
	return permissionrecord, err
}
