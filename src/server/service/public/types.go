package public

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DefaultPathJson = "service\\data\\json"
)

type Type struct {
	Name        string
	Value       int
	Description string
}

func resolvePath(basePath string, relativePath string) string {
	workingDir, _ := os.Getwd()
	return workingDir + "\\" + basePath + "\\" + relativePath
}

func ReadJsonMapTypes(basePath string, name string) ([]Type, error) {

	filePath := resolvePath(DefaultPathJson, basePath+"\\"+name+"_type.json")
	// Open the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}
	// Parse the JSON data into a map
	var items []Type
	err = json.Unmarshal(fileData, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON file: %w", err)
	}

	return items, nil
}

func FilterTypesByValue(types []Type, value int) []Type {
	var filtered []Type
	for _, t := range types {
		if t.Value == value {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
