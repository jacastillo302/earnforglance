package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type categorytemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewCategoryTemplateRepository(db mongo.Database, collection string) domain.CategoryTemplateRepository {
	return &categorytemplateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *categorytemplateRepository) CreateMany(c context.Context, items []domain.CategoryTemplate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *categorytemplateRepository) Create(c context.Context, categorytemplate *domain.CategoryTemplate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, categorytemplate)

	return err
}

func (ur *categorytemplateRepository) Update(c context.Context, categorytemplate *domain.CategoryTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": categorytemplate.ID}
	update := bson.M{
		"$set": categorytemplate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *categorytemplateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *categorytemplateRepository) Fetch(c context.Context) ([]domain.CategoryTemplate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var categorytemplates []domain.CategoryTemplate

	err = cursor.All(c, &categorytemplates)
	if categorytemplates == nil {
		return []domain.CategoryTemplate{}, err
	}

	return categorytemplates, err
}

func (tr *categorytemplateRepository) FetchByID(c context.Context, categorytemplateID string) (domain.CategoryTemplate, error) {
	collection := tr.database.Collection(tr.collection)

	var categorytemplate domain.CategoryTemplate

	idHex, err := bson.ObjectIDFromHex(categorytemplateID)
	if err != nil {
		return categorytemplate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&categorytemplate)
	return categorytemplate, err
}
