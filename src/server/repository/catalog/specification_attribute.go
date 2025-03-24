package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type specificationattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewSpecificationAttributeRepository(db mongo.Database, collection string) domain.SpecificationAttributeRepository {
	return &specificationattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *specificationattributeRepository) CreateMany(c context.Context, items []domain.SpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *specificationattributeRepository) Create(c context.Context, specificationattribute *domain.SpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, specificationattribute)

	return err
}

func (ur *specificationattributeRepository) Update(c context.Context, specificationattribute *domain.SpecificationAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": specificationattribute.ID}
	update := bson.M{
		"$set": specificationattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *specificationattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *specificationattributeRepository) Fetch(c context.Context) ([]domain.SpecificationAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var specificationattributes []domain.SpecificationAttribute

	err = cursor.All(c, &specificationattributes)
	if specificationattributes == nil {
		return []domain.SpecificationAttribute{}, err
	}

	return specificationattributes, err
}

func (tr *specificationattributeRepository) FetchByID(c context.Context, specificationattributeID string) (domain.SpecificationAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var specificationattribute domain.SpecificationAttribute

	idHex, err := primitive.ObjectIDFromHex(specificationattributeID)
	if err != nil {
		return specificationattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&specificationattribute)
	return specificationattribute, err
}
