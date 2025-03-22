package repository

import (
	"context"

	domain "earnforglance/server/domain/localization"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type localestringresourceRepository struct {
	database   mongo.Database
	collection string
}

func NewLocaleStringResourceRepository(db mongo.Database, collection string) domain.LocaleStringResourceRepository {
	return &localestringresourceRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *localestringresourceRepository) Create(c context.Context, localestringresource *domain.LocaleStringResource) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, localestringresource)

	return err
}

func (ur *localestringresourceRepository) Update(c context.Context, localestringresource *domain.LocaleStringResource) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": localestringresource.ID}
	update := bson.M{
		"$set": localestringresource,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *localestringresourceRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *localestringresourceRepository) Fetch(c context.Context) ([]domain.LocaleStringResource, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var localestringresources []domain.LocaleStringResource

	err = cursor.All(c, &localestringresources)
	if localestringresources == nil {
		return []domain.LocaleStringResource{}, err
	}

	return localestringresources, err
}

func (tr *localestringresourceRepository) FetchByID(c context.Context, localestringresourceID string) (domain.LocaleStringResource, error) {
	collection := tr.database.Collection(tr.collection)

	var localestringresource domain.LocaleStringResource

	idHex, err := primitive.ObjectIDFromHex(localestringresourceID)
	if err != nil {
		return localestringresource, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&localestringresource)
	return localestringresource, err
}
