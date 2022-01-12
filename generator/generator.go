package generator

import (
	"fmt"
	"strings"

	"github.com/vikash/api-mocker/models"
)

type Generator interface {
	Generate() interface{}
}

func GenerateObject(entity models.ModelName, store models.ModelStore) interface{} {
	structure := store.StructureForModel(entity)

	object := make(map[string]interface{})

	for fieldName, fieldDefinition := range structure {
		value, err := GenerateValueForConfig(fieldDefinition)
		if err == nil {
			object[fieldName] = value
		} else {
			// Check if a model exists with the field type
			probableType := models.ModelName(strings.ToLower(fieldDefinition["type"].(string)))
			customStructure := store.StructureForModel(probableType)
			if customStructure != nil {
				object[fieldName] = GenerateObject(probableType, store)
			}
		}
	}

	return object
}

func GenerateValueForConfig(def map[string]interface{}) (interface{}, error) {
	objectType := strings.ToLower(def["type"].(string))
	if objectType == "" {
		objectType = "string"
	}

	switch objectType {
	case "number":
		return NewNumber(def).Generate(), nil
	case "string":
		return String(def).Generate(), nil
	case "image":
		return NewImage(def).Generate(), nil
	case "array":
		return NewArray(def).Generate(), nil
	default:
		return nil, fmt.Errorf("type '%s' not defined", objectType)
	}

}
