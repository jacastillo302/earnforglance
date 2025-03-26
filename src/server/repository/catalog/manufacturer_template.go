package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type manufacturertemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewManufacturerTemplateRepository(db mongo.Database, collection string) domain.ManufacturerTemplateRepository {
	return &manufacturertemplateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *manufacturertemplateRepository) CreateMany(c context.Context, items []domain.ManufacturerTemplate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *manufacturertemplateRepository) Create(c context.Context, manufacturertemplate *domain.ManufacturerTemplate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, manufacturertemplate)

	return err
}

func (ur *manufacturertemplateRepository) Update(c context.Context, manufacturertemplate *domain.ManufacturerTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": manufacturertemplate.ID}
	update := bson.M{
		"$set": manufacturertemplate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *manufacturertemplateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *manufacturertemplateRepository) Fetch(c context.Context) ([]domain.ManufacturerTemplate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var manufacturertemplates []domain.ManufacturerTemplate

	err = cursor.All(c, &manufacturertemplates)
	if manufacturertemplates == nil {
		return []domain.ManufacturerTemplate{}, err
	}

	return manufacturertemplates, err
}

func (tr *manufacturertemplateRepository) FetchByID(c context.Context, manufacturertemplateID string) (domain.ManufacturerTemplate, error) {
	collection := tr.database.Collection(tr.collection)

	var manufacturertemplate domain.ManufacturerTemplate

	idHex, err := primitive.ObjectIDFromHex(manufacturertemplateID)
	if err != nil {
		return manufacturertemplate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&manufacturertemplate)
	return manufacturertemplate, err
}
