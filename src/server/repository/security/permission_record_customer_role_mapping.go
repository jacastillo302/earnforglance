package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type permissionrecordcustomerrolemappingRepository struct {
	database   mongo.Database
	collection string
}

func NewPermissionRecordCustomerRoleMappingRepository(db mongo.Database, collection string) domain.PermissionRecordCustomerRoleMappingRepository {
	return &permissionrecordcustomerrolemappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *permissionrecordcustomerrolemappingRepository) CreateMany(c context.Context, items []domain.PermissionRecordCustomerRoleMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *permissionrecordcustomerrolemappingRepository) Create(c context.Context, permissionrecordcustomerrolemapping *domain.PermissionRecordCustomerRoleMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, permissionrecordcustomerrolemapping)

	return err
}

func (ur *permissionrecordcustomerrolemappingRepository) Update(c context.Context, permissionrecordcustomerrolemapping *domain.PermissionRecordCustomerRoleMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": permissionrecordcustomerrolemapping.ID}
	update := bson.M{
		"$set": permissionrecordcustomerrolemapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *permissionrecordcustomerrolemappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *permissionrecordcustomerrolemappingRepository) Fetch(c context.Context) ([]domain.PermissionRecordCustomerRoleMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var permissionrecordcustomerrolemappings []domain.PermissionRecordCustomerRoleMapping

	err = cursor.All(c, &permissionrecordcustomerrolemappings)
	if permissionrecordcustomerrolemappings == nil {
		return []domain.PermissionRecordCustomerRoleMapping{}, err
	}

	return permissionrecordcustomerrolemappings, err
}

func (tr *permissionrecordcustomerrolemappingRepository) FetchByID(c context.Context, permissionrecordcustomerrolemappingID string) (domain.PermissionRecordCustomerRoleMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var permissionrecordcustomerrolemapping domain.PermissionRecordCustomerRoleMapping

	idHex, err := primitive.ObjectIDFromHex(permissionrecordcustomerrolemappingID)
	if err != nil {
		return permissionrecordcustomerrolemapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&permissionrecordcustomerrolemapping)
	return permissionrecordcustomerrolemapping, err
}
