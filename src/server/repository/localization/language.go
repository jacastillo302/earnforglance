package repository

import (
	"context"

	domain "earnforglance/server/domain/localization"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type languageRepository struct {
	database   mongo.Database
	collection string
}

func NewLanguageRepository(db mongo.Database, collection string) domain.LanguageRepository {
	return &languageRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *languageRepository) Create(c context.Context, language *domain.Language) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, language)

	return err
}

func (ur *languageRepository) Update(c context.Context, language *domain.Language) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": language.ID}
	update := bson.M{
		"$set": language,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *languageRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *languageRepository) Fetch(c context.Context) ([]domain.Language, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var languages []domain.Language

	err = cursor.All(c, &languages)
	if languages == nil {
		return []domain.Language{}, err
	}

	return languages, err
}

func (tr *languageRepository) FetchByID(c context.Context, languageID string) (domain.Language, error) {
	collection := tr.database.Collection(tr.collection)

	var language domain.Language

	idHex, err := primitive.ObjectIDFromHex(languageID)
	if err != nil {
		return language, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&language)
	return language, err
}
