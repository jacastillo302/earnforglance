package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type picturebinaryRepository struct {
	database   mongo.Database
	collection string
}

func NewPictureBinaryRepository(db mongo.Database, collection string) domain.PictureBinaryRepository {
	return &picturebinaryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *picturebinaryRepository) CreateMany(c context.Context, items []domain.PictureBinary) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *picturebinaryRepository) Create(c context.Context, picturebinary *domain.PictureBinary) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, picturebinary)

	return err
}

func (ur *picturebinaryRepository) Update(c context.Context, picturebinary *domain.PictureBinary) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": picturebinary.ID}
	update := bson.M{
		"$set": picturebinary,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *picturebinaryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *picturebinaryRepository) Fetch(c context.Context) ([]domain.PictureBinary, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var picturebinarys []domain.PictureBinary

	err = cursor.All(c, &picturebinarys)
	if picturebinarys == nil {
		return []domain.PictureBinary{}, err
	}

	return picturebinarys, err
}

func (tr *picturebinaryRepository) FetchByID(c context.Context, picturebinaryID string) (domain.PictureBinary, error) {
	collection := tr.database.Collection(tr.collection)

	var picturebinary domain.PictureBinary

	idHex, err := primitive.ObjectIDFromHex(picturebinaryID)
	if err != nil {
		return picturebinary, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&picturebinary)
	return picturebinary, err
}
