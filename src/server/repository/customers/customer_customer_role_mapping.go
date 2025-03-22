package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customercustomerrolemappingRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerCustomerRoleMappingRepository(db mongo.Database, collection string) domain.CustomerCustomerRoleMappingRepository {
	return &customercustomerrolemappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customercustomerrolemappingRepository) Create(c context.Context, customercustomerrolemapping *domain.CustomerCustomerRoleMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customercustomerrolemapping)

	return err
}

func (ur *customercustomerrolemappingRepository) Update(c context.Context, customercustomerrolemapping *domain.CustomerCustomerRoleMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customercustomerrolemapping.ID}
	update := bson.M{
		"$set": customercustomerrolemapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customercustomerrolemappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *customercustomerrolemappingRepository) Fetch(c context.Context) ([]domain.CustomerCustomerRoleMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customercustomerrolemappings []domain.CustomerCustomerRoleMapping

	err = cursor.All(c, &customercustomerrolemappings)
	if customercustomerrolemappings == nil {
		return []domain.CustomerCustomerRoleMapping{}, err
	}

	return customercustomerrolemappings, err
}

func (tr *customercustomerrolemappingRepository) FetchByID(c context.Context, customercustomerrolemappingID string) (domain.CustomerCustomerRoleMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var customercustomerrolemapping domain.CustomerCustomerRoleMapping

	idHex, err := primitive.ObjectIDFromHex(customercustomerrolemappingID)
	if err != nil {
		return customercustomerrolemapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customercustomerrolemapping)
	return customercustomerrolemapping, err
}
