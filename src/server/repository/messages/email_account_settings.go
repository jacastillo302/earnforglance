package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type emailaccountsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewEmailAccountSettingsRepository(db mongo.Database, collection string) domain.EmailAccountSettingsRepository {
	return &emailaccountsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *emailaccountsettingsRepository) Create(c context.Context, emailaccountsettings *domain.EmailAccountSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, emailaccountsettings)

	return err
}

func (ur *emailaccountsettingsRepository) Update(c context.Context, emailaccountsettings *domain.EmailAccountSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": emailaccountsettings.ID}
	update := bson.M{
		"$set": emailaccountsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *emailaccountsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *emailaccountsettingsRepository) Fetch(c context.Context) ([]domain.EmailAccountSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var emailaccountsettingss []domain.EmailAccountSettings

	err = cursor.All(c, &emailaccountsettingss)
	if emailaccountsettingss == nil {
		return []domain.EmailAccountSettings{}, err
	}

	return emailaccountsettingss, err
}

func (tr *emailaccountsettingsRepository) FetchByID(c context.Context, emailaccountsettingsID string) (domain.EmailAccountSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var emailaccountsettings domain.EmailAccountSettings

	idHex, err := primitive.ObjectIDFromHex(emailaccountsettingsID)
	if err != nil {
		return emailaccountsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&emailaccountsettings)
	return emailaccountsettings, err
}
