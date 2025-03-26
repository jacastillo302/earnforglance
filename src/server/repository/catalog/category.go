package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type categoryRepository struct {
	database   mongo.Database
	collection string
}

func NewCategoryRepository(db mongo.Database, collection string) domain.CategoryRepository {
	return &categoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *categoryRepository) CreateMany(c context.Context, items []domain.Category) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *categoryRepository) Create(c context.Context, category *domain.Category) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, category)

	return err
}

func (ur *categoryRepository) Update(c context.Context, category *domain.Category) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": category.ID}
	update := bson.M{
		"$set": category,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *categoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *categoryRepository) Fetch(c context.Context) ([]domain.Category, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var categories []domain.Category

	err = cursor.All(c, &categories)
	if categories == nil {
		return []domain.Category{}, err
	}

	return categories, err
}

func (tr *categoryRepository) FetchByID(c context.Context, categoryID string) (domain.Category, error) {
	collection := tr.database.Collection(tr.collection)

	var category domain.Category

	idHex, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return category, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&category)
	return category, err
}
