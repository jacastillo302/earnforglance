package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type searchtermreportlineRepository struct {
	database   mongo.Database
	collection string
}

func NewSearchTermReportLineRepository(db mongo.Database, collection string) domain.SearchTermReportLineRepository {
	return &searchtermreportlineRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *searchtermreportlineRepository) CreateMany(c context.Context, items []domain.SearchTermReportLine) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *searchtermreportlineRepository) Create(c context.Context, searchtermreportline *domain.SearchTermReportLine) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, searchtermreportline)

	return err
}

func (ur *searchtermreportlineRepository) Update(c context.Context, searchtermreportline *domain.SearchTermReportLine) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": searchtermreportline.ID}
	update := bson.M{
		"$set": searchtermreportline,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *searchtermreportlineRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *searchtermreportlineRepository) Fetch(c context.Context) ([]domain.SearchTermReportLine, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var searchtermreportlines []domain.SearchTermReportLine

	err = cursor.All(c, &searchtermreportlines)
	if searchtermreportlines == nil {
		return []domain.SearchTermReportLine{}, err
	}

	return searchtermreportlines, err
}

func (tr *searchtermreportlineRepository) FetchByID(c context.Context, searchtermreportlineID string) (domain.SearchTermReportLine, error) {
	collection := tr.database.Collection(tr.collection)

	var searchtermreportline domain.SearchTermReportLine

	idHex, err := bson.ObjectIDFromHex(searchtermreportlineID)
	if err != nil {
		return searchtermreportline, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&searchtermreportline)
	return searchtermreportline, err
}
