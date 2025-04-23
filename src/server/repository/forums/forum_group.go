package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type forumgroupRepository struct {
	database   mongo.Database
	collection string
}

func NewForumGroupRepository(db mongo.Database, collection string) domain.ForumGroupRepository {
	return &forumgroupRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumgroupRepository) CreateMany(c context.Context, items []domain.ForumGroup) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *forumgroupRepository) Create(c context.Context, forumgroup *domain.ForumGroup) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forumgroup)

	return err
}

func (ur *forumgroupRepository) Update(c context.Context, forumgroup *domain.ForumGroup) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumgroup.ID}
	update := bson.M{
		"$set": forumgroup,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *forumgroupRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *forumgroupRepository) Fetch(c context.Context) ([]domain.ForumGroup, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forumgroups []domain.ForumGroup

	err = cursor.All(c, &forumgroups)
	if forumgroups == nil {
		return []domain.ForumGroup{}, err
	}

	return forumgroups, err
}

func (tr *forumgroupRepository) FetchByID(c context.Context, forumgroupID string) (domain.ForumGroup, error) {
	collection := tr.database.Collection(tr.collection)

	var forumgroup domain.ForumGroup

	idHex, err := bson.ObjectIDFromHex(forumgroupID)
	if err != nil {
		return forumgroup, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forumgroup)
	return forumgroup, err
}
