package repository

import (
	"context"
	"crypto/rand"
	attributes "earnforglance/server/domain/attributes"
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

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type customerRepository struct {
	newsRepository domain.NewsLetterRepository
	database       mongo.Database
	collection     string
}

func NewCustomerRepository(db mongo.Database, collection string, news domain.NewsLetterRepository) domain.CustomerRepository {
	return &customerRepository{
		newsRepository: news,
		database:       db,
		collection:     collection,
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

	langID, _ := GetLangugaByCode(c, sigin.Lang, nr.database.Collection(localization.CollectionLanguage))

	sigin.Customer.LanguageID = &langID.ID

	isRegistered, sMessage, err := IsCustomerRegistered(nr, c, sigin)

	if err != nil {
		return result, err
	}

	if isRegistered.Result {
		result.Message = sMessage
		result.Result = isRegistered.Result
		return result, err
	}

	newCustomer, err := NewCustomer(c, sigin.Customer, nr.database.Collection(customers.CollectionCustomer))
	if err != nil {
		return result, err
	}

	if newCustomer.ID == bson.NilObjectID {
		return result, fmt.Errorf("customer id is nil")
	}

	sigin.Customer = newCustomer

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

	if getPw {
		setPw, err := NewCustomerPassword(c, customerPassword[0], nr.database.Collection(customers.CollectionCustomerPassword))
		if err != nil || !setPw {
			return result, err
		}
	}

	SubucribeNews(nr, c, sigin)
	SetCustomerAttributes(nr, c, sigin, newCustomer)
	SetCustomerAddress(nr, c, sigin)
	SetPrivacyConsents(nr, c, sigin)
	SendNotifications(nr, c, sigin)

	result.Message = sMessage
	return result, err
}

func SendNotifications(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (bool, error) {
	err := error(nil)
	result := false
	return result, err
}

func SetCustomerAddress(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (bool, error) {
	err := error(nil)
	result := false
	return result, err
}

func SetCustomerAttributes(nr *customerRepository, c context.Context, sigin domain.SingInRequest, newCustomer customers.Customer) (customers.Customer, error) {
	err := error(nil)
	result := newCustomer

	filter := []string{"GenderEnabled", "FirstNameEnabled", "LastNameEnabled", "DateOfBirthEnabled", "CompanyEnabled", "StreetAddressEnabled", "StreetAddress2Enabled", "ZipPostalCodeEnabled", "CityEnabled", "CountyEnabled", "CountryEnabled", "StateProvinceEnabled", "PhoneEnabled", "FaxEnabled"}

	settings, err := GetSettingByNames(c, filter, nr.database.Collection(configuration.CollectionSetting))
	if err != nil {
		return result, err
	}

	for i := range settings {
		switch settings[i].Name {
		case "GenderEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.Gender != "" {
					result.Gender = sigin.Customer.Gender
				}
			}
		case "FirstNameEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.FirstName != "" {
					result.FirstName = sigin.Customer.FirstName
				}
			}
		case "LastNameEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.LastName != "" {
					result.LastName = sigin.Customer.LastName
				}
			}
		case "DateOfBirthEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.DateOfBirth != nil {
					result.DateOfBirth = sigin.Customer.DateOfBirth
				}
			}
		case "CompanyEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.Company != "" {
					result.Company = sigin.Customer.Company
				}
			}
		case "StreetAddressEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.StreetAddress != "" {
					result.StreetAddress = sigin.Customer.StreetAddress
				}
			}
		case "StreetAddress2Enabled":
			if settings[i].Value == "true" {
				if sigin.Customer.StreetAddress2 != "" {
					result.StreetAddress2 = sigin.Customer.StreetAddress2
				}
			}
		case "ZipPostalCodeEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.ZipPostalCode != "" {
					result.ZipPostalCode = sigin.Customer.ZipPostalCode
				}
			}
		case "CityEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.City != "" {
					result.City = sigin.Customer.City
				}
			}
		case "CountyEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.County != "" {
					result.County = sigin.Customer.County
				}
			}
		case "CountryEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.CountryID != bson.NilObjectID {
					result.CountryID = sigin.Customer.CountryID
				}
			}
		case "StateProvinceEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.StateProvinceID != bson.NilObjectID {
					result.StateProvinceID = sigin.Customer.StateProvinceID
				}
			}
		case "PhoneEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.Phone != "" {
					result.Phone = sigin.Customer.Phone
				}
			}
		case "FaxEnabled":
			if settings[i].Value == "true" {
				if sigin.Customer.Fax != "" {
					result.Fax = sigin.Customer.Fax
				}
			}
		}
	}

	update, err := UpdateCustomer(c, result, nr.database.Collection(customers.CollectionCustomer))
	if err != nil {
		fmt.Println("err", err)
		return result, err
	}

	fmt.Println("update", update)
	if !update {
		return result, fmt.Errorf("failed to update customer fields")
	}

	SetCustomerCustomAttributes(nr, c, sigin, result)

	return result, err
}

func SetCustomerCustomAttributes(nr *customerRepository, c context.Context, sigin domain.SingInRequest, newCustomer customers.Customer) (bool, error) {
	result := false
	err := error(nil)

	customerRecord, err := GetRercordBySystemName(c, customers.CollectionCustomer, nr.database.Collection(security.CollectionPermissionRecord))
	if err != nil {
		return result, err
	}

	customAttributes, err := GetCustomAttributes(c, customerRecord.ID, nr.database.Collection(attributes.CollectionPermisionRecordAttribute))
	if err != nil {
		return result, err
	}

	for i := range customAttributes {
		order := 0

		baseAtributte, err := GetCustomAttribute(c, customAttributes[i].BaseAttributeID, nr.database.Collection(attributes.CollectionBaseAttribute))
		if err != nil {
			return result, err
		}

		value, ok := sigin.Attributes[baseAtributte.Name]
		if ok {
			preselected := false

			if order == 0 {
				preselected = true
			}

			order = i + 1
			SetCustomAttributeValue(c, customAttributes[i].ID, newCustomer.ID, value, preselected, order, nr.database.Collection(attributes.CollectionPermisionRecordAttributeValue))
		}
	}

	return result, err

}

func SetPrivacyConsents(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (bool, error) {
	err := error(nil)
	result := false
	return result, err
}

func SubucribeNews(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (bool, error) {

	err := error(nil)
	result := false

	if sigin.News {
		setting, err := GetSettingByName(c, "NewsletterEnabled", nr.database.Collection(configuration.CollectionSetting))
		if err != nil {
			return result, err
		}

		if setting.Value == "" {
			return result, fmt.Errorf("setting NewsletterEnabled is not set")
		}

		if setting.Value != "true" {
			return result, fmt.Errorf("setting NewsletterEnabled is diabled")
		}

		newsRequest := domain.NewsLetterRequest{
			Email:     sigin.Customer.Email,
			StoreID:   []string{sigin.Customer.RegisteredInStoreID},
			IpAddress: sigin.IpAddress,
			Lang:      sigin.Lang,
		}

		newResult, err := nr.newsRepository.NewsLetterSubscription(c, newsRequest, sigin.IpAddress)
		if err != nil {
			return result, err
		}

		if newResult.Result {
			result = true
		} else {
			locale, err := GetLocalebyName(c, newResult.Message, sigin.Lang, nr.database.Collection(localization.CollectionLocaleStringResource))
			if err != nil {
				return result, err
			}
			err = fmt.Errorf("%s", locale.ResourceValue)
			return result, err

		}
	}

	return result, err
}

func IsCustomerRegistered(nr *customerRepository, c context.Context, sigin domain.SingInRequest) (domain.SingInResponse, string, error) {

	var result domain.SingInResponse
	err := error(nil)
	result.Result = true
	sMessage := ""

	setting, err := GetSettingByName(c, "UsernamesEnabled", nr.database.Collection(configuration.CollectionSetting))
	if err != nil {
		return result, sMessage, err
	}

	if setting.Value == "" {
		return result, sMessage, fmt.Errorf("UsernamesEnabled settings is not set")
	}

	if setting.Value != "true" {
		customer, err := GetCustomerByEmail(c, sigin.Customer.Email, nr.database.Collection(customers.CollectionCustomer))
		if customer != nil && err == nil {
			sMessage = "Account.Register.Errors.EmailAlreadyExists"
		}
	} else {
		customer, err := GetCustomerByUserName(c, sigin.Customer.Username, nr.database.Collection(customers.CollectionCustomer))
		if customer != nil && err == nil {
			sMessage = "Account.Register.Errors.UsernameAlreadyExists"
		}
	}

	if sMessage == "" {
		result.Result = false
		sMessage = "Account.Login.WrongCredentials.CustomerNotExist"
	}

	locale, err := GetLocalebyName(c, sMessage, sigin.Customer.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
	if err != nil {
		return result, sMessage, err
	}

	sMessage = locale.ResourceValue

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

func GetCustomerRolMapping(c context.Context, customerID bson.ObjectID, rolID bson.ObjectID, collection mongo.Collection) ([]*customers.CustomerCustomerRoleMapping, error) {
	var roles []*customers.CustomerCustomerRoleMapping
	err := error(nil)
	cursor, err := collection.Find(c, bson.M{"customer_id": customerID, "customer_role_id": rolID})
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

	user.CreatedOnUtc = time.Now()
	user.CustomerGuid = uuid.New().String()

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

func UpdateCustomer(c context.Context, customer customers.Customer, collection mongo.Collection) (bool, error) {
	err := error(nil)

	query := bson.M{
		"_id": customer.ID,
	}

	update := bson.M{
		"$set": customer,
	}
	_, err = collection.UpdateOne(c, query, update)
	if err != nil {
		return false, err
	}

	return true, err
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
