package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type crosssellproductRepository struct {
	database   mongo.Database
	collection string
}

func NewCrossSellProductRepository(db mongo.Database, collection string) domain.CrossSellProductRepository {
	return &crosssellproductRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *crosssellproductRepository) Create(c context.Context, crosssellproduct *domain.CrossSellProduct) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, crosssellproduct)

	return err
}

func (ur *crosssellproductRepository) Update(c context.Context, crosssellproduct *domain.CrossSellProduct) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": crosssellproduct.ID}
	update := bson.M{
		"$set": crosssellproduct,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *crosssellproductRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *crosssellproductRepository) Fetch(c context.Context) ([]domain.CrossSellProduct, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var crosssellproducts []domain.CrossSellProduct

	err = cursor.All(c, &crosssellproducts)
	if crosssellproducts == nil {
		return []domain.CrossSellProduct{}, err
	}

	return crosssellproducts, err
}

func (tr *crosssellproductRepository) FetchByID(c context.Context, crosssellproductID string) (domain.CrossSellProduct, error) {
	collection := tr.database.Collection(tr.collection)

	var crosssellproduct domain.CrossSellProduct

	idHex, err := primitive.ObjectIDFromHex(crosssellproductID)
	if err != nil {
		return crosssellproduct, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&crosssellproduct)
	return crosssellproduct, err
}
