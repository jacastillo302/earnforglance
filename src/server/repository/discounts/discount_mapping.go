package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type discountmappingRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountMappingRepository(db mongo.Database, collection string) domain.DiscountMappingRepository {
	return &discountmappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountmappingRepository) CreateMany(c context.Context, items []domain.DiscountMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *discountmappingRepository) Create(c context.Context, discountmapping *domain.DiscountMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discountmapping)

	return err
}

func (ur *discountmappingRepository) Update(c context.Context, discountmapping *domain.DiscountMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discountmapping.ID}
	update := bson.M{
		"$set": discountmapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *discountmappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountmappingRepository) Fetch(c context.Context) ([]domain.DiscountMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discountmappings []domain.DiscountMapping

	err = cursor.All(c, &discountmappings)
	if discountmappings == nil {
		return []domain.DiscountMapping{}, err
	}

	return discountmappings, err
}

func (tr *discountmappingRepository) FetchByID(c context.Context, discountmappingID string) (domain.DiscountMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var discountmapping domain.DiscountMapping

	idHex, err := primitive.ObjectIDFromHex(discountmappingID)
	if err != nil {
		return discountmapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discountmapping)
	return discountmapping, err
}
