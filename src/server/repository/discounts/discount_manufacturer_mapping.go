package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type discountmanufacturermappingRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountManufacturerMappingRepository(db mongo.Database, collection string) domain.DiscountManufacturerMappingRepository {
	return &discountmanufacturermappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountmanufacturermappingRepository) CreateMany(c context.Context, items []domain.DiscountManufacturerMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *discountmanufacturermappingRepository) Create(c context.Context, discountmanufacturermapping *domain.DiscountManufacturerMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discountmanufacturermapping)

	return err
}

func (ur *discountmanufacturermappingRepository) Update(c context.Context, discountmanufacturermapping *domain.DiscountManufacturerMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discountmanufacturermapping.ID}
	update := bson.M{
		"$set": discountmanufacturermapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *discountmanufacturermappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountmanufacturermappingRepository) Fetch(c context.Context) ([]domain.DiscountManufacturerMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discountmanufacturermappings []domain.DiscountManufacturerMapping

	err = cursor.All(c, &discountmanufacturermappings)
	if discountmanufacturermappings == nil {
		return []domain.DiscountManufacturerMapping{}, err
	}

	return discountmanufacturermappings, err
}

func (tr *discountmanufacturermappingRepository) FetchByID(c context.Context, discountmanufacturermappingID string) (domain.DiscountManufacturerMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var discountmanufacturermapping domain.DiscountManufacturerMapping

	idHex, err := bson.ObjectIDFromHex(discountmanufacturermappingID)
	if err != nil {
		return discountmanufacturermapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discountmanufacturermapping)
	return discountmanufacturermapping, err
}
