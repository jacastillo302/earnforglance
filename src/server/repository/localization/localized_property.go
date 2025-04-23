package repository

import (
	"context"

	domain "earnforglance/server/domain/localization"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type localizedpropertyRepository struct {
	database   mongo.Database
	collection string
}

func NewLocalizedPropertyRepository(db mongo.Database, collection string) domain.LocalizedPropertyRepository {
	return &localizedpropertyRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *localizedpropertyRepository) CreateMany(c context.Context, items []domain.LocalizedProperty) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *localizedpropertyRepository) Create(c context.Context, localizedproperty *domain.LocalizedProperty) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, localizedproperty)

	return err
}

func (ur *localizedpropertyRepository) Update(c context.Context, localizedproperty *domain.LocalizedProperty) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": localizedproperty.ID}
	update := bson.M{
		"$set": localizedproperty,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *localizedpropertyRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *localizedpropertyRepository) Fetch(c context.Context) ([]domain.LocalizedProperty, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var localizedpropertys []domain.LocalizedProperty

	err = cursor.All(c, &localizedpropertys)
	if localizedpropertys == nil {
		return []domain.LocalizedProperty{}, err
	}

	return localizedpropertys, err
}

func (tr *localizedpropertyRepository) FetchByID(c context.Context, localizedpropertyID string) (domain.LocalizedProperty, error) {
	collection := tr.database.Collection(tr.collection)

	var localizedproperty domain.LocalizedProperty

	idHex, err := bson.ObjectIDFromHex(localizedpropertyID)
	if err != nil {
		return localizedproperty, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&localizedproperty)
	return localizedproperty, err
}
