package repository

import (
	"context"
	domain "earnforglance/server/domain/logging"
	"earnforglance/server/service/data/mongo"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func InsertActivity(c context.Context, recordID bson.ObjectID, recordName string, customerID bson.ObjectID, activityLogType bson.ObjectID, comment string, ipAddress string, collection mongo.Collection) (bool, error) {

	activityLog := domain.ActivityLog{
		EntityID:          &recordID,
		EntityName:        recordName,
		ActivityLogTypeID: activityLogType,
		CustomerID:        customerID,
		Comment:           comment,
		CreatedOnUtc:      time.Now(),
		IpAddress:         ipAddress,
	}

	_, err := collection.InsertOne(c, activityLog)
	if err != nil {
		return false, err
	}

	return true, err
}

func GetActivityLogTypeBySystemKeyword(c context.Context, systemKeyword string, collection mongo.Collection) (*bson.ObjectID, error) {

	var activityLogType domain.ActivityLogType

	query := bson.M{
		"system_keyword": systemKeyword,
	}

	collection.FindOne(c, query).Decode(&activityLogType)
	if activityLogType.SystemKeyword == "" {
		return nil, fmt.Errorf("activity log type not found for system keyword: %s", systemKeyword)
	}

	return &activityLogType.ID, nil
}
