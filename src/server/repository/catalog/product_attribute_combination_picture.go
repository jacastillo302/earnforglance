package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductAttributeCombinationPictureRepository struct {
	database   mongo.Database
	collection string
}

func NewProductAttributeCombinationPictureRepository(db mongo.Database, collection string) domain.ProductAttributeCombinationPictureRepository {
	return &ProductAttributeCombinationPictureRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *ProductAttributeCombinationPictureRepository) CreateMany(c context.Context, items []domain.ProductAttributeCombinationPicture) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *ProductAttributeCombinationPictureRepository) Create(c context.Context, ProductAttributeCombinationPicture *domain.ProductAttributeCombinationPicture) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, ProductAttributeCombinationPicture)

	return err
}

func (ur *ProductAttributeCombinationPictureRepository) Update(c context.Context, ProductAttributeCombinationPicture *domain.ProductAttributeCombinationPicture) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": ProductAttributeCombinationPicture.ID}
	update := bson.M{
		"$set": ProductAttributeCombinationPicture,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *ProductAttributeCombinationPictureRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *ProductAttributeCombinationPictureRepository) Fetch(c context.Context) ([]domain.ProductAttributeCombinationPicture, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var ProductAttributeCombinationPictures []domain.ProductAttributeCombinationPicture

	err = cursor.All(c, &ProductAttributeCombinationPictures)
	if ProductAttributeCombinationPictures == nil {
		return []domain.ProductAttributeCombinationPicture{}, err
	}

	return ProductAttributeCombinationPictures, err
}

func (tr *ProductAttributeCombinationPictureRepository) FetchByID(c context.Context, ProductAttributeCombinationPictureID string) (domain.ProductAttributeCombinationPicture, error) {
	collection := tr.database.Collection(tr.collection)

	var ProductAttributeCombinationPicture domain.ProductAttributeCombinationPicture

	idHex, err := primitive.ObjectIDFromHex(ProductAttributeCombinationPictureID)
	if err != nil {
		return ProductAttributeCombinationPicture, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&ProductAttributeCombinationPicture)
	return ProductAttributeCombinationPicture, err
}
