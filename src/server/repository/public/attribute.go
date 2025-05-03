package repository

import (
	"context"
	attributes "earnforglance/server/domain/attributes"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetCustomAttributes(c context.Context, permissionRecordID bson.ObjectID, collection mongo.Collection) ([]attributes.PermisionRecordAttribute, error) {
	var permisionRecordAttribute []attributes.PermisionRecordAttribute

	findOptions := options.Find().
		SetSort(bson.D{{Key: "DisplayOrder", Value: 1}})

	query := bson.M{
		"permission_record_id": permissionRecordID,
	}

	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return permisionRecordAttribute, err
	}

	err = cursor.All(c, &permisionRecordAttribute)
	if err != nil {
		return permisionRecordAttribute, err
	}

	return permisionRecordAttribute, err
}

func GetCustomAttribute(c context.Context, baseAttributeID bson.ObjectID, collection mongo.Collection) (attributes.BaseAttribute, error) {
	var baseAttribute attributes.BaseAttribute

	query := bson.M{
		"_id": baseAttributeID,
	}

	err := collection.FindOne(c, query).Decode(&baseAttribute)
	if err != nil {
		return baseAttribute, err
	}

	return baseAttribute, err
}

func SetCustomAttributeValue(c context.Context, permisionRecordAttributeID bson.ObjectID, recordID bson.ObjectID, value string, isPreSelected bool, order int, collection mongo.Collection) ([]attributes.PermisionRecordAttributeValue, error) {
	var permisionRecordAttributeValue []attributes.PermisionRecordAttributeValue

	findOptions := options.Find().
		SetSort(bson.D{{Key: "DisplayOrder", Value: 1}}).
		SetLimit(1)

	query := bson.M{
		"permission_record_attribute_id": permisionRecordAttributeID,
		"record_id":                      recordID,
	}

	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return permisionRecordAttributeValue, err
	}

	err = cursor.All(c, &permisionRecordAttributeValue)
	if err != nil {
		return permisionRecordAttributeValue, err
	}

	if len(permisionRecordAttributeValue) == 0 {
		collection.InsertOne(c, bson.M{
			"permission_record_attribute_id": permisionRecordAttributeID,
			"record_id":                      recordID,
			"value":                          value,
			"is_pre_selected":                isPreSelected,
			"display_order":                  order,
		})
	} else {
		for _, item := range permisionRecordAttributeValue {
			collection.UpdateOne(c, bson.M{"_id": item.ID}, bson.M{"$set": bson.M{"value": value, "is_pre_selected": isPreSelected, "display_order": order}})

		}
	}

	return permisionRecordAttributeValue, err
}
