package repository

import (
	"context"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type countryRepository struct {
	database   mongo.Database
	collection string
}

func NewCountryRepository(db mongo.Database, collection string) domain.CountryRepository {
	return &countryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *countryRepository) CreateMany(c context.Context, items []domain.Country) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *countryRepository) Create(c context.Context, country *domain.Country) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, country)

	return err
}

func (ur *countryRepository) Update(c context.Context, country *domain.Country) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": country.ID}
	update := bson.M{
		"$set": country,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *countryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *countryRepository) Fetch(c context.Context) ([]domain.Country, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var countries []domain.Country

	err = cursor.All(c, &countries)
	if countries == nil {
		return []domain.Country{}, err
	}

	return countries, err
}

func (tr *countryRepository) FetchByID(c context.Context, countryID string) (domain.Country, error) {
	collection := tr.database.Collection(tr.collection)

	var country domain.Country

	idHex, err := bson.ObjectIDFromHex(countryID)
	if err != nil {
		return country, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&country)
	return country, err
}
