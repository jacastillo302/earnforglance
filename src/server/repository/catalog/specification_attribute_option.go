package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type specificationattributeoptionRepository struct {
	database   mongo.Database
	collection string
}

func NewSpecificationAttributeOptionRepository(db mongo.Database, collection string) domain.SpecificationAttributeOptionRepository {
	return &specificationattributeoptionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *specificationattributeoptionRepository) CreateMany(c context.Context, items []domain.SpecificationAttributeOption) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *specificationattributeoptionRepository) Create(c context.Context, specificationattributeoption *domain.SpecificationAttributeOption) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, specificationattributeoption)

	return err
}

func (ur *specificationattributeoptionRepository) Update(c context.Context, specificationattributeoption *domain.SpecificationAttributeOption) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": specificationattributeoption.ID}
	update := bson.M{
		"$set": specificationattributeoption,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *specificationattributeoptionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *specificationattributeoptionRepository) Fetch(c context.Context) ([]domain.SpecificationAttributeOption, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var specificationattributeoptions []domain.SpecificationAttributeOption

	err = cursor.All(c, &specificationattributeoptions)
	if specificationattributeoptions == nil {
		return []domain.SpecificationAttributeOption{}, err
	}

	return specificationattributeoptions, err
}

func (tr *specificationattributeoptionRepository) FetchByID(c context.Context, specificationattributeoptionID string) (domain.SpecificationAttributeOption, error) {
	collection := tr.database.Collection(tr.collection)

	var specificationattributeoption domain.SpecificationAttributeOption

	idHex, err := primitive.ObjectIDFromHex(specificationattributeoptionID)
	if err != nil {
		return specificationattributeoption, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&specificationattributeoption)
	return specificationattributeoption, err
}
