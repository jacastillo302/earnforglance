package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type salesummaryreportlineRepository struct {
	database   mongo.Database
	collection string
}

func NewSalesSummaryReportLineRepository(db mongo.Database, collection string) domain.SalesSummaryReportLineRepository {
	return &salesummaryreportlineRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *salesummaryreportlineRepository) CreateMany(c context.Context, items []domain.SalesSummaryReportLine) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *salesummaryreportlineRepository) Create(c context.Context, salesummaryreportline *domain.SalesSummaryReportLine) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, salesummaryreportline)

	return err
}

func (ur *salesummaryreportlineRepository) Update(c context.Context, salesummaryreportline *domain.SalesSummaryReportLine) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": salesummaryreportline.ID}
	update := bson.M{
		"$set": salesummaryreportline,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *salesummaryreportlineRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *salesummaryreportlineRepository) Fetch(c context.Context) ([]domain.SalesSummaryReportLine, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var salesummaryreportlines []domain.SalesSummaryReportLine

	err = cursor.All(c, &salesummaryreportlines)
	if salesummaryreportlines == nil {
		return []domain.SalesSummaryReportLine{}, err
	}

	return salesummaryreportlines, err
}

func (tr *salesummaryreportlineRepository) FetchByID(c context.Context, salesummaryreportlineID string) (domain.SalesSummaryReportLine, error) {
	collection := tr.database.Collection(tr.collection)

	var salesummaryreportline domain.SalesSummaryReportLine

	idHex, err := bson.ObjectIDFromHex(salesummaryreportlineID)
	if err != nil {
		return salesummaryreportline, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&salesummaryreportline)
	return salesummaryreportline, err
}
