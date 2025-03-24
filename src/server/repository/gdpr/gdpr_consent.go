package repository

import (
	"context"

	domain "earnforglance/server/domain/gdpr"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type gdprconsentRepository struct {
	database   mongo.Database
	collection string
}

func NewGdprConsentRepository(db mongo.Database, collection string) domain.GdprConsentRepository {
	return &gdprconsentRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *gdprconsentRepository) CreateMany(c context.Context, items []domain.GdprConsent) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *gdprconsentRepository) Create(c context.Context, gdprconsent *domain.GdprConsent) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, gdprconsent)

	return err
}

func (ur *gdprconsentRepository) Update(c context.Context, gdprconsent *domain.GdprConsent) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": gdprconsent.ID}
	update := bson.M{
		"$set": gdprconsent,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *gdprconsentRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *gdprconsentRepository) Fetch(c context.Context) ([]domain.GdprConsent, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var gdprconsents []domain.GdprConsent

	err = cursor.All(c, &gdprconsents)
	if gdprconsents == nil {
		return []domain.GdprConsent{}, err
	}

	return gdprconsents, err
}

func (tr *gdprconsentRepository) FetchByID(c context.Context, gdprconsentID string) (domain.GdprConsent, error) {
	collection := tr.database.Collection(tr.collection)

	var gdprconsent domain.GdprConsent

	idHex, err := primitive.ObjectIDFromHex(gdprconsentID)
	if err != nil {
		return gdprconsent, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&gdprconsent)
	return gdprconsent, err
}
