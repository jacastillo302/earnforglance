package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productpictureRepository struct {
	database   mongo.Database
	collection string
}

func NewProductPictureRepository(db mongo.Database, collection string) domain.ProductPictureRepository {
	return &productpictureRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productpictureRepository) CreateMany(c context.Context, items []domain.ProductPicture) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productpictureRepository) Create(c context.Context, productpicture *domain.ProductPicture) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productpicture)

	return err
}

func (ur *productpictureRepository) Update(c context.Context, productpicture *domain.ProductPicture) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productpicture.ID}
	update := bson.M{
		"$set": productpicture,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *productpictureRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productpictureRepository) Fetch(c context.Context) ([]domain.ProductPicture, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productpictures []domain.ProductPicture

	err = cursor.All(c, &productpictures)
	if productpictures == nil {
		return []domain.ProductPicture{}, err
	}

	return productpictures, err
}

func (tr *productpictureRepository) FetchByID(c context.Context, productpictureID string) (domain.ProductPicture, error) {
	collection := tr.database.Collection(tr.collection)

	var productpicture domain.ProductPicture

	idHex, err := primitive.ObjectIDFromHex(productpictureID)
	if err != nil {
		return productpicture, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productpicture)
	return productpicture, err
}
