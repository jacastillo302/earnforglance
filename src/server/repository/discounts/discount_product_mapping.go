package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type discountproductmappingRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountProductMappingRepository(db mongo.Database, collection string) domain.DiscountProductMappingRepository {
	return &discountproductmappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountproductmappingRepository) CreateMany(c context.Context, items []domain.DiscountProductMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *discountproductmappingRepository) Create(c context.Context, discountproductmapping *domain.DiscountProductMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discountproductmapping)

	return err
}

func (ur *discountproductmappingRepository) Update(c context.Context, discountproductmapping *domain.DiscountProductMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discountproductmapping.ID}
	update := bson.M{
		"$set": discountproductmapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *discountproductmappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountproductmappingRepository) Fetch(c context.Context) ([]domain.DiscountProductMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discountproductmappings []domain.DiscountProductMapping

	err = cursor.All(c, &discountproductmappings)
	if discountproductmappings == nil {
		return []domain.DiscountProductMapping{}, err
	}

	return discountproductmappings, err
}

func (tr *discountproductmappingRepository) FetchByID(c context.Context, discountproductmappingID string) (domain.DiscountProductMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var discountproductmapping domain.DiscountProductMapping

	idHex, err := primitive.ObjectIDFromHex(discountproductmappingID)
	if err != nil {
		return discountproductmapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discountproductmapping)
	return discountproductmapping, err
}
