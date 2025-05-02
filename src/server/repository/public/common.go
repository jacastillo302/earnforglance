package repository

import (
	// ...existing imports...
	"encoding/base64"
	"fmt"
	"reflect"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ToStringAlways(v interface{}) string {

	//fmt.Printf("Go type: %T\n", v) // Print the Go type of v

	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return base64.StdEncoding.EncodeToString(val)
	case fmt.Stringer:
		return val.String()
	case map[string]interface{}:
		if bin, ok := val["$binary"]; ok {
			if binMap, ok := bin.(map[string]interface{}); ok {
				if base64Str, ok := binMap["base64"].(string); ok {
					return base64Str
				}
			}
		}
	case bson.Binary:
		// Handle bson.Binary type
		guid, err := uuid.FromBytes(val.Data)
		if err == nil {

			return guid.String()
		}
		return base64.StdEncoding.EncodeToString(val.Data)
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Slice {
			if rv.Type().Elem().Kind() == reflect.Uint8 {
				return base64.StdEncoding.EncodeToString(v.([]byte))
			}
		}
		return fmt.Sprintf("%v", v)
	}
	return ""
}
