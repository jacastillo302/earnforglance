package repository

import (
	"context"

	domain "earnforglance/server/domain/attributes"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type baseattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewBaseAttributeRepository(db mongo.Database, collection string) domain.BaseAttributeRepository {
	return &baseattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *baseattributeRepository) CreateMany(c context.Context, items []domain.BaseAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *baseattributeRepository) Create(c context.Context, baseattribute *domain.BaseAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, baseattribute)

	return err
}

func (ur *baseattributeRepository) Update(c context.Context, baseattribute *domain.BaseAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": baseattribute.ID}
	update := bson.M{
		"$set": baseattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *baseattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *baseattributeRepository) Fetch(c context.Context) ([]domain.BaseAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var baseattributes []domain.BaseAttribute

	err = cursor.All(c, &baseattributes)
	if baseattributes == nil {
		return []domain.BaseAttribute{}, err
	}

	return baseattributes, err
}

func (tr *baseattributeRepository) FetchByID(c context.Context, baseattributeID string) (domain.BaseAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var baseattribute domain.BaseAttribute

	idHex, err := bson.ObjectIDFromHex(baseattributeID)
	if err != nil {
		return baseattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&baseattribute)
	return baseattribute, err
}
