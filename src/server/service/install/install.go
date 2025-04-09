package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"crypto/rand"
	"earnforglance/server/bootstrap"
	affiliates "earnforglance/server/domain/affiliate"
	blogs "earnforglance/server/domain/blogs"
	catalog "earnforglance/server/domain/catalog"
	commons "earnforglance/server/domain/common"
	configuration "earnforglance/server/domain/configuration"
	customers "earnforglance/server/domain/customers"
	directory "earnforglance/server/domain/directory"
	discounts "earnforglance/server/domain/discounts"
	forums "earnforglance/server/domain/forums"
	response "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	loggings "earnforglance/server/domain/logging"
	media "earnforglance/server/domain/media"
	messages "earnforglance/server/domain/messages"
	news "earnforglance/server/domain/news"
	orders "earnforglance/server/domain/orders"
	polls "earnforglance/server/domain/polls"
	tasks "earnforglance/server/domain/scheduleTasks"
	security "earnforglance/server/domain/security"
	shippings "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
	topics "earnforglance/server/domain/topics"
	vendors "earnforglance/server/domain/vendors"
	tools "earnforglance/server/service/common"
	service "earnforglance/server/service/customers"
	"encoding/base64"

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

func InstallShippingMethod() (response.Install, []shippings.ShippingMethod) {
	var result response.Install
	collection := shippings.CollectionShippingMethod

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
	items := make([]shippings.ShippingMethod, 0)
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

func InstallDeliveryDate() (response.Install, []shippings.DeliveryDate) {
	var result response.Install
	collection := shippings.CollectionDeliveryDate

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
	items := make([]shippings.DeliveryDate, 0)
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

func InstallProductAvailabilityRange() (response.Install, []shippings.ProductAvailabilityRange) {
	var result response.Install
	collection := shippings.CollectionProductAvailabilityRange

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
	items := make([]shippings.ProductAvailabilityRange, 0)
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

func InstallCustomer(isSample bool) (response.Install, []customers.Customer) {
	var result response.Install
	collection := customers.CollectionCustomer

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "customers\\"+collection+sample+".json")

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

	//fmt.Println("Hash:", hash)

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

func InstallAddress(isSample bool) (response.Install, []commons.Address) {
	var result response.Install
	collection := commons.CollectionAddress

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "common\\"+collection+sample+".json")

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

func InstallCustomerCustomerRoleMapping_Sample() (response.Install, []customers.CustomerCustomerRoleMapping) {
	var result response.Install
	collection := customers.CollectionCustomerCustomerRoleMapping

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "customers\\"+collection+"_sample.json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of customers.CustomerCustomerRoleMapping
	items := make([]customers.CustomerCustomerRoleMapping, 0)
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

func InstallActivityLog() (response.Install, []loggings.ActivityLog) {
	var result response.Install
	collection := loggings.CollectionActivityLog

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "logging\\"+collection+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of customers.ActivityLog
	items := make([]loggings.ActivityLog, 0)
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

func InstallCheckoutAttribute(isSample bool) (response.Install, []orders.CheckoutAttribute) {
	var result response.Install
	collection := orders.CollectionCheckoutAttribute

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "orders\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of CheckoutAttribute
	items := make([]orders.CheckoutAttribute, 0)
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

func InstallCheckoutAttributeValue(isSample bool) (response.Install, []orders.CheckoutAttributeValue) {
	var result response.Install
	collection := orders.CollectionCheckoutAttributeValue

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "orders\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of CheckoutAttributeValue
	items := make([]orders.CheckoutAttributeValue, 0)
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

func InstallSpecificationAttribute(isSample bool) (response.Install, []catalog.SpecificationAttribute) {
	var result response.Install
	collection := catalog.CollectionSpecificationAttribute

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of SpecificationAttribute
	items := make([]catalog.SpecificationAttribute, 0)
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

func InstallSpecificationAttributeOption(isSample bool) (response.Install, []catalog.SpecificationAttributeOption) {
	var result response.Install
	collection := catalog.CollectionSpecificationAttributeOption

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of SpecificationAttributeOption
	items := make([]catalog.SpecificationAttributeOption, 0)
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

func InstallSpecificationAttributeGroup(isSample bool) (response.Install, []catalog.SpecificationAttributeGroup) {
	var result response.Install
	collection := catalog.CollectionSpecificationAttributeGroup

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of SpecificationAttributeOption
	items := make([]catalog.SpecificationAttributeGroup, 0)
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

func InstallProductAttribute(isSample bool) (response.Install, []catalog.ProductAttribute) {
	var result response.Install
	collection := catalog.CollectionProductAttribute

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of SpecificationAttributeOption
	items := make([]catalog.ProductAttribute, 0)
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

func InstallCategory(isSample bool) (response.Install, []catalog.Category) {
	var result response.Install
	collection := catalog.CollectionCategory

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Category
	items := make([]catalog.Category, 0)
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

func InstallPicture(isSample bool) (response.Install, []media.Picture) {
	var result response.Install
	collection := media.CollectionPicture

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "media\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Picture
	items := make([]media.Picture, 0)
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

func InstallManufacturer(isSample bool) (response.Install, []catalog.Manufacturer) {
	var result response.Install
	collection := catalog.CollectionManufacturer

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Manufacturer
	items := make([]catalog.Manufacturer, 0)
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

func InstallWarehouse(isSample bool) (response.Install, []shippings.Warehouse) {
	var result response.Install
	collection := shippings.CollectionWarehouse

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "shipping\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of shippings.Warehouse
	items := make([]shippings.Warehouse, 0)
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

func InstallProduct(isSample bool) (response.Install, []catalog.Product) {
	var result response.Install
	collection := catalog.CollectionProduct

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Product
	items := make([]catalog.Product, 0)
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

func InstallProductSpecificationAttribute(isSample bool) (response.Install, []catalog.ProductSpecificationAttribute) {
	var result response.Install
	collection := catalog.CollectionProductSpecificationAttribute

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductSpecificationAttribute
	items := make([]catalog.ProductSpecificationAttribute, 0)
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

func InstallProductTag(isSample bool) (response.Install, []catalog.ProductTag) {
	var result response.Install
	collection := catalog.CollectionProductTag

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductTag
	items := make([]catalog.ProductTag, 0)
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

func InstallProductAttributeValue(isSample bool) (response.Install, []catalog.ProductAttributeValue) {
	var result response.Install
	collection := catalog.CollectionProductAttributeValue

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductAttributeValue
	items := make([]catalog.ProductAttributeValue, 0)
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

func InstallProductAttributeMapping(isSample bool) (response.Install, []catalog.ProductAttributeMapping) {
	var result response.Install
	collection := catalog.CollectionProductAttributeMapping

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductAttributeMapping
	items := make([]catalog.ProductAttributeMapping, 0)
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

func InstallProductProductTagMapping(isSample bool) (response.Install, []catalog.ProductProductTagMapping) {
	var result response.Install
	collection := catalog.CollectionProductProductTagMapping

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductProductTagMapping
	items := make([]catalog.ProductProductTagMapping, 0)
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

func InstallProductPicture(isSample bool) (response.Install, []catalog.ProductPicture) {
	var result response.Install
	collection := catalog.CollectionProductPicture

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductPicture
	items := make([]catalog.ProductPicture, 0)
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

func InstallProductCategory(isSample bool) (response.Install, []catalog.ProductCategory) {
	var result response.Install
	collection := catalog.CollectionProductCategory

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductCategory
	items := make([]catalog.ProductCategory, 0)
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

func InstallTierPrice(isSample bool) (response.Install, []catalog.TierPrice) {
	var result response.Install
	collection := catalog.CollectionTierPrice

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of TierPrice
	items := make([]catalog.TierPrice, 0)
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

func InstallProductManufacturer(isSample bool) (response.Install, []catalog.ProductManufacturer) {
	var result response.Install
	collection := catalog.CollectionProductManufacturer

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductManufacturer
	items := make([]catalog.ProductManufacturer, 0)
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

func InstallProductAttributeValuePicture(isSample bool) (response.Install, []catalog.ProductAttributeValuePicture) {
	var result response.Install
	collection := catalog.CollectionProductAttributeValuePicture

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "catalog\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ProductAttributeValuePicture
	items := make([]catalog.ProductAttributeValuePicture, 0)
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

func InstallVendor(isSample bool) (response.Install, []vendors.Vendor) {
	var result response.Install
	collection := vendors.CollectionVendor

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "vendors\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of vendors.Vendor
	items := make([]vendors.Vendor, 0)
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

func InstallAffiliate(isSample bool) (response.Install, []affiliates.Affiliate) {
	var result response.Install
	collection := affiliates.CollectionAffiliate

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "affiliate\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of affiliates.Affiliate
	items := make([]affiliates.Affiliate, 0)
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

func InstallForumGroup(isSample bool) (response.Install, []forums.ForumGroup) {
	var result response.Install
	collection := forums.CollectionForumGroup

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "forums\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of ForumGroup
	items := make([]forums.ForumGroup, 0)
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

func InstallForum(isSample bool) (response.Install, []forums.Forum) {
	var result response.Install
	collection := forums.CollectionForum

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "forums\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Forum
	items := make([]forums.Forum, 0)
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

func InstallDiscount(isSample bool) (response.Install, []discounts.Discount) {
	var result response.Install
	collection := discounts.CollectionDiscount

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "discounts\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Forum
	items := make([]discounts.Discount, 0)
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

func InstallBlogPost(isSample bool) (response.Install, []blogs.BlogPost) {
	var result response.Install
	collection := blogs.CollectionBlogPost

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "blogs\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of blogs.BlogPost
	items := make([]blogs.BlogPost, 0)
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

func InstallBlogComment(isSample bool) (response.Install, []blogs.BlogComment) {
	var result response.Install
	collection := blogs.CollectionBlogComment

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "blogs\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of blogs.BlogComment
	items := make([]blogs.BlogComment, 0)
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

func InstallPoll(isSample bool) (response.Install, []polls.Poll) {
	var result response.Install
	collection := polls.CollectionPoll

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "polls\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of polls.Poll
	items := make([]polls.Poll, 0)
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

func InstallPollAnswer(isSample bool) (response.Install, []polls.PollAnswer) {
	var result response.Install
	collection := polls.CollectionPollAnswer

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "polls\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of polls.PollAnswer
	items := make([]polls.PollAnswer, 0)
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

func InstallNewsItem(isSample bool) (response.Install, []news.NewsItem) {
	var result response.Install
	collection := news.CollectionNewsItem

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "news\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of NewsItem
	items := make([]news.NewsItem, 0)
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

func InstallNewsComment(isSample bool) (response.Install, []news.NewsComment) {
	var result response.Install
	collection := news.CollectionNewsComment

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "news\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of NewsItem
	items := make([]news.NewsComment, 0)
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

func InstallSearchTerm(isSample bool) (response.Install, []commons.SearchTerm) {
	var result response.Install
	collection := commons.CollectionSearchTerm

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "common\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of commons.SearchTerm
	items := make([]commons.SearchTerm, 0)
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

func InstallDownload(isSample bool) (response.Install, []media.Download) {
	var result response.Install
	collection := media.CollectionDownload

	sample := "_sample"

	if !isSample {
		sample = ""
	}

	// Resolve the relative path
	filePath := resolvePath(DefaultPathJson, "media\\"+collection+sample+".json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read " + collection + " JSON file: " + err.Error()
		return result, nil
	}

	// Unmarshal the JSON data into a slice of Download
	items := make([]media.Download, 0)
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
