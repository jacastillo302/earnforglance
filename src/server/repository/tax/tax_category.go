package repository

import (
	"context"

	domain "earnforglance/server/domain/tax"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taxcategoryRepository struct {
	database   mongo.Database
	collection string
}

func NewTaxCategoryRepository(db mongo.Database, collection string) domain.TaxCategoryRepository {
	return &taxcategoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *taxcategoryRepository) Create(c context.Context, taxcategory *domain.TaxCategory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, taxcategory)

	return err
}

func (ur *taxcategoryRepository) Update(c context.Context, taxcategory *domain.TaxCategory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": taxcategory.ID}
	update := bson.M{
		"$set": taxcategory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *taxcategoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *taxcategoryRepository) Fetch(c context.Context) ([]domain.TaxCategory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var taxcategories []domain.TaxCategory

	err = cursor.All(c, &taxcategories)
	if taxcategories == nil {
		return []domain.TaxCategory{}, err
	}

	return taxcategories, err
}

func (tr *taxcategoryRepository) FetchByID(c context.Context, taxcategoryID string) (domain.TaxCategory, error) {
	collection := tr.database.Collection(tr.collection)

	var taxcategory domain.TaxCategory

	idHex, err := primitive.ObjectIDFromHex(taxcategoryID)
	if err != nil {
		return taxcategory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&taxcategory)
	return taxcategory, err
}
