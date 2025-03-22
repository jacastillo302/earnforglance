package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productattributevaluepictureRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAttributeValuePictureRepository(db mongo.Database, collection string) domain.ProductAttributeValuePictureRepository {
	return &productattributevaluepictureRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productattributevaluepictureRepository) Create(c context.Context, productattributevaluepicture *domain.ProductAttributeValuePicture) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productattributevaluepicture)

	return err
}

func (ur *productattributevaluepictureRepository) Update(c context.Context, productattributevaluepicture *domain.ProductAttributeValuePicture) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productattributevaluepicture.ID}
	update := bson.M{
		"$set": productattributevaluepicture,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productattributevaluepictureRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *productattributevaluepictureRepository) Fetch(c context.Context) ([]domain.ProductAttributeValuePicture, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productattributevaluepictures []domain.ProductAttributeValuePicture

	err = cursor.All(c, &productattributevaluepictures)
	if productattributevaluepictures == nil {
		return []domain.ProductAttributeValuePicture{}, err
	}

	return productattributevaluepictures, err
}

func (tr *productattributevaluepictureRepository) FetchByID(c context.Context, productattributevaluepictureID string) (domain.ProductAttributeValuePicture, error) {
	collection := tr.database.Collection(tr.collection)

	var productattributevaluepicture domain.ProductAttributeValuePicture

	idHex, err := primitive.ObjectIDFromHex(productattributevaluepictureID)
	if err != nil {
		return productattributevaluepicture, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productattributevaluepicture)
	return productattributevaluepicture, err
}
