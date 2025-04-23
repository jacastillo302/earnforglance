package repository

import (
	"context"
	middleware "earnforglance/server/domain/middleware"
	security "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MiddlewareRepository struct {
	database mongo.Database
}

func NewInstallRepository(db mongo.Database) middleware.MiddlewareRepository {
	return &MiddlewareRepository{
		database: db,
	}
}

func (tu *MiddlewareRepository) GetPermissionsCustumer(c context.Context, custumerID string) ([]middleware.Middleware, error) {

	collection := tu.database.Collection(security.CollectionPermissionRecordCustomerRoleMapping)

	customer, err := bson.ObjectIDFromHex(custumerID)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(c, bson.M{"customer_id": customer})
	if err != nil {
		return nil, err
	}

	var permissions []middleware.Middleware

	err = cursor.All(c, &permissions)
	if permissions == nil {
		return []middleware.Middleware{}, err
	}

	return permissions, err
}
