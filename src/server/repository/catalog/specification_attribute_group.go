package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type specificationattributegroupRepository struct {
	database   mongo.Database
	collection string
}

func NewSpecificationAttributeGroupRepository(db mongo.Database, collection string) domain.SpecificationAttributeGroupRepository {
	return &specificationattributegroupRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *specificationattributegroupRepository) CreateMany(c context.Context, items []domain.SpecificationAttributeGroup) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *specificationattributegroupRepository) Create(c context.Context, specificationattributegroup *domain.SpecificationAttributeGroup) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, specificationattributegroup)

	return err
}

func (ur *specificationattributegroupRepository) Update(c context.Context, specificationattributegroup *domain.SpecificationAttributeGroup) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": specificationattributegroup.ID}
	update := bson.M{
		"$set": specificationattributegroup,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *specificationattributegroupRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *specificationattributegroupRepository) Fetch(c context.Context) ([]domain.SpecificationAttributeGroup, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var specificationattributegroups []domain.SpecificationAttributeGroup

	err = cursor.All(c, &specificationattributegroups)
	if specificationattributegroups == nil {
		return []domain.SpecificationAttributeGroup{}, err
	}

	return specificationattributegroups, err
}

func (tr *specificationattributegroupRepository) FetchByID(c context.Context, specificationattributegroupID string) (domain.SpecificationAttributeGroup, error) {
	collection := tr.database.Collection(tr.collection)

	var specificationattributegroup domain.SpecificationAttributeGroup

	idHex, err := bson.ObjectIDFromHex(specificationattributegroupID)
	if err != nil {
		return specificationattributegroup, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&specificationattributegroup)
	return specificationattributegroup, err
}
