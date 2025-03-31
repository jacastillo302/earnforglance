package controller

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"earnforglance/server/bootstrap"
	configuration "earnforglance/server/domain/configuration"
	response "earnforglance/server/domain/install"
	stores "earnforglance/server/domain/stores"
)

type InstallService struct {
	InstallUsecase response.InstallRepository
	SettingUsecase configuration.SettingRepository
	StoresUsecase  stores.StoreRepository
	Env            *bootstrap.Env
}

func InstallationSettings() response.Install {
	var result response.Install

	result.Status = true
	result.Details = "Installation settings initialized successfully"
	result.CreatedOnUtc = time.Now()

	return result
}

func resolvePath(relativePath string) string {
	workingDir, _ := os.Getwd()
	return workingDir + "/" + relativePath
}

func InstallStores(storesUsecase stores.StoreRepository) response.Install {
	var result response.Install

	// Resolve the relative path
	filePath := resolvePath("service\\data\\json\\store\\store.json")

	// Read the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		result.Status = false
		result.Details = "Failed to read stores JSON file: " + err.Error()
		return result
	}

	// Unmarshal the JSON data into a slice of stores.Store
	stores := make([]stores.Store, 0)
	err = json.Unmarshal(fileData, &stores)
	if err != nil {
		result.Status = false
		result.Details = "Failed to parse stores JSON file: " + err.Error()
		return result
	}

	// Call StoresUsercase.CreateMany to save the data
	errnew := storesUsecase.CreateMany(context.Background(), stores)
	if errnew != nil {
		result.Status = false
		result.Details = "Failed to save stores data: " + errnew.Error()
		return result
	}

	// Success response
	result.Status = true
	result.Details = "Stores data installed successfully"
	result.CreatedOnUtc = time.Now()

	return result
}
