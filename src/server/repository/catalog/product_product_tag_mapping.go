package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productproducttagmappingRepository struct {
	database   mongo.Database
	collection string
}

func NewProductProductTagMappingRepository(db mongo.Database, collection string) domain.ProductProductTagMappingRepository {
	return &productproducttagmappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productproducttagmappingRepository) Create(c context.Context, productproducttagmapping *domain.ProductProductTagMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productproducttagmapping)

	return err
}

func (ur *productproducttagmappingRepository) Update(c context.Context, productproducttagmapping *domain.ProductProductTagMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": productproducttagmapping.ID}
	update := bson.M{
		"$set": productproducttagmapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *productproducttagmappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productproducttagmappingRepository) Fetch(c context.Context) ([]domain.ProductProductTagMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productproducttagmappings []domain.ProductProductTagMapping

	err = cursor.All(c, &productproducttagmappings)
	if productproducttagmappings == nil {
		return []domain.ProductProductTagMapping{}, err
	}

	return productproducttagmappings, err
}

func (tr *productproducttagmappingRepository) FetchByID(c context.Context, productproducttagmappingID string) (domain.ProductProductTagMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var productproducttagmapping domain.ProductProductTagMapping

	idHex, err := primitive.ObjectIDFromHex(productproducttagmappingID)
	if err != nil {
		return productproducttagmapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productproducttagmapping)
	return productproducttagmapping, err
}
