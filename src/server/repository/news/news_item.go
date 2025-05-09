package repository

import (
	"context"

	domain "earnforglance/server/domain/news"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type newsitemRepository struct {
	database   mongo.Database
	collection string
}

func NewNewsItemRepository(db mongo.Database, collection string) domain.NewsItemRepository {
	return &newsitemRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *newsitemRepository) CreateMany(c context.Context, items []domain.NewsItem) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *newsitemRepository) Create(c context.Context, newsitem *domain.NewsItem) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, newsitem)

	return err
}

func (ur *newsitemRepository) Update(c context.Context, newsitem *domain.NewsItem) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": newsitem.ID}
	update := bson.M{
		"$set": newsitem,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *newsitemRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *newsitemRepository) Fetch(c context.Context) ([]domain.NewsItem, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var newsitems []domain.NewsItem

	err = cursor.All(c, &newsitems)
	if newsitems == nil {
		return []domain.NewsItem{}, err
	}

	return newsitems, err
}

func (tr *newsitemRepository) FetchByID(c context.Context, newsitemID string) (domain.NewsItem, error) {
	collection := tr.database.Collection(tr.collection)

	var newsitem domain.NewsItem

	idHex, err := bson.ObjectIDFromHex(newsitemID)
	if err != nil {
		return newsitem, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&newsitem)
	return newsitem, err
}
