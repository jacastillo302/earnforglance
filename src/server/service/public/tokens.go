package public

import (
	"encoding/json"
	"fmt"
	"os"
)

type TokenGroup struct {
	Group  string   `json:"group"`
	Tokens []string `json:"tokens"`
}

func ReadJsonTokens(groupNames []string) ([]string, error) {
	var tokens []string
	filePath := resolvePath(DefaultPathJson, "messages\\tokens.json")
	// Open the JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}
	// Parse the JSON data into a map
	var groups []TokenGroup
	if err := json.Unmarshal(fileData, &groups); err != nil {
		return nil, err
	}

	for _, groupName := range groupNames {
		for _, g := range groups {
			if g.Group == groupName {
				tokens = append(tokens, g.Tokens...)
			}
		}
	}

	return tokens, nil
}
