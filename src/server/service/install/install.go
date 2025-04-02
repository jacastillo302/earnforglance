package controller

import (
	"encoding/json"
	"os"
	"time"

	"earnforglance/server/bootstrap"
	configuration "earnforglance/server/domain/configuration"
	directory "earnforglance/server/domain/directory"
	response "earnforglance/server/domain/install"
	lang "earnforglance/server/domain/localization"
	stores "earnforglance/server/domain/stores"
	taxes "earnforglance/server/domain/tax"
	common "earnforglance/server/service/common"

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

func InstallationSettings() response.Install {
	var result response.Install

	result.Status = true
	result.Details = "Installation settings initialized successfully"
	result.CreatedOnUtc = time.Now()

	return result
}

func resolvePath(basePath string, relativePath string) string {
	workingDir, _ := os.Getwd()
	return workingDir + "\\" + basePath + "\\" + relativePath
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

func GetDefaultLanguage() []lang.Language {

	// Create a default language object with the necessary fields
	defaultLanguage := lang.Language{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		Name:              common.DefaultLanguageName,
		LanguageCulture:   common.DefaultLanguageCulture,
		UniqueSeoCode:     common.DefaultLocalePattern,
		FlagImageFileName: common.DefaultLocalePattern + ".png",
		Rtl:               false,
		LimitedToStores:   false,
		DefaultCurrencyID: primitive.NewObjectID(),
		Published:         true,
		DisplayOrder:      1,
	}

	langs := []lang.Language{defaultLanguage}

	return langs
}
