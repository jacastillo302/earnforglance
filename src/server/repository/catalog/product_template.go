package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type producttemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewProductTemplateRepository(db mongo.Database, collection string) domain.ProductTemplateRepository {
	return &producttemplateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *producttemplateRepository) CreateMany(c context.Context, items []domain.ProductTemplate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *producttemplateRepository) Create(c context.Context, producttemplate *domain.ProductTemplate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, producttemplate)

	return err
}

func (ur *producttemplateRepository) Update(c context.Context, producttemplate *domain.ProductTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": producttemplate.ID}
	update := bson.M{
		"$set": producttemplate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *producttemplateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *producttemplateRepository) Fetch(c context.Context) ([]domain.ProductTemplate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var producttemplates []domain.ProductTemplate

	err = cursor.All(c, &producttemplates)
	if producttemplates == nil {
		return []domain.ProductTemplate{}, err
	}

	return producttemplates, err
}

func (tr *producttemplateRepository) FetchByID(c context.Context, producttemplateID string) (domain.ProductTemplate, error) {
	collection := tr.database.Collection(tr.collection)

	var producttemplate domain.ProductTemplate

	idHex, err := primitive.ObjectIDFromHex(producttemplateID)
	if err != nil {
		return producttemplate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&producttemplate)
	return producttemplate, err
}
