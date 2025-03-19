package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bestcustomerreportlineRepository struct {
	database   mongo.Database
	collection string
}

func NewBestCustomerReportLineRepository(db mongo.Database, collection string) domain.BestCustomerReportLineRepository {
	return &bestcustomerreportlineRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *bestcustomerreportlineRepository) Create(c context.Context, bestcustomerreportline *domain.BestCustomerReportLine) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, bestcustomerreportline)

	return err
}

func (ur *bestcustomerreportlineRepository) Update(c context.Context, bestcustomerreportline *domain.BestCustomerReportLine) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": bestcustomerreportline.ID}
	update := bson.M{
		"$set": bestcustomerreportline,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *bestcustomerreportlineRepository) Delete(c context.Context, bestcustomerreportline *domain.BestCustomerReportLine) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": bestcustomerreportline.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *bestcustomerreportlineRepository) Fetch(c context.Context) ([]domain.BestCustomerReportLine, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var bestcustomerreportlines []domain.BestCustomerReportLine

	err = cursor.All(c, &bestcustomerreportlines)
	if bestcustomerreportlines == nil {
		return []domain.BestCustomerReportLine{}, err
	}

	return bestcustomerreportlines, err
}

func (tr *bestcustomerreportlineRepository) FetchByID(c context.Context, bestcustomerreportlineID string) (domain.BestCustomerReportLine, error) {
	collection := tr.database.Collection(tr.collection)

	var bestcustomerreportline domain.BestCustomerReportLine

	idHex, err := primitive.ObjectIDFromHex(bestcustomerreportlineID)
	if err != nil {
		return bestcustomerreportline, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&bestcustomerreportline)
	return bestcustomerreportline, err
}
