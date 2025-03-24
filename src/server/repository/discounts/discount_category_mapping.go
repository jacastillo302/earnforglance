package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type discountcategorymappingRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountCategoryMappingRepository(db mongo.Database, collection string) domain.DiscountCategoryMappingRepository {
	return &discountcategorymappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountcategorymappingRepository) CreateMany(c context.Context, items []domain.DiscountCategoryMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *discountcategorymappingRepository) Create(c context.Context, discountcategorymapping *domain.DiscountCategoryMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discountcategorymapping)

	return err
}

func (ur *discountcategorymappingRepository) Update(c context.Context, discountcategorymapping *domain.DiscountCategoryMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discountcategorymapping.ID}
	update := bson.M{
		"$set": discountcategorymapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *discountcategorymappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountcategorymappingRepository) Fetch(c context.Context) ([]domain.DiscountCategoryMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discountcategorymappings []domain.DiscountCategoryMapping

	err = cursor.All(c, &discountcategorymappings)
	if discountcategorymappings == nil {
		return []domain.DiscountCategoryMapping{}, err
	}

	return discountcategorymappings, err
}

func (tr *discountcategorymappingRepository) FetchByID(c context.Context, discountcategorymappingID string) (domain.DiscountCategoryMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var discountcategorymapping domain.DiscountCategoryMapping

	idHex, err := primitive.ObjectIDFromHex(discountcategorymappingID)
	if err != nil {
		return discountcategorymapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discountcategorymapping)
	return discountcategorymapping, err
}
