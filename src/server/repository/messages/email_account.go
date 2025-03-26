package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type emailaccountRepository struct {
	database   mongo.Database
	collection string
}

func NewEmailAccountRepository(db mongo.Database, collection string) domain.EmailAccountRepository {
	return &emailaccountRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *emailaccountRepository) CreateMany(c context.Context, items []domain.EmailAccount) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *emailaccountRepository) Create(c context.Context, emailaccount *domain.EmailAccount) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, emailaccount)

	return err
}

func (ur *emailaccountRepository) Update(c context.Context, emailaccount *domain.EmailAccount) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": emailaccount.ID}
	update := bson.M{
		"$set": emailaccount,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *emailaccountRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *emailaccountRepository) Fetch(c context.Context) ([]domain.EmailAccount, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var emailaccounts []domain.EmailAccount

	err = cursor.All(c, &emailaccounts)
	if emailaccounts == nil {
		return []domain.EmailAccount{}, err
	}

	return emailaccounts, err
}

func (tr *emailaccountRepository) FetchByID(c context.Context, emailaccountID string) (domain.EmailAccount, error) {
	collection := tr.database.Collection(tr.collection)

	var emailaccount domain.EmailAccount

	idHex, err := primitive.ObjectIDFromHex(emailaccountID)
	if err != nil {
		return emailaccount, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&emailaccount)
	return emailaccount, err
}
