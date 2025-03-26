package repository

import (
	"context"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type discountrequirementRepository struct {
	database   mongo.Database
	collection string
}

func NewDiscountRequirementRepository(db mongo.Database, collection string) domain.DiscountRequirementRepository {
	return &discountrequirementRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *discountrequirementRepository) CreateMany(c context.Context, items []domain.DiscountRequirement) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *discountrequirementRepository) Create(c context.Context, discountrequirement *domain.DiscountRequirement) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, discountrequirement)

	return err
}

func (ur *discountrequirementRepository) Update(c context.Context, discountrequirement *domain.DiscountRequirement) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": discountrequirement.ID}
	update := bson.M{
		"$set": discountrequirement,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *discountrequirementRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *discountrequirementRepository) Fetch(c context.Context) ([]domain.DiscountRequirement, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var discountrequirements []domain.DiscountRequirement

	err = cursor.All(c, &discountrequirements)
	if discountrequirements == nil {
		return []domain.DiscountRequirement{}, err
	}

	return discountrequirements, err
}

func (tr *discountrequirementRepository) FetchByID(c context.Context, discountrequirementID string) (domain.DiscountRequirement, error) {
	collection := tr.database.Collection(tr.collection)

	var discountrequirement domain.DiscountRequirement

	idHex, err := primitive.ObjectIDFromHex(discountrequirementID)
	if err != nil {
		return discountrequirement, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&discountrequirement)
	return discountrequirement, err
}
