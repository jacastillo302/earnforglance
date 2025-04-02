package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"earnforglance/server/bootstrap"
	configuration "earnforglance/server/domain/configuration"
	directory "earnforglance/server/domain/directory"
	response "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	messages "earnforglance/server/domain/messages"
	security "earnforglance/server/domain/security"
	shipping "earnforglance/server/domain/shipping"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"

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
