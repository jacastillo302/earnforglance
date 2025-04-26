package repository

import (
	"context"

	setting "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewLoginRepository(db mongo.Database, collection string) domain.LoginRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *customers.Customer) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]customers.Customer, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []customers.Customer

	err = cursor.All(c, &users)
	if users == nil {
		return []customers.Customer{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (customers.Customer, error) {
	collection := ur.database.Collection(customers.CollectionCustomer)
	var user customers.Customer
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByUserName(c context.Context, usermame string) (customers.Customer, error) {
	collection := ur.database.Collection(customers.CollectionCustomer)
	var user customers.Customer
	err := collection.FindOne(c, bson.M{"username": usermame}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetPasw(c context.Context, CustumerID string) (customers.CustomerPassword, error) {
	collection := ur.database.Collection(customers.CollectionCustomerPassword)
	var user customers.CustomerPassword

	idHex, err := bson.ObjectIDFromHex(CustumerID)
	if err != nil {
		return user, err
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "created_on_utc", Value: 1}})

	cursor, err := collection.Find(c, bson.M{"customer_id": idHex}, opts)
	if err != nil {
		return user, err
	}

	var pasws []customers.CustomerPassword

	err = cursor.All(c, &pasws)
	if pasws == nil {
		return customers.CustomerPassword{}, err
	}

	user = pasws[0]

	return user, err
}

func (ur *userRepository) GetSettingByName(c context.Context, name string) (setting.Setting, error) {
	item, err := GetSettingByName(c, name, ur.database.Collection(setting.CollectionSetting))
	if err != nil {
		return item, err
	}
	return item, err
}

func (ur *userRepository) GetLangugaByCode(c context.Context, lang string) (localization.Language, error) {
	err := error(nil)
	locale, err := GetLangugaByCode(c, lang, ur.database.Collection(localization.CollectionLanguage))
	if err != nil {
		return locale, err
	}
	return locale, err
}

func (ur *userRepository) GetLocalebyName(c context.Context, name string, languageID string) (localization.LocaleStringResource, error) {
	err := error(nil)
	locale, err := GetLocalebyName(c, name, languageID, ur.database.Collection(localization.CollectionLocalizedProperty))
	if err != nil {
		return locale, err
	}
	return locale, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (customers.Customer, error) {
	collection := ur.database.Collection(ur.collection)

	var user customers.Customer

	idHex, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
