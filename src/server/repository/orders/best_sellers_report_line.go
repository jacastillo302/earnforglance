package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (ur *bestSellersReportLineRepository) Delete(c context.Context, bestSellersReportLine *domain.BestSellersReportLine) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": bestSellersReportLine.ProductID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *bestSellersReportLineRepository) Fetch(c context.Context) ([]domain.BestSellersReportLine, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := primitive.ObjectIDFromHex(bestSellersReportLineID)
	if err != nil {
		return bestSellersReportLine, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&bestSellersReportLine)
	return bestSellersReportLine, err
}
