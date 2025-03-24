package repository

import (
	"context"

	domain "earnforglance/server/domain/vendors"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type vendorattributevalueRepository struct {
	database   mongo.Database
	collection string
}

func NewVendorAttributeValueRepository(db mongo.Database, collection string) domain.VendorAttributeValueRepository {
	return &vendorattributevalueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *vendorattributevalueRepository) CreateMany(c context.Context, items []domain.VendorAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *vendorattributevalueRepository) Create(c context.Context, vendorattributevalue *domain.VendorAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, vendorattributevalue)

	return err
}

func (ur *vendorattributevalueRepository) Update(c context.Context, vendorattributevalue *domain.VendorAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": vendorattributevalue.ID}
	update := bson.M{
		"$set": vendorattributevalue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *vendorattributevalueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *vendorattributevalueRepository) Fetch(c context.Context) ([]domain.VendorAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var vendorattributevalues []domain.VendorAttributeValue

	err = cursor.All(c, &vendorattributevalues)
	if vendorattributevalues == nil {
		return []domain.VendorAttributeValue{}, err
	}

	return vendorattributevalues, err
}

func (tr *vendorattributevalueRepository) FetchByID(c context.Context, vendorattributevalueID string) (domain.VendorAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var vendorattributevalue domain.VendorAttributeValue

	idHex, err := primitive.ObjectIDFromHex(vendorattributevalueID)
	if err != nil {
		return vendorattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&vendorattributevalue)
	return vendorattributevalue, err
}
