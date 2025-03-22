package repository

import (
	"context"

	domain "earnforglance/server/domain/vendors"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type vendorRepository struct {
	database   mongo.Database
	collection string
}

func NewVendorRepository(db mongo.Database, collection string) domain.VendorRepository {
	return &vendorRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *vendorRepository) Create(c context.Context, vendor *domain.Vendor) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, vendor)

	return err
}

func (ur *vendorRepository) Update(c context.Context, vendor *domain.Vendor) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": vendor.ID}
	update := bson.M{
		"$set": vendor,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *vendorRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *vendorRepository) Fetch(c context.Context) ([]domain.Vendor, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var vendors []domain.Vendor

	err = cursor.All(c, &vendors)
	if vendors == nil {
		return []domain.Vendor{}, err
	}

	return vendors, err
}

func (tr *vendorRepository) FetchByID(c context.Context, vendorID string) (domain.Vendor, error) {
	collection := tr.database.Collection(tr.collection)

	var vendor domain.Vendor

	idHex, err := primitive.ObjectIDFromHex(vendorID)
	if err != nil {
		return vendor, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&vendor)
	return vendor, err
}
