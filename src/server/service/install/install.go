package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"crypto/rand"
	"encoding/base64"

	"earnforglance/server/bootstrap"
	catalog "earnforglance/server/domain/catalog"
	commons "earnforglance/server/domain/common"
	configuration "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	directory "earnforglance/server/domain/directory"
	response "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	loggings "earnforglance/server/domain/logging"
	messages "earnforglance/server/domain/messages"
	orders "earnforglance/server/domain/orders"
	tasks "earnforglance/server/domain/scheduleTasks"
	security "earnforglance/server/domain/security"
	shipping "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
	topics "earnforglance/server/domain/topics"
	tools "earnforglance/server/service/common"
	service "earnforglance/server/service/customers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InstallService struct {
	InstallUsecase response.InstallRepository
	SettingUsecase configuration.SettingRepository
	StoresUsecase  stores.StoreRepository
	Env            *bootstrap.Env
}

const (
	DefaultPathJson = "service\\data\\json"
)

func resolvePath(basePath string, relativePath string) string {
	workingDir, _ := os.Getwd()
	return workingDir + "\\" + basePath + "\\" + relativePath
}

func ReadJsonMap(filePath string) (map[string]interface{}, error) {
	// Open the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	// Parse the JSON data into a map
	var settings map[string]interface{}
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON file: %w", err)
	}

	return settings, nil
}

func InstallPermissionRecord() (response.Install, []security.PermissionRecord) {
	var result response.Install
	collection := security.CollectionPermissionRecord

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "security\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of security.PermissionRecord
	items := make([]security.PermissionRecord, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCurrency() (response.Install, []directory.Currency) {
	var result response.Install
	collection := directory.CollectionCurrency

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "directory\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]directory.Currency, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallMeasureDimension() (response.Install, []directory.MeasureDimension) {
	var result response.Install
	collection := directory.CollectionMeasureDimension

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "directory\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]directory.MeasureDimension, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallMeasureWeight() (response.Install, []directory.MeasureWeight) {
	var result response.Install
	collection := directory.CollectionMeasureWeight

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "directory\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]directory.MeasureWeight, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallTaxCategory() (response.Install, []taxes.TaxCategory) {
	var result response.Install
	collection := taxes.CollectionTaxCategory

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "tax\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of taxes.TaxCategory
	items := make([]taxes.TaxCategory, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallLanguages() (response.Install, []lang.Language) {
	var result response.Install
	collection := lang.CollectionLanguage

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "localization\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of lang.Language
	items := make([]lang.Language, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallLocaleStringResource(languageID string, culture string) (response.Install, []lang.LocaleStringResource) {
	var result response.Install
	collection := lang.CollectionLocaleStringResource

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "localization\\"+collection+"."+culture+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	type NameValue struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	}

	nameValueItems := make([]NameValue, 0)
	err = json.Unmarshal(fileData, &nameValueItems)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	ID, err := primitive.ObjectIDFromHex(languageID)
	if err != nil {
		result.Status = false
		result.Details = "Invalid languageID: " + err.Error()
		return result, nil
	}

	// Convert NameValue items to lang.LocaleStringResource
	localeStringResources := make([]lang.LocaleStringResource, 0)
	for _, item := range nameValueItems {
		localeStringResources = append(localeStringResources, lang.LocaleStringResource{
			ID:            primitive.NewObjectID(), // Generate a new ObjectID
			LanguageID:    ID,                      // Replace with actual LanguageID if available
			ResourceName:  item.Name,
			ResourceValue: item.Value,
		})
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, localeStringResources
}

func InstallStores() (response.Install, []stores.Store) {
	var result response.Install
	collection := stores.CollectionStore

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "store\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]stores.Store, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallSettings() (response.Install, []configuration.Setting) {
	var result response.Install
	collection := configuration.CollectionSetting

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "store\\"+stores.CollectionStore+".json")

	// Read the JSON file
	storeData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	stores := make([]stores.Store, 0)
	err = json.Unmarshal(storeData, &stores)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse store JSON file: " + err.Error()
		return result, nil
	}

	// Resolve the relative path
	filePath = resolvePath(DefaultPathJson, "configuration\\"+collection+".json")

	fileData, err := ReadJsonMap(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil

	}

	// Iterate over the grouped settings
	settings := make([]configuration.Setting, 0)
	for group, values := range fileData {
		switch v := values.(type) {
		case map[string]interface{}:
			for key, value := range v {
				strValue, ok := value.(string) // Type assertion to convert value to string
				if !ok {
					fmt.Printf("  Skipping key %s: value is not a string\n", key)
					continue
				}
				settings = append(settings, configuration.Setting{
					ID:      primitive.NewObjectID(),
					Name:    key,
					Value:   strValue,
					StoreID: stores[0].ID,
				})

			}
		default:
			fmt.Printf("  Unknown type for group %s\n", group)
			result.Status = false
			result.Details = "Invalid storeID: " + "Unknown type for group " + group
			return result, nil
		}
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, settings
}

func InstallCountries() (response.Install, []directory.Country) {
	var result response.Install
	collection := directory.CollectionCountry

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "directory\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]directory.Country, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallStateProvince() (response.Install, []directory.StateProvince) {
	var result response.Install
	collection := directory.CollectionStateProvince

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "directory\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of directory.StateProvince
	items := make([]directory.StateProvince, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallShippingMethod() (response.Install, []shipping.ShippingMethod) {
	var result response.Install
	collection := shipping.CollectionShippingMethod

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "shipping\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]shipping.ShippingMethod, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallDeliveryDate() (response.Install, []shipping.DeliveryDate) {
	var result response.Install
	collection := shipping.CollectionDeliveryDate

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "shipping\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of stores.Store
	items := make([]shipping.DeliveryDate, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallProductAvailabilityRange() (response.Install, []shipping.ProductAvailabilityRange) {
	var result response.Install
	collection := shipping.CollectionProductAvailabilityRange

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "shipping\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductAvailabilityRange
	items := make([]shipping.ProductAvailabilityRange, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallEmailAccount() (response.Install, []messages.EmailAccount) {
	var result response.Install
	collection := messages.CollectionEmailAccount

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "messages\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of messages.EmailAccount
	items := make([]messages.EmailAccount, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallMessageTemplate() (response.Install, []messages.MessageTemplate) {
	var result response.Install
	collection := messages.CollectionMessageTemplate

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "messages\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of messages.MessageTemplate
	items := make([]messages.MessageTemplate, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallTopicTemplate() (response.Install, []topics.TopicTemplate) {
	var result response.Install
	collection := topics.CollectionTopicTemplate

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "topics\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of topics.TopicTemplate
	items := make([]topics.TopicTemplate, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallTopic() (response.Install, []topics.Topic) {
	var result response.Install
	collection := topics.CollectionTopic

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "topics\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of topics.TopicTemplate
	items := make([]topics.Topic, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCustomerRole() (response.Install, []customers.CustomerRole) {
	var result response.Install
	collection := customers.CollectionCustomerRole

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "customers\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of customers.CustomerRole
	items := make([]customers.CustomerRole, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCustomer() (response.Install, []customers.Customer) {
	var result response.Install
	collection := customers.CollectionCustomer

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "customers\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of customers.CustomerRole
	items := make([]customers.Customer, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCustomerPassword(customerID primitive.ObjectID, psw string) (response.Install, []customers.CustomerPassword) {
	var result response.Install
	collection := customers.CollectionCustomer

	buff := make([]byte, 32) // Assuming NopCustomerServicesDefaults.PasswordSaltKeySize is 32
	_, err := rand.Read(buff)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	saltKey := base64.StdEncoding.EncodeToString(buff)

	hashAlgorithm := service.CustomerServicesDefaults{}

	data := []byte(psw)

	hash, err := tools.CreateHash(data, hashAlgorithm.DefaultHashedPasswordFormat(), hashAlgorithm.PasswordSaltKeySize())
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse  create  Hash:" + err.Error()
		return result, nil
	}

	fmt.Println("Hash:", hash)

	items := []customers.CustomerPassword{
		{
			ID:               primitive.NewObjectID(),
			CustomerID:       customerID,
			Password:         hash,
			PasswordFormatID: int(customers.Hashed),
			PasswordSalt:     saltKey,
			CreatedOnUTC:     time.Now(),
		},
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCustomerAddress() (response.Install, []commons.Address) {
	var result response.Install
	collection := commons.CollectionAddress

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "common\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of customers.CustomerRole
	items := make([]commons.Address, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCustomerAddressMapping(customerID primitive.ObjectID, addressID primitive.ObjectID) (response.Install, customers.CustomerAddressMapping) {
	var result response.Install
	collection := customers.CollectionCustomerAddressMapping

	items := customers.CustomerAddressMapping{
		CustomerID: customerID,
		AddressID:  addressID,
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCustomerCustomerRoleMapping(customerID primitive.ObjectID, rolesID []customers.CustomerRole) (response.Install, []customers.CustomerCustomerRoleMapping) {
	var result response.Install
	collection := customers.CollectionCustomerCustomerRoleMapping

	items := make([]customers.CustomerCustomerRoleMapping, 0)
	for _, item := range rolesID {
		items = append(items, customers.CustomerCustomerRoleMapping{
			ID:             primitive.NewObjectID(),
			CustomerID:     customerID,
			CustomerRoleID: item.ID,
		})
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallActivityLogType() (response.Install, []loggings.ActivityLogType) {
	var result response.Install
	collection := loggings.CollectionActivityLogType

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "logging\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of customers.CustomerRole
	items := make([]loggings.ActivityLogType, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallProductTemplate() (response.Install, []catalog.ProductTemplate) {
	var result response.Install
	collection := catalog.CollectionProductTemplate

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of catalog.ProductTemplate
	items := make([]catalog.ProductTemplate, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallCategoryTemplate() (response.Install, []catalog.CategoryTemplate) {
	var result response.Install
	collection := catalog.CollectionCategoryTemplate

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of catalog.ProductTemplate
	items := make([]catalog.CategoryTemplate, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallManufacturerTemplate() (response.Install, []catalog.ManufacturerTemplate) {
	var result response.Install
	collection := catalog.CollectionManufacturerTemplate

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of catalog.ProductTemplate
	items := make([]catalog.ManufacturerTemplate, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallScheduleTask() (response.Install, []tasks.ScheduleTask) {
	var result response.Install
	collection := tasks.CollectionScheduleTask

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "scheduleTasks\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of tasks.ScheduleTask
	items := make([]tasks.ScheduleTask, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallReturnRequestReason() (response.Install, []orders.ReturnRequestReason) {
	var result response.Install
	collection := orders.CollectionReturnRequestReason

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "orders\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ReturnRequestReason
	items := make([]orders.ReturnRequestReason, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}

func InstallReturnRequestAction() (response.Install, []orders.ReturnRequestAction) {
	var result response.Install
	collection := orders.CollectionReturnRequestAction

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "orders\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ReturnRequestAction
	items := make([]orders.ReturnRequestAction, 0)
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Success response
	result.Status = true
	result.Details = collection + " data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result, items
}
