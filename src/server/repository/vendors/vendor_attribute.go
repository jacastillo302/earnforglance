package repository

import (
	"context"

	domain "earnforglance/server/domain/vendors"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type vendorattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewVendorAttributeRepository(db mongo.Database, collection string) domain.VendorAttributeRepository {
	return &vendorattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *vendorattributeRepository) CreateMany(c context.Context, items []domain.VendorAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *vendorattributeRepository) Create(c context.Context, vendorattribute *domain.VendorAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, vendorattribute)

	return err
}

func (ur *vendorattributeRepository) Update(c context.Context, vendorattribute *domain.VendorAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": vendorattribute.ID}
	update := bson.M{
		"$set": vendorattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *vendorattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *vendorattributeRepository) Fetch(c context.Context) ([]domain.VendorAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var vendorattributes []domain.VendorAttribute

	err = cursor.All(c, &vendorattributes)
	if vendorattributes == nil {
		return []domain.VendorAttribute{}, err
	}

	return vendorattributes, err
}

func (tr *vendorattributeRepository) FetchByID(c context.Context, vendorattributeID string) (domain.VendorAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var vendorattribute domain.VendorAttribute

	idHex, err := primitive.ObjectIDFromHex(vendorattributeID)
	if err != nil {
		return vendorattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&vendorattribute)
	return vendorattribute, err
}
