package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type robotstxtsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewRobotsTxtSettingsRepository(db mongo.Database, collection string) domain.RobotsTxtSettingsRepository {
	return &robotstxtsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *robotstxtsettingsRepository) CreateMany(c context.Context, items []domain.RobotsTxtSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *robotstxtsettingsRepository) Create(c context.Context, robotstxtsettings *domain.RobotsTxtSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, robotstxtsettings)

	return err
}

func (ur *robotstxtsettingsRepository) Update(c context.Context, robotstxtsettings *domain.RobotsTxtSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": robotstxtsettings.ID}
	update := bson.M{
		"$set": robotstxtsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *robotstxtsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *robotstxtsettingsRepository) Fetch(c context.Context) ([]domain.RobotsTxtSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var robotstxtsettingss []domain.RobotsTxtSettings

	err = cursor.All(c, &robotstxtsettingss)
	if robotstxtsettingss == nil {
		return []domain.RobotsTxtSettings{}, err
	}

	return robotstxtsettingss, err
}

func (tr *robotstxtsettingsRepository) FetchByID(c context.Context, robotstxtsettingsID string) (domain.RobotsTxtSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var robotstxtsettings domain.RobotsTxtSettings

	idHex, err := primitive.ObjectIDFromHex(robotstxtsettingsID)
	if err != nil {
		return robotstxtsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&robotstxtsettings)
	return robotstxtsettings, err
}
