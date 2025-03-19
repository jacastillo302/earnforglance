package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type searchtermRepository struct {
	database   mongo.Database
	collection string
}

func NewSearchTermRepository(db mongo.Database, collection string) domain.SearchTermRepository {
	return &searchtermRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *searchtermRepository) Create(c context.Context, searchterm *domain.SearchTerm) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, searchterm)

	return err
}

func (ur *searchtermRepository) Update(c context.Context, searchterm *domain.SearchTerm) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": searchterm.ID}
	update := bson.M{
		"$set": searchterm,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *searchtermRepository) Delete(c context.Context, searchterm *domain.SearchTerm) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": searchterm.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *searchtermRepository) Fetch(c context.Context) ([]domain.SearchTerm, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var searchterms []domain.SearchTerm

	err = cursor.All(c, &searchterms)
	if searchterms == nil {
		return []domain.SearchTerm{}, err
	}

	return searchterms, err
}

func (tr *searchtermRepository) FetchByID(c context.Context, searchtermID string) (domain.SearchTerm, error) {
	collection := tr.database.Collection(tr.collection)

	var searchterm domain.SearchTerm

	idHex, err := primitive.ObjectIDFromHex(searchtermID)
	if err != nil {
		return searchterm, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&searchterm)
	return searchterm, err
}
