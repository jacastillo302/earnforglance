package repository

import (
	"context"

	domain "earnforglance/server/domain/affiliate"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type affiliateRepository struct {
	database   mongo.Database
	collection string
}

func NewAffiliateRepository(db mongo.Database, collection string) domain.AffiliateRepository {
	return &affiliateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *affiliateRepository) Create(c context.Context, affiliate *domain.Affiliate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, affiliate)

	return err
}

func (ur *affiliateRepository) CreateMany(c context.Context, items []domain.Affiliate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *affiliateRepository) Update(c context.Context, affiliate *domain.Affiliate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": affiliate.ID}
	update := bson.M{
		"$set": affiliate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *affiliateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *affiliateRepository) Fetch(c context.Context) ([]domain.Affiliate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var affiliates []domain.Affiliate

	err = cursor.All(c, &affiliates)
	if affiliates == nil {
		return []domain.Affiliate{}, err
	}

	return affiliates, err
}

func (tr *affiliateRepository) FetchByID(c context.Context, affiliateID string) (domain.Affiliate, error) {
	collection := tr.database.Collection(tr.collection)

	var affiliate domain.Affiliate

	idHex, err := bson.ObjectIDFromHex(affiliateID)
	if err != nil {
		return affiliate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&affiliate)
	return affiliate, err
}
