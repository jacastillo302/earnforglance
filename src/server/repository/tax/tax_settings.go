package repository

import (
	"context"

	domain "earnforglance/server/domain/tax"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taxsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewTaxSettingsRepository(db mongo.Database, collection string) domain.TaxSettingsRepository {
	return &taxsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *taxsettingsRepository) Create(c context.Context, taxsettings *domain.TaxSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, taxsettings)

	return err
}

func (ur *taxsettingsRepository) Update(c context.Context, taxsettings *domain.TaxSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": taxsettings.ID}
	update := bson.M{
		"$set": taxsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *taxsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *taxsettingsRepository) Fetch(c context.Context) ([]domain.TaxSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var taxsettingss []domain.TaxSettings

	err = cursor.All(c, &taxsettingss)
	if taxsettingss == nil {
		return []domain.TaxSettings{}, err
	}

	return taxsettingss, err
}

func (tr *taxsettingsRepository) FetchByID(c context.Context, taxsettingsID string) (domain.TaxSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var taxsettings domain.TaxSettings

	idHex, err := primitive.ObjectIDFromHex(taxsettingsID)
	if err != nil {
		return taxsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&taxsettings)
	return taxsettings, err
}
