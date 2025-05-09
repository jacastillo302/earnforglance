package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type bestSellersReportLineRepository struct {
	database   mongo.Database
	collection string
}

func NewBestSellersReportLineRepository(db mongo.Database, collection string) domain.BestSellersReportLineRepository {
	return &bestSellersReportLineRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *bestSellersReportLineRepository) CreateMany(c context.Context, items []domain.BestSellersReportLine) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *bestSellersReportLineRepository) Create(c context.Context, bestSellersReportLine *domain.BestSellersReportLine) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, bestSellersReportLine)

	return err
}

func (ur *bestSellersReportLineRepository) Update(c context.Context, bestSellersReportLine *domain.BestSellersReportLine) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": bestSellersReportLine.ProductID}
	update := bson.M{
		"$set": bestSellersReportLine,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *bestSellersReportLineRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *bestSellersReportLineRepository) Fetch(c context.Context) ([]domain.BestSellersReportLine, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var bestSellersReportLines []domain.BestSellersReportLine

	err = cursor.All(c, &bestSellersReportLines)
	if bestSellersReportLines == nil {
		return []domain.BestSellersReportLine{}, err
	}

	return bestSellersReportLines, err
}

func (tr *bestSellersReportLineRepository) FetchByID(c context.Context, bestSellersReportLineID string) (domain.BestSellersReportLine, error) {
	collection := tr.database.Collection(tr.collection)

	var bestSellersReportLine domain.BestSellersReportLine

	idHex, err := bson.ObjectIDFromHex(bestSellersReportLineID)
	if err != nil {
		return bestSellersReportLine, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&bestSellersReportLine)
	return bestSellersReportLine, err
}
