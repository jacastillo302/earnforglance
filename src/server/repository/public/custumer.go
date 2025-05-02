package repository

import (
	"context"
	"crypto/rand"
	configuration "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	seo "earnforglance/server/domain/seo"
	stores "earnforglance/server/domain/stores"
	service "earnforglance/server/service/customers"
	"earnforglance/server/service/data/mongo"
	encryption "earnforglance/server/service/security"
	"encoding/base64"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type customerRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerRepository(db mongo.Database, collection string) domain.CustomerRepository {
	return &customerRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *customerRepository) SingIn(c context.Context, sigin domain.SingInRequest) (domain.SingInResponse, error) {
	var result domain.SingInResponse
	err := error(nil)

	result, err = PrepareSingIn(cr, c, sigin)

	return result, err
}

func PrepareSingIn(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (domain.SingInResponse, error) {
	var result domain.SingInResponse
	err := error(nil)

	setting, err := GetSettingByName(c, "UserRegistrationType", nr.database.Collection(configuration.CollectionSetting))
	if err != nil {
		return result, err
	}

	if setting.Value == "" {
		return result, fmt.Errorf("UserRegistrationType is not set")
	}

	if setting.Value == fmt.Sprintf("%d", customers.Disabled) {
		return result, fmt.Errorf("registration is disabled")
	}

	var store *stores.Store
	store = nil

	if sigin.Customer.RegisteredInStoreID != "" {
		storeID, err := bson.ObjectIDFromHex(sigin.Customer.RegisteredInStoreID)
		if err != nil {
			return result, err
		}

		if storeID == bson.NilObjectID {
			return result, fmt.Errorf("store id is nil")
		}

		store, err = GetStoreByID(c, storeID, nr.database.Collection(stores.CollectionStore))
		if err != nil {
			return result, err
		}
	} else {
		store, err = GetDefaultStore(c, nr.database.Collection(stores.CollectionStore))
		if err != nil {
			return result, err
		}
	}

	if store == nil {
		return result, fmt.Errorf("store not found")
	}

	if store.Deleted {
		return result, fmt.Errorf("store is deleted")
	}

	sigin.Customer.RegisteredInStoreID = store.ID.Hex()

	isRegistered, sMessage, err := IsCustomerRegistered(nr, c, sigin)
	if err != nil {
		return result, err
	}

	if isRegistered.Result {

		if sMessage != "" {
			locale, err := GetLocalebyName(c, sMessage, sigin.Lang, nr.database.Collection(localization.CollectionLocaleStringResource))
			if err != nil {
				return result, err
			}
			return result, fmt.Errorf("%s", locale.ResourceValue)
		}

		return result, err
	}

	sigin.Customer.CreatedOnUtc = time.Now()
	newCustomer, err := NewCustomer(c, sigin.Customer, nr.database.Collection(customers.CollectionCustomer))
	if err != nil {
		return result, err
	}
	if newCustomer.ID == bson.NilObjectID {
		return result, fmt.Errorf("customer id is nil")
	}

	sigin.Customer.ID = newCustomer.ID
	role, err := GetCustomerRolByName(c, customers.CustomerDefaults.GuestsRoleName, nr.database.Collection(customers.CollectionCustomerRole))
	if err != nil || role == nil {
		return result, fmt.Errorf("rol no found")
	}

	_, err = NewCustomerRol(c, newCustomer.ID, role.ID, nr.database.Collection(customers.CollectionCustomerCustomerRoleMapping))
	if err != nil {
		return result, err
	}

	getPw, customerPassword := GetCustomerPassword(newCustomer.ID, sigin.Password)
	if !getPw {
		return result, fmt.Errorf("password not generated")
	}

	setPw, err := NewCustomerPassword(c, customerPassword[0], nr.database.Collection(customers.CollectionCustomerPassword))
	if err != nil {
		return result, err
	}
	if !setPw {
		return result, fmt.Errorf("password not set")
	}
	if customerPassword[0].ID == bson.NilObjectID {
		return result, fmt.Errorf("password id is nil")
	}

	fmt.Println("sigin:", sigin.Customer)
	return result, err
}

func IsCustomerRegistered(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (domain.SingInResponse, string, error) {

	var result domain.SingInResponse
	err := error(nil)
	result.Result = false
	sMessage := ""

	setting, err := GetSettingByName(c, "UsernamesEnabled", nr.database.Collection(configuration.CollectionSetting))
	if err != nil {
		return result, sMessage, err
	}

	if setting.Value == "" {
		return result, sMessage, fmt.Errorf("UsernamesEnabled settings is not set")
	}

	var customer *customers.Customer

	if setting.Value != "true" {
		sMessage = "Account.Register.Errors.EmailAlreadyExists"
		customer, err = GetCustomerByEmail(c, sigin.Customer.Email, nr.database.Collection(customers.CollectionCustomer))
		if err != nil {
			return result, sMessage, err
		}
	} else {
		sMessage = "Account.Register.Errors.UsernameAlreadyExists"
		customer, err = GetCustomerByUserName(c, sigin.Customer.Username, nr.database.Collection(customers.CollectionCustomer))
		if err != nil {
			return result, sMessage, err
		}
	}

	if customer != nil {

		result.Result = true

		if customer.Deleted {
			return result, sMessage, fmt.Errorf("customer is deleted")
		}

		if customer.RegisteredInStoreID != sigin.Customer.RegisteredInStoreID {
			return result, sMessage, fmt.Errorf("customer is not registered in this store")
		}

		typeResult, err := GetCustomerExist(nr, c, customer)
		if err != nil {
			return result, sMessage, err
		}
		switch typeResult {
		case customers.LockedOut:
			sMessage = "Account.Login.WrongCredentials.LockedOut"
		case customers.Deleted:
			sMessage = "Account.Login.WrongCredentials.Deleted"
		case customers.NotRegistered:
			sMessage = "Account.Login.WrongCredentials.NotRegistered"
		case customers.CustomerNotExist:
			sMessage = "Account.Login.WrongCredentials.CustomerNotExist"
		case customers.NotActive:
			sMessage = "Account.Login.WrongCredentials.NotActive"
		case customers.WrongPassword:
			sMessage = "Account.Login.WrongCredentials"
		case customers.Successful:
			err = fmt.Errorf("%s", sMessage)
			result.Result = true
		}

		return result, sMessage, err
	}

	return result, sMessage, err
}

func GetCustomerExist(nr *customerRepository, c context.Context, customerResult *customers.Customer) (customers.CustomerLoginResults, error) {
	var result customers.CustomerLoginResults
	err := error(nil)

	if customerResult == nil {
		result = customers.CustomerNotExist
		return result, err
	}

	if customerResult.Deleted {
		result = customers.Deleted
		return result, err
	}

	currentTime := time.Now()
	if customerResult.CannotLoginUntilDateUtc != nil {
		if customerResult.CannotLoginUntilDateUtc.After(currentTime) {
			result = customers.LockedOut
			return result, err
		}
	}

	role, err := GetCustomerRolByName(c, customers.CustomerDefaults.RegisteredRoleName, nr.database.Collection(customers.CollectionCustomerRole))
	if err != nil || role == nil {
		return result, fmt.Errorf("rol no found")
	}

	maprol, err := GetCustomerRolMapping(c, customerResult.ID, role.ID, nr.database.Collection(customers.CollectionCustomerCustomerRoleMapping))
	if err != nil || maprol == nil {
		result = customers.NotRegistered
		return result, err
	}

	result = customers.Successful

	return result, err
}

func GetCustomerByUserName(c context.Context, username string, collection mongo.Collection) (*customers.Customer, error) {
	var user *customers.Customer
	err := error(nil)

	err = collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	return user, err
}

func GetCustomerByEmail(c context.Context, email string, collection mongo.Collection) (*customers.Customer, error) {
	var user *customers.Customer
	err := error(nil)

	err = collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func GetCustomerRolMapping(c context.Context, customerID bson.ObjectID, rolID bson.ObjectID, collection mongo.Collection) (*customers.CustomerCustomerRoleMapping, error) {
	var roles *customers.CustomerCustomerRoleMapping
	err := error(nil)

	cursor, err := collection.Find(c, bson.M{"customer_id": customerID, "customer_rol_id": rolID})
	if err != nil {
		return roles, err
	}

	err = cursor.All(c, &roles)
	if err != nil {
		return roles, err
	}
	defer cursor.Close(c)

	return roles, err
}

func GetCustomerRolByName(c context.Context, name string, collection mongo.Collection) (*customers.CustomerRole, error) {
	var rol *customers.CustomerRole
	err := error(nil)

	err = collection.FindOne(c, bson.M{"name": name}).Decode(&rol)
	return rol, err
}

func GetCustomerRolByID(c context.Context, rolID bson.ObjectID, collection mongo.Collection) (*customers.CustomerRole, error) {
	var rol *customers.CustomerRole
	err := error(nil)

	err = collection.FindOne(c, bson.M{"_id": rolID}).Decode(&rol)
	return rol, err
}

func NewCustomer(c context.Context, user customers.Customer, collection mongo.Collection) (customers.Customer, error) {
	err := error(nil)
	var result customers.Customer
	insertResult, err := collection.InsertOne(c, user)
	if err != nil {
		return result, err
	}
	result = user
	if oid, ok := insertResult.(bson.ObjectID); ok {
		result.ID = oid
	} else {
		return result, fmt.Errorf("failed to assert insertResult to bson.ObjectID")
	}
	return result, err
}

func NewCustomerRol(c context.Context, customerID bson.ObjectID, RolID bson.ObjectID, collection mongo.Collection) (customers.CustomerCustomerRoleMapping, error) {
	err := error(nil)
	result := customers.CustomerCustomerRoleMapping{
		CustomerID:     customerID,
		CustomerRoleID: RolID,
	}

	insertResult, err := collection.InsertOne(c, result)
	if err != nil {
		return result, err
	}
	if oid, ok := insertResult.(bson.ObjectID); ok {
		result.ID = oid
	} else {
		return result, fmt.Errorf("failed to assert insertResult to bson.ObjectID")
	}
	return result, err
}

func NewCustomerPassword(c context.Context, customerPw customers.CustomerPassword, collection mongo.Collection) (bool, error) {
	err := error(nil)

	insertResult, err := collection.InsertOne(c, customerPw)
	if err != nil {
		return false, err
	}
	if oid, ok := insertResult.(bson.ObjectID); ok {
		if oid != bson.NilObjectID {
			customerPw.ID = oid
			return true, nil
		}
	} else {
		return false, fmt.Errorf("failed to assert insertResult to bson.ObjectID")
	}

	return false, err
}

func GetCustomerPassword(customerID bson.ObjectID, psw string) (bool, []customers.CustomerPassword) {
	result := false

	buff := make([]byte, 32)
	_, err := rand.Read(buff)
	if err != nil {
		result = false
		return result, nil
	}

	saltKey := base64.StdEncoding.EncodeToString(buff)

	hashAlgorithm := service.CustomerServicesDefaults{}

	hash, err := encryption.CreatePasswordHash(psw, saltKey, hashAlgorithm.DefaultHashedPasswordFormat())
	if err != nil {
		result := false
		return result, nil
	}

	items := []customers.CustomerPassword{
		{
			ID:               bson.NewObjectID(),
			CustomerID:       customerID,
			Password:         hash,
			PasswordFormatID: int(customers.Hashed),
			PasswordSalt:     saltKey,
			CreatedOnUTC:     time.Now(),
		},
	}

	// Success response
	result = true

	return result, items
}

func (cu *customerRepository) GetSlugs(c context.Context, record string) ([]string, error) {

	urlRecord, err := GetRercordBySystemName(c, record, cu.database.Collection(security.CollectionPermissionRecord))
	if err != nil {
		return nil, err
	}

	slugs := []string{}
	urls, err := GetSlugsByRecord(c, urlRecord.ID, cu.database.Collection(seo.CollectionUrlRecord))
	if err != nil {
		return nil, err
	}

	for _, url := range urls {
		slugs = append(slugs, url.Slug)
	}

	return slugs, nil
}
